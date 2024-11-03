package json

import (
	"github.com/go-faster/jx"
	"github.com/shopspring/decimal"
)

func EncodeDecimal(e *jx.Encoder, v decimal.Decimal) {
	e.Raw([]byte(v.String()))
}

func DecodeDecimal(d *jx.Decoder) (v decimal.Decimal, err error) {
	str, err := d.Str()
	if err != nil {
		return v, err
	}

	return decimal.NewFromString(str)
}

func EncodeStringDecimal(e *jx.Encoder, v string) {
	e.Raw([]byte("\"" + e.String() + "\""))
}

func DecodeStringDecimal(d *jx.Decoder) (v decimal.Decimal, err error) {
	str, err := d.StrBytes()
	if err != nil {
		return v, err
	}

	return decimal.NewFromString(unquote(str))
}

func unquote(bytes []byte) string {
	return string(bytes[1 : len(bytes)-1])
}
