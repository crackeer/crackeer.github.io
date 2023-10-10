package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/dgrijalva/jwt-go"
)

// GenJwt
//  @param args
//  @param list
//  @return interface{}
func GenJwt(this js.Value, args []js.Value) interface{} {
	raws := args[0].String()
	var tmp map[string]interface{}
	json.Unmarshal([]byte(raws), &tmp)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(tmp))

	data, _ := token.SignedString([]byte("GenJwt"))
	return data

}

func main() {
	// 将golang中foo函数注入到window.foo中
	js.Global().Set("GenJwt", js.FuncOf(GenJwt))
	select {}
}
