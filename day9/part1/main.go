package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
  EMPTY = "."
)

func main() {
	diskRep := MustGetInput()
  disk := MustGetDisk(diskRep)
  firstEmptyIdx := seekToNextEmpty(disk, -1)
  lastFileByteIdx := seekToPrevFile(disk, len(disk))
  for ; firstEmptyIdx < lastFileByteIdx; {
    disk[firstEmptyIdx] = disk[lastFileByteIdx]
    disk[lastFileByteIdx] = EMPTY
    firstEmptyIdx = seekToNextEmpty(disk, firstEmptyIdx)
    lastFileByteIdx = seekToPrevFile(disk, lastFileByteIdx)
  }
  fmt.Printf("disk: %v\n", disk)
  // fmt.Printf("first empty: %d, last file byte idx: %d\n", firstEmptyIdx, lastFileByteIdx)
  checksum := 0
  for idx, item := range disk {
    if item == EMPTY {
      continue
    }
    itemInt, err := strconv.Atoi(item)
    if err != nil {
      panic(err)
    }
    checksum += (itemInt * idx)
  }
  fmt.Printf("checksum -> %d\n", checksum)
}

func seekToNextEmpty(disk []string, i int) int {
  r, n := i + 1, len(disk)
  for {
    if r >= n {
      return -1
    }
    if disk[r] == EMPTY {
      return r
    }
    r++
  }
}

func seekToPrevFile(disk []string, i int) int {
  r := i - 1
  for {
    if r < 0 {
      return -1
    }
    if disk[r] != EMPTY {
      return r
    }
    r--
  }
}

// Given the compressed representation of the disk, returns the actual disk.
func MustGetDisk(diskRep []string) []string {
  r := []string{}
  currFileId := 0
  for idx, ch := range diskRep {
    if ch == "\n" {
      break
    }
    if idx % 2 == 0 {
      occBlocks, err := strconv.Atoi(ch)
      if err != nil {
        panic(fmt.Errorf("failed to get occupied blocks: %w", err))
      }
      strCurrFileId := strconv.Itoa(currFileId)
      for range occBlocks {
        r = append(r, strCurrFileId)
      }
      currFileId++
    } else {
      // Free space.
      freeBlocks, err := strconv.Atoi(ch)
      if err != nil {
        panic(fmt.Errorf("failed to get free blocks: %w", err))
      }
      r = append(r,strings.Split(strings.Repeat(EMPTY, freeBlocks), "")...)
    }
  }
  return r
}

func MustGetInput() []string {
	contents, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
  return strings.Split(string(contents), "")
}
