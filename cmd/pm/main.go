package main

import (
	"bufio"
	"github.com/alexflint/go-arg"
	"log"
	"os"
	"pipe-monitor/internal"
	"time"
)

type args struct {
	Input_File string `arg:"positional" help:"Optional input file. If not provided input will be read from STDIN"`
	Size       int64  `arg:"-s,--size" help:"Size of input from STDIN. Ignored if using INPUT_FILE"`
	Name       string `arg:"-n,--name" help:"A NAME tag for this output. Will be pre-pended to default FORMAT string"`
	Format     string `arg:"-f,--format" help:"Output format string. Allowed keys: %name, %size, %time, %eta, %percent, %written, %buffered"`
}

func main() {

	arguments := new(args)

	arg.MustParse(arguments)

	run(*arguments)
}

func run(params args) {
	var err error

	var stats internal.Stats
	stats.Name = params.Name
	stats.Size = params.Size
	stats.Format = params.Format
	stats.StartTime = time.Now()
	var file = os.Stdin
	if len(params.Input_File) > 1 {
		file, err = os.Open(params.Input_File)
		fileDes, _ := file.Stat()
		stats.Size = fileDes.Size()
	}

	err = runPipe(file, &stats)

	if err != nil {
		log.Fatal(err)
	}
}

func runPipe(input *os.File, stats *internal.Stats) error {

	stats.WriteStats(true)

	bufferSize := 10 * 1024 * 1024
	readSize := 4 * 1024

	var count int
	reader := bufio.NewReaderSize(input, bufferSize)

	data := make([]byte, readSize)

	var read int
	var err error
	stats.Buffered = reader.Buffered()
	for read, err = reader.Read(data); err == nil; read, err = reader.Read(data) {
		stats.Buffered = reader.Buffered()
		count += read
		_, err := os.Stdout.Write(data[:read])
		if err != nil {
			return err
		}
		stats.Record(read)
		stats.WriteStats(false)
	}
	stats.WriteStats(true)

	return nil
}
