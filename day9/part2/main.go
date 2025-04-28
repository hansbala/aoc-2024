package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type FileInfo struct {
	id  int
	startIdx int
	endIdx int
}

type FreeSpace struct {
	startIdx int
	endIdx   int
}

type Disk struct {
	files     []*FileInfo
	freeSpaces []*FreeSpace
}

func main() {
	diskInput := MustGetDiskInput()
	disk := MustGetDisk(diskInput)
	// Sort the files DESC `id`.
	slices.SortFunc(disk.files, func(f1 *FileInfo, f2 *FileInfo) int {
		return f2.id - f1.id
	})
	// Sort the freespace ASC `startIdx`.
	slices.SortFunc(disk.freeSpaces, func(f1 *FreeSpace, f2 *FreeSpace) int {
		return f1.startIdx - f2.startIdx
	})
	// PrintDisk(&disk)
	for _, file := range disk.files {
		fileLen := file.endIdx - file.startIdx
		for _, freeSpace := range disk.freeSpaces {
			if freeSpace.startIdx == -1 || freeSpace.endIdx == -1 {
				continue
			}
			if freeSpace.startIdx > file.startIdx {
				// We don't want to move the file right only left.
				continue
			}
			freeSpaceLen := freeSpace.endIdx - freeSpace.startIdx
			if fileLen > freeSpaceLen {
				continue
			}
			// We can place the file in here.
			file.startIdx = freeSpace.startIdx
			file.endIdx = file.startIdx + fileLen
			freeSpace.startIdx = freeSpace.startIdx + fileLen + 1
			if freeSpace.startIdx > freeSpace.endIdx {
				freeSpace.startIdx = -1
				freeSpace.endIdx = -1
			}
		}
	}
	fmt.Printf("checksum: %d\n", ComputeChecksum(GetDiskRep(&disk)))
}

func GetDiskRep(disk *Disk) []string {
	diskSize := -1
	for _, file := range disk.files {
		diskSize = int(math.Max(float64(diskSize), float64(file.endIdx)))
	}
	diskSize++
	r := []string{}
	for range diskSize {
		r = append(r, ".")
	}
	for _, file := range disk.files {
		fileIdStr := strconv.Itoa(file.id)
		for i := file.startIdx; i <= file.endIdx; i++ {
			r[i] = fileIdStr
		}
	}
	return r
}

func PrintDisk(disk *Disk) {
	for _, file := range disk.files {
		fmt.Printf("(%d, %d, %d) ", file.id, file.startIdx, file.endIdx)
	}
	fmt.Println()
	for _, freeSpace := range disk.freeSpaces {
		fmt.Printf("(%d, %d) ", freeSpace.startIdx, freeSpace.endIdx)
	}
	fmt.Println()
}

func MustGetDisk(inp string) Disk {
	r := Disk{files: []*FileInfo{}, freeSpaces: []*FreeSpace{}}
	input := strings.Split(inp, "")
	fileId := 0
	diskCursor := 0
	for i, ch := range input {
		if ch == "\n" {
			continue
		}
		if i%2 == 0 {
			// This is a file.
			fileLen, err := strconv.Atoi(ch)
			if err != nil {
				panic(err)
			}
			r.files = append(r.files, &FileInfo{id: fileId, startIdx: diskCursor, endIdx: diskCursor + fileLen - 1})
			fileId++
			diskCursor += fileLen
		} else {
			// This is amount of free space.
			freeSpaceLen, err := strconv.Atoi(ch)
			if err != nil {
				panic(err)
			}
			r.freeSpaces = append(r.freeSpaces, &FreeSpace{startIdx: diskCursor, endIdx: diskCursor + freeSpaceLen - 1})
			diskCursor += freeSpaceLen
		}
	}
	return r
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

func MustGetDiskInput() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(content)
}
