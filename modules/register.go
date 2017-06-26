package amdsi

import (
	"regexp"
	"strconv"
)

// Register represent a register
type Register struct {
	Type  string
	Index int
}

var regexTable = map[string]string{
	`registers`: `v(?P<vreg>\d+)|s(?P<sreg>\d+)|v\[(?P<vregs_start>\d+):(?P<vregs_end>\d+)\]|s\[(?P<sregs_start>\d+):(?P<sregs_end>\d+)\]`,
}

// NewRegister creates a register
func NewRegister(index int, regType string) *Register {
	reg := new(Register)
	reg.Index = index
	reg.Type = regType
	return reg
}

// IsRegister return if the string represents register
func IsRegister(asm string) bool {
	re := regexp.MustCompile(regexTable["registers"])
	result := re.MatchString(asm)
	return result
}

// ParseRegisters parse the asm string and returns a list of registes
func ParseRegisters(asm string) []*Register {
	regs := []*Register{}

	registerMap := parseNamedGroup(regexTable["registers"], asm)
	for key, val := range registerMap {
		if val != "" {
			switch key {
			case "vreg":
				idx, _ := strconv.Atoi(val)
				reg := NewRegister(idx, "v")
				regs = append(regs, reg)
			case "sreg":
				idx, _ := strconv.Atoi(val)
				reg := NewRegister(idx, "s")
				regs = append(regs, reg)
			case "vregs_start":
				startIdx, _ := strconv.Atoi(val)
				endIdx, _ := strconv.Atoi(registerMap["vregs_end"])
				for idx := startIdx; idx <= endIdx; idx++ {
					reg := NewRegister(idx, "v")
					regs = append(regs, reg)
				}
			case "sregs_start":
				startIdx, _ := strconv.Atoi(val)
				endIdx, _ := strconv.Atoi(registerMap["sregs_end"])
				for idx := startIdx; idx <= endIdx; idx++ {
					reg := NewRegister(idx, "s")
					regs = append(regs, reg)
				}
			case "vregs_end", "sregs_end":
			default:
				panic("Invalid key")
			}
		}
	}

	return regs
}
