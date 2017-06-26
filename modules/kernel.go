package amdsi

import (
	"fmt"
	"sort"
)

// Kernel represents an kernel binary
type Kernel struct {
	Name     string
	Metadata []string
	Args     []string
	Insts    []*Instruction
	MaxVRegs int
	MaxSRegs int
}

// NewKernel returns a kernel
func NewKernel(name string) *Kernel {
	k := new(Kernel)
	k.Name = name
	return k
}

func (k *Kernel) PrintGPRs() {
	fmt.Print(k.Name, ": VGPRs = ", k.MaxVRegs, " SGPRs = ", k.MaxSRegs, "\n")
	for _, inst := range k.Insts {
		fmt.Print(inst.InstType, "\t")
		inst.PrintVRegs(k.MaxVRegs)
		fmt.Print(" ")
		inst.PrintSRegs(k.MaxSRegs)
		// fmt.Print(" ")
		// fmt.Print(inst.Raw)
		fmt.Println()
	}
}

type InstCount struct {
	InstTypeCount      map[string]int
	InstTypeVRegsCount map[string][102]int
	InstTypeSRegsCount map[string][255]int
}

func (k *Kernel) PrintInstCount() {
	fmt.Println(k.Name)
	instCount := map[string]int{}
	instTypeCount := map[string]int{}
	instTypeVRegsCount := map[string]map[int]bool{}
	instTypeSRegsCount := map[string]map[int]bool{}

	for _, inst := range k.Insts {
		instType := inst.InstType

		// Count of each instruction
		if _, ok := instCount[inst.InstText]; ok {
			instCount[inst.InstText]++
		} else {
			instCount[inst.InstText] = 1
		}

		// Count instruction by type
		if _, ok := instTypeCount[instType]; ok {
			instTypeCount[instType]++
		} else {
			instTypeCount[instType] = 1
		}

		// Count vregs by instuction type
		vregs := inst.VRegs
		if _, ok := instTypeVRegsCount[instType]; !ok {
			instTypeVRegsCount[instType] = map[int]bool{}
		}
		for idx := 0; idx < len(vregs); idx++ {
			vregStatus := vregs[idx]
			if vregStatus != 0 {
				instTypeVRegsCount[instType][idx] = true
			}
		}

		// Count sregs by instuction type
		sregs := inst.SRegs
		if _, ok := instTypeSRegsCount[instType]; !ok {
			instTypeSRegsCount[instType] = map[int]bool{}
		}
		for idx := 0; idx < len(sregs); idx++ {
			sregStatus := sregs[idx]
			if sregStatus != 0 {
				instTypeSRegsCount[instType][idx] = true
			}
		}
	}

	fmt.Printf("\nBy inst: %d\n", len(instCount))
	keys := []string{}
	for key := range instCount {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s\t%d\n",
			key, instCount[key])
	}

	fmt.Printf("\nBy type: %d\n", len(instTypeCount))

	keys = []string{}
	for key := range instTypeCount {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s\t%d\t%d\t%d\n",
			key, instTypeCount[key],
			len(instTypeVRegsCount[key]), len(instTypeSRegsCount[key]))
	}
}
