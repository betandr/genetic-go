package main

import (
	"fmt"
	g "genetic/genetic"
)

func main() {

	eval := g.Evaluator{
		Candidate: "0101010101010101010101010101010101010101010101010101010101010101",
	}

	pop := g.Population{
		Size: 50,
	}

	solution := solve(eval, pop)
	fmt.Println("candidate: ", eval.Candidate)
	fmt.Println("solution:  ", solution)

}

func solve(evl g.Evaluator, pop g.Population) string {
	var s func(pop g.Population, generation int) string

	s = func(pop g.Population, generation int) string {
		fittest := evl.Fittest(pop)
		fitness := evl.Fitness(fittest)

		fmt.Printf("generation %d: chromosome: %s fitness: %.1f\n", generation, fittest.Chromosome, fitness)

		if fitness >= 1.0 {
			return string(fittest.Chromosome)
		}

		return s(pop.Evolve(true, evl), generation+1)
	}

	return s(pop, 1)
}
