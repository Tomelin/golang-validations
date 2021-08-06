package validations

import "github.com/paemuri/brdoc"

func IsCPFCNPJValid(doc string) bool {

	return brdoc.IsCPF(doc)
}
