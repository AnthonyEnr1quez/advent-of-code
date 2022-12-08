package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	totalDiskSize        = 70000000
	unusedDiskSizeNeeded = 30000000
)

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rootDir := Directory{Name: "/"}
	var currentDir *Directory

	var cmdBase string
	for scanner.Scan() {
		input := scanner.Text()
		firstChar := input[:1]

		if firstChar == "$" {
			fullCommand := input[2:]
			cmd := strings.Split(fullCommand, " ")

			cmdBase = cmd[0]

			if cmdBase == "cd" {
				switch cmd[1] {
				case "/":
					currentDir = &rootDir
				case "..":
					currentDir = currentDir.ParentDir
					fmt.Print()
				default:
					for _, dir := range currentDir.Directories {
						if cmd[1] == dir.Name {
							currentDir = dir
							break
						}
					}
				}
			}
		} else {
			file := strings.Split(input, " ")
			if file[0] == "dir" {
				currentDir.Directories = append(currentDir.Directories, &Directory{Name: file[1], ParentDir: currentDir})
			} else if size, err := strconv.Atoi(file[0]); err == nil {
				currentDir.Files = append(currentDir.Files, &File{Name: file[1], Size: size})
			}
		}
	}

	validDeletionDirs := rootDir.FindValidDeletionDirs()
	sizeOfValidDeletionDirs := 0
	for _, dir := range validDeletionDirs {
		sizeOfValidDeletionDirs += dir.Size()
	}

	sizeStillNeededToFree := unusedDiskSizeNeeded - (totalDiskSize - rootDir.Size())

	var dirToDelete *Directory
	for _, child := range rootDir.AllChildren() {
		if dirToDelete == nil && child.Size() >= sizeStillNeededToFree {
			dirToDelete = child
		} else if child.Size() >= sizeStillNeededToFree && child.Size() < dirToDelete.Size() {
			dirToDelete = child
		}
	}

	fmt.Println("Sum of directory sizes with <= 100000:", sizeOfValidDeletionDirs)
	fmt.Println("Size of directory to delete:", dirToDelete.Size())
}

type File struct {
	Name string
	Size int
}

type Directory struct {
	Name        string
	Files       []*File
	Directories []*Directory
	ParentDir   *Directory
}

func (d *Directory) Size() (size int) {
	for _, file := range d.Files {
		size += file.Size
	}
	for _, dir := range d.Directories {
		size += dir.Size()
	}
	return
}

func (d *Directory) FindValidDeletionDirs() (dirs []*Directory) {
	for _, dir := range d.AllChildren() {
		if dir.Size() <= 100000 {
			dirs = append(dirs, dir)
		}
	}
	return
}

func (d *Directory) AllChildren() (dirs []*Directory) {
	for _, dir := range d.Directories {
		if len(dir.Directories) != 0 {
			dirs = append(dirs, dir.AllChildren()...)
		}
		dirs = append(dirs, dir)
	}
	return
}
