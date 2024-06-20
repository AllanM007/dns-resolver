package main

t.Run("Should encode an header into bytes", func(t *testing.T) {
	header := NewHeader(22, RECURSION_FLAG, 1, 0, 0, 0)

	encodedHeader := header.ToBytes()

	expected, err := hex.DecodeString("0016010000010000000000000")
	assert.NotNil(t, err)
	assert.Equal(t, expected, encodedHeader)
})