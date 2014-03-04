package main

import (
	"fmt"
	"log"
	// "net"
	// "net/http"
	"crypto/md5"
	"encoding/hex"
	"os"
	"path"
	"strconv"
	// "net/http/httputil"
	"github.com/cznic/exp/dbm"
)

const REVERSE_DIR = "reverse"
const SPLIT_SIZE = 2
const PRINTABLE_START = ' '
const PRINTABLE_END = '~'

const ASCII_PRINTABLE_CHARACTERS = ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~`
const ASCII_PRINTABLE_CHARACTERS_LENGTH = len(ASCII_PRINTABLE_CHARACTERS)

func next(current string) string {
	currentLen := len(current)
	if currentLen == 0 {
		// init
		return string(PRINTABLE_START)
	} else {
		pre, last := current[:currentLen-1], current[currentLen-1]
		if last == PRINTABLE_END {
			return next(pre) + string(PRINTABLE_START)
		} else {
			return pre + string(last+1)
		}
	}
}
func getPath(prefix, hash string) string {
	if len(hash) < SPLIT_SIZE {
		panic("hash string length < " + strconv.Itoa(SPLIT_SIZE))
	} else {
		return path.Join(prefix, hash[:SPLIT_SIZE])
	}
}

func mkdirs(path string) error {
	return os.MkdirAll(path, os.ModeDir)
}

func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	_ = log.Fatal
	_ = fmt.Println
	_ = dbm.ACIDNone
	current := ""
	size0 := 0
	size1 := 0
	size2 := 0
	size3 := 0
	for i := 0; i < 1000000; i++ {
		// md5 := getMd5String(current)
		// dir1 := getPath(REVERSE_DIR, md5)
		// log.Println(current, md5, dir1)
		// fmt.Printf(`"%s", `, current)
		size := len(current)
		switch size {
		case 0:
			size0++
		case 1:
			size1++
		case 2:
			size2++
		case 3:
			size3++
		}
		if current == "eggy" {
			fmt.Println(current, i)
		}
		current = next(current)
	}
	log.Println(size0, size1, size2, size3)
}
