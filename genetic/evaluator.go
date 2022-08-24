package genetic

type Evaluator struct {
	Candidate string
}

func (e *Evaluator) Fittest(pop Population) Organism {
	return Organism{
		Chromosome: e.Candidate,
	}
}

func (e *Evaluator) Fitness(org Organism) float32 {
	return 1.0
}
