package amdsi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInstruction(t *testing.T)  {
  assert := assert.New(t)

  testAsm := "	 tbuffer_load_format_x  v6, v6, s[12:15], 0 offen format:[BUF_DATA_FORMAT_32,BUF_NUM_FORMAT_FLOAT] // 00000158: EBA01000 80030606"
  result := NewInstruction(testAsm)
  assert.Equal("MTBUF", result.InstType)
  assert.Equal("tbuffer_load_format_x", result.InstText)
  assert.Equal("v", result.DstRegs[0].Type)
  assert.Equal(6, result.DstRegs[0].Index)
  assert.Equal("v", result.SrcRegs[0].Type)
  assert.Equal(6, result.SrcRegs[0].Index)
  assert.Equal("s", result.SrcRegs[1].Type)
  assert.Equal(12, result.SrcRegs[1].Index)
  assert.Equal("s", result.SrcRegs[2].Type)
  assert.Equal(13, result.SrcRegs[2].Index)
  assert.Equal("s", result.SrcRegs[3].Type)
  assert.Equal(14, result.SrcRegs[3].Index)
  assert.Equal("s", result.SrcRegs[4].Type)
  assert.Equal(15, result.SrcRegs[4].Index)
}
