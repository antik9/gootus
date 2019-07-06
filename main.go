package main

import (
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	src   = kingpin.Flag("from", "Source file").Required().String()
	dest  = kingpin.Flag("to", "Destincation file").Required().String()
	limit = kingpin.Flag("limit", "Limit bytes").Default(
		strconv.FormatInt(DefaultLimit, 10)).Int64()
	offset = kingpin.Flag("offset", "Offset from the beginning").Default(
		strconv.FormatInt(DefaultOffset, 10)).Int64()
)

func main() {
	kingpin.Parse()
	Copy(*src, *dest, *offset, *limit)
}
