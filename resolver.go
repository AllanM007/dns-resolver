package main

import (
	"bytes"
	"encoding/binary"
	"errors"
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

func (q *Question) ToBytes() []byte {
	encodedQuestion := new(bytes.Buffer)
	binary.Write(encodedQuestion, binary.BigEndian, q.QName)
	binary.Write(encodedQuestion, binary.BigEndian, q.QType)
	binary.Write(encodedQuestion, binary.BigEndian, q.QClass)

	return encodedQuestion.Bytes()
}

func encodeDnsName(qname []byte) []byte {
	var encoded []byte
	parts := bytes.Split([]byte(qname), []byte{'.'})
	for _, part := range parts {
		encoded = append(encoded, byte(len(part)))
		encoded = append(encoded, part...)
	}
	return append(encoded, 0x00)
}

func NewQuery(header *Header, question *Question) []byte {
	var query []byte

	query = append(query, header.ToBytes()...)
	query = append(query, question.ToBytes()...)

	return query
}

func ParseHeader(reader *bytes.Reader) (*Header, error) {
	var header Header

	binary.Read(reader, binary.BigEndian, &header.Id)
	binary.Read(reader, binary.BigEndian, &header.Flags)
	switch header.Flags & 0b1111 {
	case 1:
		return nil, errors.New("error with the query")
	case 2:
		return nil, errors.New("error with the server")
	case 3:
		return nil, errors.New("the domain doesn't exist")
	}
	binary.Read(reader, binary.BigEndian, &header.QdCount)
	binary.Read(reader, binary.BigEndian, &header.AnCount)
	binary.Read(reader, binary.BigEndian, &header.NsCount)
	binary.Read(reader, binary.BigEndian, &header.ArCount)

	return &header, nil
}

func ParseQuestion(reader *bytes.Reader) *Question {
	var question Question

	question.QName = []byte(DecodeName(reader))
	binary.Read(reader, binary.BigEndian, &question.QType)
	binary.Read(reader, binary.BigEndian, &question.QClass)

	return &question
}
