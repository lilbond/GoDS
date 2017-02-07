package sach

import (
  "fmt"
  "testing"
)

func TestDb(t *testing.T) {
  config := Config{"/Users/ayadav/Documents/comparepy/data"}
  sachdb := Sach{config}

  topicHandler := sachdb.GetTopicHandler("my_topic")
  topicHandler.Write([]byte("hello"))
  fmt.Println(topicHandler.Read())
}
