SnowFlake Decoder
===

This tool decode snowflake ID.
*What is Snowflake*
(Snowflake is a network service for generating unique ID numbers at high scale with some simple guarantees.) [Implementations link](https://github.com/twitter-archive/snowflake/tree/b3f6a3c6ca8e1b6847baa6ff42bf72201e2c2231)

## Background
Looking at [this series of tweets](https://twitter.com/conspirator0/status/1227063249440509952), I thought that it would be nice to be able to easily decode snowflake and created it.

## Installation
```sh
go get github.com/famasoon/snowflake_decoder
```

## Usage
```sh
$ snowflake_decoder [SnowFlake ID]
```

### Example
```sh
$ snowflake_decoder 1228330164804837377
ID: 1228330164804837377
TimeStamp: 292856732560
UnixTime: 1581691707217
Date: 2020-02-14 23:48:27.00000001 +0900 JST
WorkerID: 14
DatacenterID: 11
Sequence: 1
```

## For reference
- [Tweets](https://twitter.com/conspirator0/status/1227063249440509952)
- [Snowflake github](https://github.com/twitter-archive/snowflake/tree/b3f6a3c6ca8e1b6847baa6ff42bf72201e2c2231)
- [Explanation of snowflake](https://yoshipc.net/tweet-id-to-mili-sec/)
- [Explanation of snowflake 2](https://www.slideshare.net/moaikids/20130901-snowflake)