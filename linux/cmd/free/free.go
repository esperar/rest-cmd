package free

import (
	"fmt"
	"os/exec"
	"strings"
)

func Free() (MemoryInfo, error) {
	cmd := exec.Command("free", "-h")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return MemoryInfo{}, err
	}

	lines := strings.Split(string(output), "\n")
	if len(lines) < 2 {
		return MemoryInfo{}, fmt.Errorf("unexepcted output format free")
	}

	fields := strings.Fields(lines[1])
	if len(fields) < 7 {
		return MemoryInfo{}, fmt.Errorf("unexpected nuber of fileds in output free")
	}

	return MemoryInfo{
		Total:     fields[1],
		Used:      fields[2],
		Free:      fields[3],
		Shared:    fields[4],
		Buffer:    fields[5],
		Available: fields[6],
	}, nil
}
