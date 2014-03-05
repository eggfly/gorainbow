package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strconv"
)

const REVERSE_DIR = "reverse"
const SPLIT_SIZE = 2
const PRINTABLE_START = ' '
const PRINTABLE_END = '~'

// const ASCII_PRINTABLE_CHARACTERS = ` !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~`
// const ASCII_PRINTABLE_CHARACTERS_LENGTH = len(ASCII_PRINTABLE_CHARACTERS)

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	initDB()

	current := dbStatusGetCurrent()
	var count int64 = 0
	totalCount := dbStatusGetCount()

	running := true

	defer func() {
		db.Close()
		log.Println("after db.Close()")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		for _ = range sigChan {
			running = false
			log.Println("Ctrl+C detected, stopping..")
		}
	}()

	for running {
		md5 := getMd5String(current)
		dbHash.Set(md5, current)
		dbReverse.Set(current, md5)
		current = next(current)
		count++
		if count%10000 == 0 {
			log.Printf("Calculate Count: %d, Total Count: %d, Current: %s\n",
				count, totalCount+count, current)
		}
	}

	// finally save status and close db in defer func
	dbStatus.Set(current, _DB_STATUS_CURRENT)
	dbStatus.Set(totalCount+count, _DB_STATUS_COUNT)
}
