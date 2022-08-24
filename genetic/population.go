package genetic

import "crypto/rand"

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

func (p *Population) Evolve(elitist bool, eval Evaluator) *Population {
	return p // todo
}
