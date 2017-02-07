package main

import (
	"golds/oops"
	"fmt"
	"os"
	"path/filepath"
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

	fmt.Println(exists("." + string(filepath.Separator) + "verifier/DocumentVerifie1r.go"))
	createDirectory("." + string(filepath.Separator) + "verifier1")
}

func packageRun() {
	creature := oops.Creature{"Anuj", true}
	creature.Dump()
	runner := oops.Runner(creature)

	flyingCreature := oops.FlyingCreature{creature, false}
	flyingCreature.Dump()

	runner.Run()
}

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
