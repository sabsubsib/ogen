package json

import (
	"github.com/go-faster/jx"
	"github.com/shopspring/decimal"
)

func EncodeDecimal(e *jx.Encoder, v decimal.Decimal) {
	bytes, _ := v.MarshalJSON()
	e.Raw(bytes)
}

func DecodeDecimal(d *jx.Decoder) (v decimal.Decimal, err error) {
	bytes, err := d.Raw()
	if err != nil {
		return v, err
	}

	err = v.UnmarshalJSON(bytes)

	return v, err
}

func EncodeStringDecimal(e *jx.Encoder, v string) {
	EncodeDecimal(e, decimal.RequireFromString(v))
}

func DecodeStringDecimal(d *jx.Decoder) (v decimal.Decimal, err error) {
	return DecodeDecimal(d)
}

func unquote(bytes []byte) string {
	return string(bytes[1 : len(bytes)-1])
}
