package genetic

import "strings"

type Organism struct {
	Genes []byte
}

func (o Organism) String() string {
	var sb strings.Builder

	for _, g := range o.Genes {
		if g == 1 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}

	return sb.String()
}
