package main

import (
	"fmt"
	"github.com/jbrukh/bayesian"
)

func main() {
	const (
		Good bayesian.Class = "Good"
		Bad  bayesian.Class = "Bad"
	)

	classifier := bayesian.NewClassifier(Good, Bad)
	goodStuff := []string{"tall", "rich", "handsome"}
	badStuff := []string{"poor", "smelly", "ugly"}
	classifier.Learn(goodStuff, Good)
	classifier.Learn(badStuff, Bad)

	scores, likely, _ := classifier.LogScores(
		[]string{"tall", "girl"},
	)
	fmt.Printf("Score: %v, likely: %v\n", scores, likely)

	probs, likely, _ := classifier.ProbScores(
		[]string{"tall", "girl"},
	)
	fmt.Printf("Probs: %v, likely: %v\n", probs, likely)

}
