package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
)

func dirTree(out io.Writer, path string, files bool) error {
	printTree(out, path, files, []bool{})
	return nil
}

func printTree(out io.Writer, path string, areFilesPrinted bool, level []bool)  {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err.Error())
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	if !areFilesPrinted {
		withoutFiles := files[:0]
		for _, el := range files{
			if el.IsDir() {
				withoutFiles = append(withoutFiles, el)
			}
		}
		files = withoutFiles
	}
	if len(files) == 0 {
		return
	}
	tabString := getSpaceString(level)

	for _, file := range files[:len(files) - 1] {
		printLine(out, file, "├", tabString)
		if file.IsDir() {
			printTree(out, path + string(os.PathSeparator) + file.Name(), areFilesPrinted, append(level, true))
		}
	}

	file := files[len(files) - 1]
	printLine(out, file, "└", tabString)
	if file.IsDir() {
		printTree(out, path + string(os.PathSeparator) + file.Name(), areFilesPrinted, append(level, false))
	}
}


func getSpaceString(levels []bool) string {
	returnValue := ""
	for _, l := range levels {
		if l {
			returnValue += "│\t"
		} else {
			returnValue += "\t"
		}
	}
	return returnValue
}

func printLine(out io.Writer, file os.FileInfo, char string, spaceBeforeName string)  {
	strInfo := fmt.Sprint(file.Size())
	if strInfo == "0"{
		strInfo = "empty"
	} else {
		strInfo += "b"
	}
	if file.IsDir(){
		out.Write([]byte(spaceBeforeName + char + "───" + file.Name() + "\n"))
	} else {
		out.Write([]byte(spaceBeforeName + char + "───" + file.Name() + " (" + strInfo  + ")" + "\n"))
	}
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		fmt.Print("Bye")
		panic(err.Error())
	}
}


