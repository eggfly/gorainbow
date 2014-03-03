package main

import (
	// "fmt"
	// "log"
	// "net"
	// "net/http"
	"crypto/md5"
	"encoding/hex"
	// "net/http/httputil"
)

func main() {
	// empty := md5.Sum([]byte(""))
	// log.Println(empty, hex.EncodeToString([]byte(empty)))
}

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
