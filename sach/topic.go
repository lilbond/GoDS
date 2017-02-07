package sach

import (
  "fmt"
  "io/ioutil"
  "os"
)

type TopicHandler interface {
  Write(data []byte)
  Read() []byte
}

type topic struct {
  Name string
  Path string
}

func (t topic) Write(data []byte) {
  for index := 0; index < 10; index++ {
    err := ioutil.WriteFile(t.Path, data, os.FileMode.ModeAppend)
    if err != nil {
      fmt.Println(err)
    }
  }
}

func (t topic) Read() []byte {
  return []byte("hello")
}
