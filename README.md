# go-fastcdc

go-fastcdc is a Golang package implementing the fastCDC chunking algorithm.

This is a work in progress.

```go
    chunkerOpts := fastcdc.ChunkerOpts{
        NormalSize : 8 * 1024 * 1024,   // 8KB
    }

    chunker, err := fastcdc.NewChunker(rd, &chunkerOpts)
    if err != nil {
        log.Fatal(err)
    }

    for {
        chunk, err := chunker.Next()
        if err != nil {
            if err == io.EOF {
                // no more chunks to read
                break
            }
            log.Fatal(err)
        }

        fmt.Println(chunk.Offset, chunk.Size)
        // data is in chunk.Data
    }
```