package test

import (
	"github.com/orestonce/Ip2regionTool"
	"testing"
)

func TestParse(t *testing.T) {
	errMsg := Ip2regionTool.ConvertDb(Ip2regionTool.ConvertDbReq{
		FromName:         "ip2region_v4.xdb",
		FromType:         "Linsoul2014v2, Linsoul2014 第二版xdb格式",
		ToName:           "1.csv",
		ToType:           "csv,纯文本格式",
		VerifyFullUint32: false,
		FillFullUint32:   false,
		MergeIpRange:     false,
	})
	t.Log(errMsg)
}
