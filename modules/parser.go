package amdsi

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

func Parse(path string) (err error) {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		glog.Fatal(err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	prg := NewProgram()
	section := ""

	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ".global") {
			namedGroup := parseNamedGroup(`\.global (?P<name>\w+)`, line)
			kernel := NewKernel(namedGroup["name"])
			prg.Kernels = append(prg.Kernels, kernel)
			continue
		} else if strings.Contains(line, ".metadata") {
			section = ".metadata"
			continue
		} else if strings.Contains(line, ".args") {
			section = ".args"
			continue
		} else if strings.Contains(line, ".text") {
			section = ".text"
			continue
		}

		kernel := prg.Kernels[len(prg.Kernels)-1]
		switch section {
		case ".metadata":
			kernel.Metadata = append(kernel.Metadata, line)
			namedGroup := parseNamedGroup(`(VGPRs = (?P<vgprs>\d+))|(SGPRs = (?P<sgprs>\d+))`, line)
			maxvgprs := namedGroup["vgprs"]
			if maxvgprs != "" {
				kernel.MaxVRegs, _ = strconv.Atoi(maxvgprs)
			}
			maxsgprs := namedGroup["sgprs"]
			if maxsgprs != "" {
				kernel.MaxSRegs, _ = strconv.Atoi(maxsgprs)
			}
		case ".args":
			kernel.Args = append(kernel.Args, line)
		case ".text":
			inst := NewInstruction(line)
			if IsValidInst(inst.InstText) {
				kernel.Insts = append(kernel.Insts, inst)
			}
		}
	}

	for _, kernel := range prg.Kernels {
		// kernel.PrintGPRs()
		kernel.PrintInstCount()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}
