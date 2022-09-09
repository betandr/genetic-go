package genetic

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

type Population struct {
	Organisms []Organism
}

var chromosomeSize int = 64
var mutationRate float64 = 0.015
var mixingRatio float64 = 0.5

// Populate creates a population of Organisms with random Genes
func (p *Population) Populate(size int) {
	p.Organisms = make([]Organism, size)

	for i := range p.Organisms {
		genes := make([]byte, chromosomeSize)

		for j := 0; j < chromosomeSize; j++ {
			genes[j] = byte(math.Round(rand.Float64()))
		}

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

// Evolve the population using crossover and mutation
func (p *Population) Evolve(elitist bool, ev Evaluator) *Population {
	nextGeneration := Population{}
	nextGeneration.Populate(len(p.Organisms))

	offset := 0

	if elitist {
		eliteOrganism := ev.Fittest(*p)
		nextGeneration.AddOrganism(0, p.mutate(eliteOrganism))
		offset++
	}

	for i := offset; i < len(p.Organisms); i++ {
		parent1 := p.selectOrganism(ev)
		parent2 := p.selectOrganism(ev)
		child := p.crossover(parent1, parent2)

		nextGeneration.AddOrganism(i, p.mutate(child))
	}

	return &nextGeneration
}

// mutate an organism with a random rate of 0.015
func (p *Population) mutate(org Organism) Organism {
	for i := range org.Genes {
		if rand.Float64() <= mutationRate {
			org.Genes[i] = byte(math.Round(rand.Float64()))
		}
	}
	return org
}

// crossover creates a child organism from two parents
func (p *Population) crossover(parent1 Organism, parent2 Organism) Organism {
	chromosomes := make([]byte, len(parent2.Genes))

	for i, gene := range parent1.Genes {
		if rand.Float64() <= mixingRatio {
			chromosomes[i] = gene
		} else {
			chromosomes[i] = parent2.Genes[i]
		}
	}

	return Organism{chromosomes}
}

// selectOrganism from the population using stochastic universal sampling
func (p *Population) selectOrganism(eval Evaluator) Organism {
	numRounds := 10
	populationSize := len(p.Organisms)
	tournament := Population{}
	tournament.Populate(populationSize)

	for i := 0; i < numRounds; i++ {
		randomOrganism := p.Organisms[rand.Intn(populationSize)]
		tournament.AddOrganism(i, randomOrganism)
	}

	return eval.Fittest(tournament)
}
