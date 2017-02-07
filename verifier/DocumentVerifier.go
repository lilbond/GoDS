package verifier

import (
  "golds/verifier/entities"
)

func Verify(d entities.Document) bool{
  if (d.Value == 10) {
    return true
  }

  return false
}
