package sach

import (
  "fmt"
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
  fmt.Println(data)
}

func (t topic) Read() []byte {
  return []byte("hello")
}
