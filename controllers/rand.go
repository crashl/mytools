package controllers

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var rows, length *int

func init() {
	rows = flag.Int("r", 1, "random number rows")
	length = flag.Int("l", 16, "random number length")
	flag.Parse()
}

func RandomString(rows, length int) {
	// 根据前端选项，源端切片要重新生成，可以利用ASCII码表来操作，或者四个字符串，通过判断来决定哪些来组合成新的字符串
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str) //字节切片
	res := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < rows; i++ {
		for j := 0; j < length; j++ {
			res = append(res, bytes[r.Intn(len(bytes))])
		}
		fmt.Println(string(res))
		res = res[0:0]
	}

}

//func main() {
//	RandomString(*rows, *length)
//}
