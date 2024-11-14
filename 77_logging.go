// https://gobyexample.com/logging

package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
	// "log/slog"
)

func main() {
	log.Println("standart logger")  // 2024/11/13 16:40:39 standart logger

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro") // 2024/11/13 16:40:39.956573 with micro

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")  // 2024/11/13 16:40:39 77_logging.go:21: with file/line

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")  // my:2024/11/13 16:40:39 from mylog

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog") // ohmy:2024/11/13 16:42:21 from mylog

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello")
	fmt.Print("from buflog:", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")
	myslog.Info("hello again", "key", "val", "age", 25)
	myslog.Error("error mess", "fff")

}

// 16:47:28 $ go run 77_logging.go 

// 2024/11/13 16:48:08 standart logger
// 2024/11/13 16:48:08.184775 with micro
// 2024/11/13 16:48:08 77_logging.go:21: with file/line
// my:2024/11/13 16:48:08 from mylog
// ohmy:2024/11/13 16:48:08 from mylog
// from buflog:buf:2024/11/13 16:48:08 hello
// {"time":"2024-11-13T16:48:08.18486587+03:00","level":"INFO","msg":"hi there"}
// {"time":"2024-11-13T16:48:08.18491486+03:00","level":"INFO","msg":"hello again","key":"val","age":25}
// {"time":"2024-11-13T16:48:08.18493106+03:00","level":"ERROR","msg":"error mess","!BADKEY":"fff"}
