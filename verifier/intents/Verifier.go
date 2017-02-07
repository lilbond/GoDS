package verifier

import (
  "golds/verifier/entities"
)
type Verifier interface {
  Verify(d entities.Document) bool
}
