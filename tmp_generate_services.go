package main

import (
	"os"
	"time"

	"github.com/groovy-sky/aws-docs/internal/write"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	writer := write.New(cwd)
	err = writer.WriteServicesIndex("docs", write.ServicesIndexRunInfo{
		Mode:        "manual",
		Status:      "success",
		Error:       "",
		GeneratedAt: time.Now().UTC(),
	})
	if err != nil {
		panic(err)
	}
}
