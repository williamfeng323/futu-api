package futugrpc

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
)

var serialNo uint32 = 0 //uint32(rand.Int31())

type Pack struct {
	szHeaderFlag  [2]uint8  // u8_t szHeaderFlag[2];
	nProtoID      uint32    // u32_t nProtoID;
	nProtoFmtType uint8     // u8_t nProtoFmtType;
	nProtoVer     uint8     // u8_t nProtoVer;
	nSerialNo     uint32    // u32_t nSerialNo;
	nBodyLen      uint32    // u32_t nBodyLen;
	arrBodySHA1   [20]uint8 // u8_t arrBodySHA1[20];
	arrReserved   [8]uint8  // u8_t arrReserved[8];
	body          []byte    // []byte add;
}

// SetProtoID set nProtoID
func (p *Pack) SetProtoID(nProtoID uint32) {
	p.nProtoID = nProtoID
}

// SetBody set body
func (p *Pack) SetBody(body []byte) {
	p.body = body
	p.nBodyLen = uint32(len(body))

	sha := sha1.New()
	sha.Write(p.body)
	arrBodySHA1 := sha.Sum(nil)
	copy(p.arrBodySHA1[:], arrBodySHA1)
}

func (p *Pack) GetSerialNoStr() string {
	return fmt.Sprintf("p_%d", p.nSerialNo)
}

// Unpack unpack
func (p *Pack) Unpack(b []byte) error {
	var err error
	reader := bytes.NewReader(b)
	if err = binary.Read(reader, binary.LittleEndian, &p.szHeaderFlag); err != nil {
		errors.Errorf("binary.Read failed: %v", err)
		return err
	}
	if err = binary.Read(reader, binary.LittleEndian, &p.nProtoID); err != nil {
		errors.Errorf("binary.Read failed: %v", err)
		return err
	}
	if err = binary.Read(reader, binary.LittleEndian, &p.nProtoFmtType); err != nil {
		errors.Errorf("binary.Read failed: %v", err)
		return err
	}
	if err = binary.Read(reader, binary.LittleEndian, &p.nProtoVer); err != nil {
		errors.Errorf("binary.Read failed: %v", err)
		return err
	}
	if err = binary.Read(reader, binary.LittleEndian, &p.nSerialNo); err != nil {
		errors.Errorf("binary.Read failed: %v", err)
		return err
	}
	if err = binary.Read(reader, binary.LittleEndian, &p.nBodyLen); err != nil {
		errors.Errorf("binary.Read failed: %v", err)
		return err
	}
	if err = binary.Read(reader, binary.LittleEndian, &p.arrBodySHA1); err != nil {
		errors.Errorf("binary.Read failed: %v", err)
		return err
	}
	if err = binary.Read(reader, binary.LittleEndian, &p.arrReserved); err != nil {
		errors.Errorf("binary.Read failed: %v", err)
		return err
	}

	p.body = make([]byte, p.nBodyLen)
	err = binary.Read(reader, binary.LittleEndian, &p.body)

	return err
}

// to string
func (p *Pack) String() string {
	return fmt.Sprintf("nBodyLen: %d body: %s",
		p.nBodyLen,
		p.body,
	)
}

// GetBody get body data
func (p *Pack) GetBody() []byte {
	return p.body
}
