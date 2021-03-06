package amdsi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRegister(t *testing.T) {
	assert := assert.New(t)

	testAsm := "v23"
	result := IsRegister(testAsm)
	assert.True(result)

	testAsm = "s10"
	result = IsRegister(testAsm)
	assert.True(result)

	testAsm = "v[10:14]"
	result = IsRegister(testAsm)
	assert.True(result)

	testAsm = "s[10:14]"
	result = IsRegister(testAsm)
	assert.True(result)

	testAsm = "label"
	result = IsRegister(testAsm)
	assert.False(result)
}

func TestParseRegisters(t *testing.T) {
	assert := assert.New(t)

	testAsm := "v11"
	result := ParseRegisters(testAsm)
	assert.Equal(11, result[0].Index)
	assert.Equal("v", result[0].Type)

	testAsm = "s20"
	result = ParseRegisters(testAsm)
	assert.Equal(20, result[0].Index)
	assert.Equal("s", result[0].Type)

	testAsm = "v[20:23]"
	result = ParseRegisters(testAsm)
	assert.Equal(20, result[0].Index)
	assert.Equal("v", result[0].Type)
	assert.Equal(21, result[1].Index)
	assert.Equal("v", result[1].Type)
	assert.Equal(22, result[2].Index)
	assert.Equal("v", result[2].Type)
	assert.Equal(23, result[3].Index)
	assert.Equal("v", result[3].Type)

	testAsm = "s[20:23]"
	result = ParseRegisters(testAsm)
	assert.Equal(20, result[0].Index)
	assert.Equal("s", result[0].Type)
	assert.Equal(21, result[1].Index)
	assert.Equal("s", result[1].Type)
	assert.Equal(22, result[2].Index)
	assert.Equal("s", result[2].Type)
	assert.Equal(23, result[3].Index)
	assert.Equal("s", result[3].Type)
}
