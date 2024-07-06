package main

t.Run("Should encode an header into bytes", func(t *testing.T) {
	header := NewHeader(22, RECURSION_FLAG, 1, 0, 0, 0)

	encodedHeader := header.ToBytes()

	expected, err := hex.DecodeString("0016010000010000000000000")
	assert.NotNil(t, err)
	assert.Equal(t, expected, encodedHeader)
})

t.Run("Should encode a question into bytes", func(t *testing.T) {
	question := NewQuestion("dns.google.com", TYPE_A, CLASS_IN)

	encodedQuestion := question.ToBytes()

	expected, _ := hex.DecodeString("03646e7306676f6f676c6503636f6d0000010001")
	assert.NotNil(t, expected)
	assert.Equal(t, expected, encodedQuestion)
})

t.Run("Should encode the dns name", func(t *testing.T) {
	encodedDnsName := encodeDnsName([]byte("dns.google.com"))
	assert.Equal(t, []byte("\x03dns\x06google\x03com\x00"), encodedDnsName)
})

t.Run("Should create a query", func(t *testing.T) {
	header := NewHeader(22, RECURSION_FLAG, 1, 0, 0, 0)
	question := NewQuestion("dns.google.com", TYPE_A, CLASS_IN)

	query := NewQuery(header, question)

	expected, err := hex.DecodeString("00160100000100000000000003646e7306676f6f676c6503636f6d0000010001")
	assert.Nil(t, err)
	assert.Equal(t, expected, query)
})

t.Run("Should check if the response starts with the same ID as the query", func(t *testing.T) {
	query, _ := hex.DecodeString("00160100000100000000000003646e7306676f6f676c6503636f6d0000010001")
	response, _ := hex.DecodeString("00168080000100020000000003646e7306676f6f676c6503636f6d0000010001c00c0001000100000214000408080808c00c0001000100000214000408080404")

	assert.True(t, hasTheSameID(query, response))
})

t.Run("Should create an header from a response", func(t *testing.T) {
	response, _ := hex.DecodeString("001680800001000200000000")
	header, _ := ParseHeader(bytes.NewReader(response))

	assert.Equal(t, &Header{
		Id:      0x16,
		Flags:   1<<15 | 1<<7, // QR (Response) bit = 1, OPCODE = 0 (standard query), AA = 1, TC = 0, RD (Recursion Desired) bit = 1, RA = 1, Z = 0, RCODE = 0
		QdCount: 0x1,
		AnCount: 0x2,
		NsCount: 0x0,
		ArCount: 0x0,
	}, header)
})

t.Run("Should return an error if the header flags contains a query error", func(t *testing.T) {
	response, _ := hex.DecodeString("001680810001000200000000")

	header, err := ParseHeader(bytes.NewReader(response))

	assert.Nil(t, header)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "error with the query")
})