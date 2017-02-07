package tests

import (
  "testing"
)

func TestTruth(t *testing.T) {
  if true != true {
    t.Error("Everything I know is wrong")
  }
}
