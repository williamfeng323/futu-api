package futugrpc

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/williamfeng323/futu-api/protos/pb/qotcommon"
	"github.com/williamfeng323/futu-api/protos/pb/qotgetstaticinfo"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func TestClient(t *testing.T) {
	client := NewClient(context.Background(), nil, WithAddr("localhost:11111"), WithKeepAlive(5*time.Second), WithTimeout(10*time.Second), WithMaxRetry(3), WithClientId("Futu-API"))
	assert.NotNil(t, client)
	err := client.Connect()
	assert.Nil(t, err)
	market := int32(qotcommon.QotMarket_QotMarket_CNSH_Security)
	// secType := int32(qotcommon.SecurityType_SecurityType_Eqty)
	req := qotgetstaticinfo.Request{
		C2S: &qotgetstaticinfo.C2S{
			Market: &market,
			// SecType: &secType,
		},
	}
	resp := qotgetstaticinfo.Response{}
	client.DoRequest(uint32(3202), &req, &resp)
	fmt.Printf("%s", ConvertByte2String(resp.S2C.ProtoReflect().GetUnknown(), GB18030))
	assert.NotNil(t, &resp.S2C)
	time.Sleep(5 * time.Second)
	client.Close()
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}
