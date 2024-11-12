package main

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"
)

var bufPool sync.Pool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func main() {
	Log(os.Stdout, "path", "/search?keyword=google.com")
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.WriteString(time.Now().String())
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)
}
