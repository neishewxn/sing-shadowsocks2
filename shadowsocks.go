package shadowsocks

import (
	C "github.com/neishewxn/sing-shadowsocks2/cipher"
	_ "github.com/neishewxn/sing-shadowsocks2/shadowaead"
	_ "github.com/neishewxn/sing-shadowsocks2/shadowaead_2022"
	_ "github.com/neishewxn/sing-shadowsocks2/shadowstream"
)

type (
	Method        = C.Method
	MethodOptions = C.MethodOptions
)

func CreateMethod(method string, options MethodOptions) (Method, error) {
	return C.CreateMethod(method, options)
}
