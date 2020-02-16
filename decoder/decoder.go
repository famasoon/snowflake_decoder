package decoder

import (
	"strconv"
	"time"
)

type SnowFlake struct {
	ID           string `json:"id"`
	Timestamp    int64  `json:"timestamp"`
	Unixtime     int64  `json:"unixtime"`
	Date         string `json:"date"`
	DatacenterID int64  `json:"datacenterId"`
	WorkerID     int64  `json:"workerId"`
	Sequence     int64  `json:"sequence"`
}

const (
	unixTimeStart = 1288834974657
)

func DecodeSnowFlake(input int64) SnowFlake {
	snowFlake := SnowFlake{}
	snowFlake.ID = strconv.FormatInt(input, 10)

	timeStamp := input >> 22
	snowFlake.Timestamp = timeStamp

	snowFlake.Unixtime = timeStamp + unixTimeStart

	jst, _ := time.LoadLocation("Asia/Tokyo")
	snowFlake.Date = time.Unix((snowFlake.Timestamp+unixTimeStart)/1000.000, 10).In(jst).String()

	datecenter := input >> 17
	datecenter = datecenter & 0x1F
	snowFlake.DatacenterID = datecenter

	workerTemp := input >> 12
	workerTemp = workerTemp & 0x1F
	snowFlake.WorkerID = workerTemp

	snowFlake.Sequence = input & 0xFFF

	return snowFlake
}
