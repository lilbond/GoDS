package main

import (
	"golds/oops"
	"fmt"
	"os"
	"path/filepath"
	"io"
)

type User struct {
	Name string
}

func (u *User) mutate() {
	fmt.Println(u)
	u.Name = "Anuj"
}

func main() {
	u := &User{"Bhupesh"}
	fmt.Println(u)
	u.mutate()
	fmt.Println(u.Name)

	//createDirectory("." + string(filepath.Separator) + "verifier")
	//createFile("." + string(filepath.Separator) + "verifier/DocumentVerifier.go")
	read("." + string(filepath.Separator) + "verifier/tests/DocumentVerifier_test.go")
	write("." + string(filepath.Separator) + "verifier/tests/DocumentVerifier_test.go", "\n")
	write("." + string(filepath.Separator) + "verifier/tests/DocumentVerifier_test.go", "here comes my line\n")
}

func packageRun() {
	creature := oops.Creature{"Anuj", true}
	creature.Dump()
	runner := oops.Runner(creature)

	flyingCreature := oops.FlyingCreature{creature, false}
	flyingCreature.Dump()

	runner.Run()
}

//file creation
func exists(filePath string) (exists bool) {
  exists = true

  if _, err := os.Stat(filePath); os.IsNotExist(err) {
    exists = false
  }

  return
}

func createDirectory(directoryPath string) {
	//choose your permissions well
	pathErr := os.MkdirAll(directoryPath,0777)

	//check if you need to panic, fallback or report
	if pathErr != nil {
		fmt.Println(pathErr)
	}
}

func createFile(filePath string) {
	file, err := os.Create(filePath)
	defer file.Close()

	//REVIEW: panic
	if err != nil {
		fmt.Println(err)
	}
}

//file read-write
func read(filePath string) {
	f, ferr := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if ferr != nil {
		fmt.Println(ferr)
		return
	}

	defer f.Close()
	data := make([]byte, 4096)
	for {
		n, err := f.Read(data)
    if err != nil {
      if err == io.EOF {
          break
      }
      fmt.Println(err)
      return
    }
    data = data[:n]
		fmt.Println(data)
	}
}

func write(filePath string, content string) {
	f, ferr := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if ferr != nil {
		fmt.Println(ferr)
		return
	}
	defer f.Close()

	_, err := f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}
}
