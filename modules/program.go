package amdsi

// Program contains all kernels
type Program struct {
  Kernels []*Kernel
}

// NewProgram creates an program object
func NewProgram() *Program {
  prg := new(Program)
  return prg
}
