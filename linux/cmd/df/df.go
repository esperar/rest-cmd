package df

import (
	"os/exec"
	"strconv"
	"strings"
)

type DiskUsage struct {
	Filesystem string `json:"filesystem"`
	Blocks     int    `json:"512-blocks"`
	Used       int    `json:"used"`
	Available  int    `json:"available"`
	Capacity   string `json:"capacity"`
	Iused      int    `json:"iused"`
	Ifree      int    `json:"ifree"`
	IusedPerc  string `json:"%iused"`
	MountedOn  string `json:"mounted_on"`
}

func Df() ([]DiskUsage, error) {
	cmd := exec.Command("df", "-i")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	diskUsages := make([]DiskUsage, 0, len(lines)-1)

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if fields[0] == "map" {
			// Filesystem is "map auto_home"
			fields[1] = "map " + fields[1]
			fields = fields[1:]
		}

		fileystem := fields[0]
		blocks, _ := strconv.Atoi(fields[1])
		used, _ := strconv.Atoi(fields[2])
		available, _ := strconv.Atoi(fields[3])
		capacity := fields[4]
		iused, _ := strconv.Atoi(fields[5])
		ifree, _ := strconv.Atoi(fields[6])
		iusedPerc := fields[7]
		mountedOn := fields[8]

		diskUsages = append(diskUsages,
			DiskUsage{
				Filesystem: fileystem,
				Blocks:     blocks,
				Used:       used,
				Available:  available,
				Capacity:   capacity,
				Iused:      iused,
				Ifree:      ifree,
				IusedPerc:  iusedPerc,
				MountedOn:  mountedOn,
			})
	}

	return diskUsages, nil
}
