package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type DISK_TYPE int

const (
	FREESPACE DISK_TYPE = iota
	FILE
)

func main() {
	fmt.Printf(
		"checksum: %d\n", 
		ComputeChecksum(MustMoveDisk(MustGetDisk(MustGetDiskInput()))),
	)
}

func ComputeChecksum(disk []string) int {
	r := 0
	for i, ch := range disk {
		if ch == "." {
			continue
		}
		fileId, err := strconv.Atoi(ch)
		if err != nil {
			panic(err)
		}
		r += (fileId * i)
	}
	return r
}

func MustMoveDisk(disk []string) []string {
	r := slices.Clone(disk)
	freeIdx := SeekAhead(disk, FREESPACE, 0)
	fileIdx := SeekBack(disk, FILE, len(disk))
	for ; freeIdx < fileIdx ; {
		r[freeIdx] = r[fileIdx]
		r[fileIdx] = "."
		// Now we move the pointers forward and back.
		if freeIdx = SeekAhead(disk, FREESPACE, freeIdx); freeIdx == -1 {
			panic("free idx was -1")
		}
		if fileIdx = SeekBack(disk, FILE, fileIdx); fileIdx == -1 {
			panic("file idx was -1")
		}
	}
	return r
}

// SeekAhead looks in (fromIdx, end]
func SeekAhead(disk []string, lookFor DISK_TYPE, fromIdx int) int {
	for i := fromIdx + 1; i < len(disk); i++ {
		if lookFor == FREESPACE && disk[i] == "." {
			return i
		}
		if lookFor == FILE && disk[i] != "." {
			return i
		}
	}
	return -1
}

// SeekBack looks back in (fromIdx ... start]
func SeekBack(disk []string, lookFor DISK_TYPE, fromIdx int) int {
	for i := fromIdx - 1; i >= 0; i-- {
		if lookFor == FREESPACE && disk[i] == "." {
			return i
		}
		if lookFor == FILE && disk[i] != "." {
			return i
		}
	}
	return -1
}

func MustGetDisk(inp string) []string {
	r := []string{}
	input := strings.Split(inp, "")
	fileId := 0
	for i, ch := range input {
		if ch == "\n" {
			continue
		}
		if i % 2 == 0 {
			// This is a file.
			fileLen, err := strconv.Atoi(ch)
			if err != nil {
				panic(err)
			}
			fileIdStr := strconv.Itoa(fileId)
			for range fileLen {
				r = append(r, fileIdStr)
			}
			fileId++
		} else {
			// This is amount of free space.
			freeSpaceLen, err := strconv.Atoi(ch)
			if err != nil {
				panic(err)
			}
			for range freeSpaceLen {
				r = append(r, ".")
			}
		}
	}
	return r
}

func MustGetDiskInput() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(content)
}
