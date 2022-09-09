package main

import (
	"fmt"
	g "genetic/genetic"
)

func main() {

	candidate := "0101010101010101010101010101010101010101010101010101010101010101"

	eval := g.Evaluator{}
	eval.Load(candidate)

	pop := g.Population{}
	pop.Populate(50)

	solution := run(candidate, eval, pop)
	fmt.Println("candidate: ", eval.Candidate)
	fmt.Println("solution:  ", solution)
}

func run(cand string, ev g.Evaluator, pop g.Population) g.Organism {
	var s func(candidate string, pop *g.Population, generation int) g.Organism

	s = func(candidate string, pop *g.Population, generation int) g.Organism {
		fittest := ev.Fittest(*pop)
		fitness := ev.Fitness(fittest)

		fmt.Printf("generation %d: chromosome: %s fitness: %.1f\n", generation, fittest, fitness)

		if fitness >= 1.0 {
			return fittest
		}

		return s(candidate, pop.Evolve(true, ev), generation+1)
	}

	return s(cand, &pop, 1)
}
