package genetic

type Population struct {
	Size int
}

func (p *Population) Evolve(elitist bool, eval Evaluator) Population {
	return *p // todo
}
