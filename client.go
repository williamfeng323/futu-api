package futugrpc

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/williamfeng323/futu-api/protos/pb/initconnect"
	"github.com/williamfeng323/futu-api/protos/pb/keepalive"
	"go.uber.org/zap"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

var defaultOptions = Options{
	addr:      "localhost:13888",
	keepAlive: 5 * time.Second,
	timeout:   10 * time.Second,
	clientId:  "Futu-API",
}
var ClientVersion int32 = 100

type Client struct {
	opts        *Options
	ctx         context.Context
	packChanMap sync.Map
	conn        net.Conn
	logger      Logger
	close       chan int
	isConnected bool
}

func NewClient(ctx context.Context, logger Logger, opts ...Option) *Client {
	if ctx == nil {
		ctx = context.Background()
	}
	if logger == nil {
		zlogger, _ := zap.NewProduction()
		defer zlogger.Sync() // flushes buffer, if any
		logger = zlogger.Sugar()

	}
	c := Client{
		ctx:    ctx,
		logger: logger,
		close:  make(chan int),
	}
	c.opts = &defaultOptions
	for _, opt := range opts {
		opt(c.opts)
	}
	return &c
}

func (c *Client) Close() error {
	if err := c.conn.Close(); err != nil {
		c.logger.Errorf("close connection failed: %v", err)
		return err
	}
	c.isConnected = false
	c.close <- 1
	return nil
}

func (c *Client) DoRequest(protoId uint32, request proto.Message, response proto.Message) error {
	pack := &Pack{}
	pack.SetProtoID(protoId)
	body, err := proto.Marshal(request)
	if err != nil {
		return err
	}
	pack.SetBody(body)
	respPack, err := c.SyncDo(pack)
	if err != nil {
		return err
	}
	b := respPack.GetBody()
	return proto.Unmarshal(b, response)
}

func (c *Client) SyncDo(p *Pack) (responsePack *Pack, err error) {
	ch := make(chan *Pack)
	defer close(ch)

	// szHeaderFlag
	p.szHeaderFlag[0] = uint8('F')
	p.szHeaderFlag[1] = uint8('T')

	// nProtoFmtType
	p.nProtoFmtType = uint8(0)

	// nProtoVer
	p.nProtoVer = 0

	// nSerialNo
	//serialNo += 1
	atomic.AddUint32(&serialNo, 1)
	p.nSerialNo = atomic.LoadUint32(&serialNo)

	// arrBodySHA1
	// fmt.Println("debug: ", p.arrBodySHA1)
	// os.Exit(1)

	// arrReserved
	var arrReservedTmp [20]uint8
	copy(p.arrReserved[:], arrReservedTmp[:20])

	if err = binary.Write(c.conn, binary.LittleEndian, &p.szHeaderFlag); err != nil {
		err = errors.Errorf("binary.Write failed: %v", err)
		return
	}
	if err = binary.Write(c.conn, binary.LittleEndian, &p.nProtoID); err != nil {
		err = errors.Errorf("binary.Write failed: %v", err)
		return
	}
	if err = binary.Write(c.conn, binary.LittleEndian, &p.nProtoFmtType); err != nil {
		err = errors.Errorf("binary.Write failed: %v", err)
		return
	}
	if err = binary.Write(c.conn, binary.LittleEndian, &p.nProtoVer); err != nil {
		err = errors.Errorf("binary.Write failed: %v", err)
		return
	}
	if err = binary.Write(c.conn, binary.LittleEndian, &p.nSerialNo); err != nil {
		err = errors.Errorf("binary.Write failed: %v", err)
		return
	}
	if err = binary.Write(c.conn, binary.LittleEndian, &p.nBodyLen); err != nil {
		err = errors.Errorf("binary.Write failed: %v", err)
		return
	}
	if err = binary.Write(c.conn, binary.LittleEndian, &p.arrBodySHA1); err != nil {
		err = errors.Errorf("binary.Write failed: %v", err)
		return
	}
	if err = binary.Write(c.conn, binary.LittleEndian, &p.arrReserved); err != nil {
		err = errors.Errorf("binary.Write failed: %v", err)
		return
	}
	if err = binary.Write(c.conn, binary.LittleEndian, &p.body); err != nil {
		err = errors.Errorf("binary.Write failed: %v", err)
		return
	}

	sid := p.GetSerialNoStr()
	c.packChanMap.Store(sid, ch)
	responsePack = <-ch
	return
}
func (c *Client) WatchMessage() {
	scanner := bufio.NewScanner(c.conn)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == 'F' {
			if len(data) > 44 {
				length := uint32(0)
				binary.Read(bytes.NewReader(data[12:16]), binary.LittleEndian, &length)
				if int(length)+4 <= len(data) {
					return int(length) + 44, data[:int(length)+44], nil
				}
			}
		}
		return
	})
	//buf := make([]byte, 64 * 1024)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	//scanner.Buffer(buf, bufio.MaxScanTokenSize)

	for scanner.Scan() {
		if !c.isConnected {
			c.logger.Info("futu api client is disconnected")
			return
		}
		scannedPack := new(Pack)
		err := scannedPack.Unpack(scanner.Bytes())
		if err != nil {
			return
		}

		sid := scannedPack.GetSerialNoStr()
		chInterface, ok := c.packChanMap.Load(sid)
		if ok {
			ch := chInterface.(chan *Pack)
			ch <- scannedPack
		}
	}
	if err := scanner.Err(); err != nil {
		c.logger.Errorf("无效数据包 %v", err)
	}
}

func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", c.opts.addr)
	if err != nil {
		return fmt.Errorf("failed to establish tcp connection to %s", c.opts.addr)
	}
	c.conn = conn
	c.isConnected = true
	go func() {
		c.logger.Info("futu api client is established, addr:", c.opts.addr)
		c.WatchMessage()
	}()
	c.initConnect()
	go func() {
		tm := time.NewTicker(c.opts.keepAlive)
		for {
			select {
			case <-c.close:
				c.logger.Info("futu api client done")
				close(c.close)
				tm.Stop()
			case <-c.ctx.Done():
				c.logger.Info("futu api client done")
				close(c.close)
				tm.Stop()
			case <-tm.C:
				c.keepAlive()
			}
		}
	}()
	return nil
}

func (c *Client) initConnect() (resp *initconnect.Response, err error) {
	req := &initconnect.Request{}
	req.Reset()
	req.C2S = &initconnect.C2S{
		ClientID:  &c.opts.clientId,
		ClientVer: &ClientVersion,
	}
	resp = &initconnect.Response{}
	err = c.DoRequest(uint32(1001), req, resp)
	return
}

func (c *Client) keepAlive() (resp *keepalive.Response, err error) {
	t := int64(time.Now().Second())
	req := &keepalive.Request{
		C2S: &keepalive.C2S{
			Time: &t,
		},
	}
	resp = &keepalive.Response{}
	err = c.DoRequest(uint32(1004), req, resp)
	return
}
