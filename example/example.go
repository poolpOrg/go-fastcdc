package main

import (
        "fmt"
        "io"
        "log"
        "os"
        "time"

        "github.com/poolpOrg/go-fastcdc"
)

func main() {
        if len(os.Args) != 2 {
                log.Fatal("a parameter is required")
        }

        rd, err := os.Open(os.Args[1])
        if err != nil {
                log.Fatal(err)
        }

        chunkerOptions := fastcdc.NewChunkerOptions()

        /*
                bufPool := &sync.Pool{
                        New: func() interface{} {
                                b := make([]byte, chunkerOptions.MaxSize)
                                return &b
                        },
                }
                chunkerOptions.BufferAllocate = func() *[]byte {
                        return bufPool.Get().(*[]byte)
                }

                chunkerOptions.BufferRelease = func(buffer *[]byte) {
                        bufPool.Put(buffer)
                }
        */

        chunker, err := fastcdc.NewChunker(rd, chunkerOptions)
        if err != nil {
                log.Fatal(err)
        }

        t0 := time.Now()
        for {
                chunk, err := chunker.Next()
                if err != nil {
                        if err == io.EOF {
                                break
                        }
                        log.Fatal(err)
                }
                fmt.Println(chunk.Offset, chunk.Size)
        }
        t1 := time.Since(t0)
        fmt.Println(t1)

}
