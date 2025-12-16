// Package memory provides a single method reporting total system memory
// accessible to the kernel.
package memory

// TotalMemory returns the total accessible system memory in bytes.
//
// The total accessible memory is installed physical memory size minus reserved
// areas for the kernel and hardware, if such reservations are reported by
// the operating system.
//
// If accessible memory size could not be determined, then 0 is returned.
func TotalMemory() uint64 {
	return sysTotalMemory()
}

// FreeMemory returns the total free system memory in bytes.
//
// The total free memory is installed physical memory size minus reserved
// areas for other applications running on the same system.
//
// If free memory size could not be determined, then 0 is returned.
func FreeMemory() uint64 {
	return sysFreeMemory()
}

// AvailableMemory returns the available system memory in bytes.
//
// Available memory is an estimate of how much memory is available for
// starting new applications, without swapping. It accounts for reclaimable
// memory such as buffers and cache that can be freed if needed.
//
// On Linux, this reads MemAvailable from /proc/meminfo (available since
// kernel 3.14). On other systems, it may fall back to FreeMemory.
//
// If available memory size could not be determined, then 0 is returned.
func AvailableMemory() uint64 {
	return sysAvailableMemory()
}
