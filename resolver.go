package main

import (
	"bytes"
	"encoding/binary"
)

const RECURSION_FLAG uint16 = 1 << 8

type Header struct {
	Id      uint16
	Flags   uint16
	QdCount uint16
	AnCount uint16
	NsCount uint16
	ArCount uint16
}

type Mapper struct {
	Domain string
	IPV4   string
	IPV6   string
}

type Question struct {
	QName  string //e.g 3dns6google3com
	QType  string //e.g A,MX
	QClass string //e.g internet
}

func (h *Header) ToBytes() []byte {
	encodedHeader := new(bytes.Buffer)
	binary.Write(encodedHeader, binary.BigEndian, h.Id)
	binary.Write(encodedHeader, binary.BigEndian, h.Flags)
	binary.Write(encodedHeader, binary.BigEndian, h.QdCount)
	binary.Write(encodedHeader, binary.BigEndian, h.AnCount)
	binary.Write(encodedHeader, binary.BigEndian, h.NsCount)
	binary.Write(encodedHeader, binary.BigEndian, h.ArCount)

	return encodedHeader.Bytes()
}
