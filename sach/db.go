package sach

import (
  "fmt"
  "os"
  "path/filepath"
)

var (
  topics []*topic
)

type Config struct {
    DataDirectory string
}

type Sach struct {
  SachConfig Config
}

//a topic name should not have spaces
//may not work on all systens path should be
//data directory could have / at end
func (sach *Sach) GetTopicHandler(name string) TopicHandler{
  path := sach.SachConfig.DataDirectory + "/" + name
  createPath(path)

  topicHandler := newTopicHandler(name, path)
  return topicHandler
}

func newTopicHandler(name string, path string) TopicHandler {
  topicHandler := topic{name, path}
  return topicHandler
}

func createPath(path string) {
  if !exists(path) {
    pathErr := os.MkdirAll(path,0777)

    // REVIEW: panic
    if pathErr != nil {
      fmt.Println(pathErr)
    }
  }

  filePath := path + string(filepath.Separator) + "data.sach"
  if !exists(filePath) {
    f, err := os.Create(filePath)
    defer f.Close()

    //REVIEW: panic
    if err != nil {
      fmt.Println(err)
    }
  }
}

func exists(filePath string) (exists bool) {
  exists = true

  if _, err := os.Stat(filePath); os.IsNotExist(err) {
    exists = false
  }

  return
}
