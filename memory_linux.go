//go:build linux
// +build linux

package memory

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func sysTotalMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Totalram) * uint64(in.Unit)
}

func sysFreeMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Freeram) * uint64(in.Unit)
}

func sysAvailableMemory() uint64 {
	// MemAvailable was added in Linux 3.14 and provides a more accurate
	// estimate of available memory than Freeram, as it accounts for
	// reclaimable buffers/cache.
	f, err := os.Open("/proc/meminfo")
	if err != nil {
		// Fall back to Freeram if we can't read /proc/meminfo
		return sysFreeMemory()
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "MemAvailable:") {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				val, err := strconv.ParseUint(fields[1], 10, 64)
				if err != nil {
					return sysFreeMemory()
				}
				// Value in /proc/meminfo is in kB
				return val * 1024
			}
		}
	}

	// Fall back to Freeram if MemAvailable is not found
	return sysFreeMemory()
}
