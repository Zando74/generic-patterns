package main

import (
	"fmt"
	"generic-patterns/structural"
)

type Searcher interface {
	search(keyword string)
}

type File struct {
	structural.Composable
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

type Folder struct {
	structural.Composable
	structural.Composite[Searcher]
	name string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.Children {
		composite.search(keyword)
	}
}

func NewFolder(name string) *Folder {
	return &Folder{name: name}
}

func NewFile(name string) *File {
	return &File{name: name}
}

func MainCompositeExample() {
	folder1 := NewFolder("Folder 1")
	folder2 := NewFolder("Folder 2")
	file1 := NewFile("File 1")
	file2 := NewFile("File 2")

	folder1.Add(file1)
	folder1.Add(file2)
	folder1.Add(folder2)

	folder1.search("rose")
}
