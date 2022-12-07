package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawLines, _ := os.ReadFile("data.txt")
	lineArray := strings.Split(string(rawLines), "\n")
	root := ParseLines(lineArray)

	// start := int64(0)
	// fmt.Println(TotalUnderLimit(&root, int64(100000), &start))

	limit := 30000000 - (70000000 - root.totalSize)
	fmt.Println(limit)
	out := make(map[string]int64)
	DirsOverLimit(&root, int64(limit), out)

	fmt.Println(Min(out))
}

type File struct {
	name string
	size int64
}

type Directory struct {
	name      string
	parent    *Directory
	dirs      map[string]*Directory
	files     map[string]*File
	fileSize  int64
	dirSize   int64
	totalSize int64
}

func ParseLines(lines []string) Directory {
	var root *Directory = &Directory{name: "/", dirs: make(map[string]*Directory), files: make(map[string]*File)}
	var currentDir *Directory = nil
	for i := 0; i < len(lines); i++ {
		command := strings.Split(lines[i], " ")
		if command[0] == "$" {
			switch command[1] {
			case "cd":
				if command[2] == "/" {
					currentDir = root
				} else if command[2] == ".." && currentDir != root {
					currentDir = currentDir.parent
				} else {
					currentDir = currentDir.dirs[command[2]]
				}
			case "ls":
				for i += 1; i < len(lines); i++ {
					line := lines[i]
					if line[0] == '$' {
						i -= 1
						break
					}
					output := strings.Split(line, " ")
					if output[0] == "dir" {
						currentDir.dirs[output[1]] = &Directory{name: output[1], parent: currentDir, dirs: make(map[string]*Directory), files: make(map[string]*File)}
					} else {
						sizei, _ := strconv.ParseInt(output[0], 10, 64)
						currentDir.files[output[1]] = &File{name: output[1], size: sizei}
					}
				}
			}
		}
	}

	populateSize(root)

	return *root
}

func populateSize(dir *Directory) int64 {
	var fileSize int64 = 0
	var dirSize int64 = 0
	for _, f := range dir.files {
		fileSize += f.size
	}
	for _, d := range dir.dirs {
		dirSize += populateSize(d)
	}
	dir.fileSize = fileSize
	dir.dirSize = dirSize
	dir.totalSize = fileSize + dirSize
	return dir.totalSize
}

func TotalUnderLimit(dir *Directory, limit int64, start *int64) int64 {
	if dir.totalSize <= limit {
		*start += dir.totalSize
	}
	for _, d := range dir.dirs {
		TotalUnderLimit(d, limit, start)
	}
	return *start
}

func DirsOverLimit(dir *Directory, limit int64, output map[string]int64) {
	if dir.totalSize >= limit {
		output[dir.name] = dir.totalSize
		fmt.Println(dir.name, dir.totalSize)
	}
	for _, d := range dir.dirs {
		DirsOverLimit(d, limit, output)
	}
}

func Min(input map[string]int64) string {
	max := int64(300000000)
	name := ""
	for n, s := range input {
		if s <= max {
			max = s
			name = n
		}
	}
	return name
}
