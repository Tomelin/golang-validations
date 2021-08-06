package validations

import "github.com/paemuri/brdoc"

func IsCEPValid(doc string) bool {
	return brdoc.IsCEP(doc)
}
