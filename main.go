package main

import (
	"fmt"
	"os"
	"snowflake_decoder/decoder"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s [ID]", os.Args[0])
		os.Exit(0)
	}

	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	snowFlake := decoder.DecodeSnowFlake(int64(id))
	fmt.Printf("ID: %s\n", snowFlake.ID)
	fmt.Printf("TimeStamp: %v\n", snowFlake.Timestamp)
	fmt.Printf("UnixTime: %d\n", snowFlake.Unixtime)
	fmt.Printf("Date: %s\n", snowFlake.Date)
	fmt.Printf("WorkerID: %d\n", snowFlake.WorkerID)
	fmt.Printf("DatacenterID: %d\n", snowFlake.DatacenterID)
	fmt.Printf("Sequence: %d\n", snowFlake.Sequence)
}
