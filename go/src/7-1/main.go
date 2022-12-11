package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Plan:
// build up a representation of a file system
// it is a tree type data structure made up of directories
// each directory can point to one or more child directories
// the file system will need to keep track of its root directory
// to navigate down the tree, each directory will need to know its children.
// will have to use recursion to search the tree from its root
// each directory can contain 0 or more files.
// an empty directory is size 0.
// a directory containing files has size equal to the sum of its files
// in a single directory, two files cant have the file name + file extension
// in the file system, two directories with the same name can exist in different locations
// separate the population of the file system with data into a constructor function
// once the file system is constructed, use recursion to print it so can debug any issue

func getInput() string {
	file, err := os.Open("../../input/7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	return string(data)
}

func main() {
	i := getInput()

	fs := buildFilesystem(i)

	directories := fs.findDirectories()

	total := 0
	for _, dir := range directories {
		total += dir.size()
	}

	fmt.Println(total)
}

type directory struct {
	totalSize int
	name      string
	children  []*directory
	files     map[string]int // map of file name to size
}

// String is a custom stringer to print a representation of a directory.
func (d directory) String() string {
	dirs := []*directory{&d}
	indent := 0
	printDirectories(dirs, indent)
	return ""
}

// printDirectories is a recursive function that prints all directories and children.
func printDirectories(directories []*directory, indent int) {
	for _, dir := range directories {
		fmt.Printf("%s- %s (dir, size=%d)\n", space(indent), dir.name, dir.size())
		indent++
		for name, size := range dir.files {
			fmt.Printf("%s- %s (file, size=%d)\n", space(indent), name, size)
		}
		printDirectories(dir.children, indent)
	}
}

// space is a helper that returns a blank string n chars long.
func space(n int) string {
	return strings.Repeat(" ", n)
}

// size calculates the size of a directory including child directories
func (d directory) size() int {
	// if the total size has already been calculated
	if d.totalSize < 0 {
		return d.totalSize // return it.
	}

	// otherwise calculate size:
	// size includes all files in this directory
	for _, filesize := range d.files {
		d.totalSize += filesize
	}

	// plus the size of all child directories.
	for _, child := range d.children {
		d.totalSize += child.size()
	}

	return d.totalSize
}

func (d directory) searchForUnderMax(max int) []*directory {
	result := []*directory{}

	if d.size() < max {
		result = append(result, &d)
	}

	for _, child := range d.children {
		result = append(result, child.searchForUnderMax(max)...)
	}

	return result
}

func (d *directory) getChild(name string) (*directory, error) {
	for _, child := range d.children {
		if child.name == name {
			return child, nil
		}
	}
	return nil, errors.New("not found")
}

func (d *directory) addFile(name string, size int) {
	d.files[name] = size
}

func (d *directory) addChild(name string) {
	child := newDirectory(name)
	d.children = append(d.children, child)
}

func newDirectory(name string) *directory {
	return &directory{
		name:     name,
		children: []*directory{},
		files:    map[string]int{},
	}
}

type filesystem struct {
	root *directory
	path []*directory
}

func (f filesystem) String() string {
	return f.root.String()
}

// The working directory is the last directory in the path.
func (f *filesystem) getWorkingDir() *directory {
	return f.path[len(f.path)-1]
}

func (f *filesystem) createFile(name string, size int) {
	f.getWorkingDir().addFile(name, size)
}

func (f *filesystem) createDirectory(name string) {
	f.getWorkingDir().addChild(name)
}

func (f *filesystem) moveUpPath() {
	// to move up the path, the last element of the path is removed.
	f.path = f.path[:len(f.path)-1]
}

func (f *filesystem) moveTo(target string) {
	targetDir, err := f.getWorkingDir().getChild(target)
	if err != nil {
		log.Fatalf("unable to cd to directory %s, error %s\n", target, err.Error())
	}
	f.path = append(f.path, targetDir)
}

// newFilesystem is a constructor for a filesystem.
// it creates the root directory and sets the working directory to root.
func newFilesystem() filesystem {
	root := newDirectory("/")

	return filesystem{
		root: root,
		path: []*directory{root},
	}
}

func (f *filesystem) findDirectories() []*directory {
	dirs := []*directory{}
	max := 100000

	if f.root.size() < max {
		dirs = append(dirs, f.root)
	}

	dirs = append(dirs, f.root.searchForUnderMax(max)...)
	return dirs
}

// buildFilesystem initialises a new filesystem with data from input.
func buildFilesystem(input string) filesystem {
	scanner := bufio.NewScanner(strings.NewReader(input))

	filesystem := newFilesystem()

	scanner.Scan() // first line is useless info, scan it to skip it.
	for scanner.Scan() {
		line := scanner.Text()

		if line == "$ ls" {
			// this is the command to list files in a directory
			continue
		}

		if strings.Contains(line, "$ cd") {
			parts := strings.Split(line, " ")
			target := parts[2]

			if target == ".." {
				// this is the command to move up one directory
				filesystem.moveUpPath()

			} else {
				// this is the command to move into a specific directory
				filesystem.moveTo(target)
			}

			continue
		}

		if strings.HasPrefix(line, "dir") {
			// this is the command to create a directory
			parts := strings.Split(line, " ")
			name := parts[1]
			filesystem.createDirectory(name)
			continue
		}

		// if none of the other commands have been matched,
		// then a file is being created.
		parts := strings.Split(line, " ")

		size, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("unable to parse int, error %s\n", err.Error())
		}
		name := parts[1]
		filesystem.createFile(name, size)
	}
	return filesystem
}
