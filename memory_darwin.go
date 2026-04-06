// +build darwin

package memory

func sysTotalMemory() uint64 {
	s, err := sysctlUint64("hw.memsize")
	if err != nil {
		return 0
	}
	return s
}

func sysFreeMemory() uint64 {
	pageSize, err := sysctlUint32("vm.pagesize")
	if err != nil || pageSize == 0 {
		pageSize = 16384
	}

	freePages, err := sysctlUint32("vm.page_free_count")
	if err != nil {
		return 0
	}

	return uint64(freePages) * uint64(pageSize)
}

func sysAvailableMemory() uint64 {
	pageSize, err := sysctlUint32("vm.pagesize")
	if err != nil || pageSize == 0 {
		pageSize = 16384
	}

	free, _ := sysctlUint32("vm.page_free_count")
	purgeable, _ := sysctlUint32("vm.page_purgeable_count")
	speculative, _ := sysctlUint32("vm.page_speculative_count")

	// Available memory includes free pages plus pages that can be reclaimed
	// without swapping (purgeable and speculative pages)
	return uint64(free+purgeable+speculative) * uint64(pageSize)
}
