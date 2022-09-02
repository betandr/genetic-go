package genetic

import (
	"crypto/rand"
	"fmt"
	"strings"
)

type Population struct {
	Organisms []Organism
}

var chromosomeSize int = 64
var mutationRate float32 = 0.015
var mixingRatio float32 = 0.5

// Populate creates a population of Organisms with random Genes
func (p *Population) Populate(size int) {
	p.Organisms = make([]Organism, size)

	for i := range p.Organisms {
		genes := make([]byte, chromosomeSize)
		rand.Read(genes) // randomise bytes
		p.Organisms[i] = Organism{Genes: genes}
	}
}

func (p *Population) String() string {
	var sb strings.Builder

	for _, o := range p.Organisms {
		sb.WriteString(o.String())
		sb.WriteString(", ")
	}
	return fmt.Sprintf("[%s]", strings.TrimRight(sb.String(), ", "))
}

func (p *Population) AddOrganism(index int, org Organism) {
	p.Organisms[index] = org
}

func (p *Population) Evolve(elitist bool, eval Evaluator) *Population {
	return p // todo
}

func (p *Population) mutate(org Organism) Organism {
	return org // todo
}

func (p *Population) crossover(parent1 Organism, parent2 Organism) Organism {
	return Organism{}
}
