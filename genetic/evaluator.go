package genetic

type Evaluator struct {
	Candidate []byte
	solution  []byte
}

// LoadSolution sets the solution to evaluate against
func (e *Evaluator) Load(s string) {
	e.solution = make([]byte, len(s))
	copy(e.solution[:], s)
}

// Fittest finds the most fit organism in a population
func (e *Evaluator) Fittest(pop Population) Organism {
	org := pop.Organisms[0]

	for i := 0; i < len(pop.Organisms); i++ {
		nextOrg := pop.Organisms[i]

		if e.Fitness(nextOrg) > e.Fitness(org) {
			org = pop.Organisms[i]
		}
	}

	return org
}

// Fitness calculates an organism's fitness by comparing it to the optimal solution
func (e *Evaluator) Fitness(org Organism) float32 {
	score := float32(0)

	for i, gene := range org.Genes {
		if e.solution[i] == gene {
			score = score + 1
		}
	}

	maxScore := len(org.Genes)
	return 1.0 - ((float32(maxScore) - float32(score)) / 100)
}
