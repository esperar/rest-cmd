package vm_stat

import (
	"os/exec"
	"strings"
)

const pageSize = 4096 // macos page size in bytes

func Vmstat() (MemoryStats, error) {
	cmd := exec.Command("vm_stat")
	output, err := cmd.Output()
	if err != nil {
		return MemoryStats{}, err
	}
	return parseVMStatOutput(string(output))
}

func parseVMStatOutput(input string) (MemoryStats, error) {
	stats := MemoryStats{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// Split the line into key and value
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue // Skip lines that don't match the format
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Pages free":
			stats.PagesFree = value
		case "Pages active":
			stats.PagesActive = value
		case "Pages inactive":
			stats.PagesInactive = value
		case "Pages speculative":
			stats.PagesSpeculative = value
		case "Pages throttled":
			stats.PagesThrottled = value
		case "Pages wired down":
			stats.PagesWiredDown = value
		case "Pages purgeable":
			stats.PagesPurgeable = value
		case "Translation faults":
			stats.TranslationFaults = value
		case "Pages copy-on-write":
			stats.PagesCopyOnWrite = value
		case "Pages zero filled":
			stats.PagesZeroFilled = value
		case "Pages reactivated":
			stats.PagesReactivated = value
		case "Pages purged":
			stats.PagesPurged = value
		case "File-backed pages":
			stats.FileBackedPages = value
		case "Anonymous pages":
			stats.AnonymousPages = value
		case "Pages stored in compressor":
			stats.PagesStoredInCompressor = value
		case "Pages occupied by compressor":
			stats.PagesOccupiedByCompressor = value
		case "Decompressions":
			stats.Decompressions = value
		case "Compressions":
			stats.Compressions = value
		case "Pageins":
			stats.Pageins = value
		case "Pageouts":
			stats.Pageouts = value
		case "Swapins":
			stats.Swapins = value
		case "Swapouts":
			stats.Swapouts = value
		}
	}

	return stats, nil
}
