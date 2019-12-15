// Package namegen generates randomized comet names
package namegen

import (
	"fmt"
	"math/rand"
	"time"
)

const nameFormat string = "%s%s%s"

// GenerateName returns a random name in form <adj><adj><noun>
func GenerateName() string {
	name := fmt.Sprintf(nameFormat,
		getRandomAdjective(),
		getRandomAdjective(),
		getRandomNoun())

	return name
}

func getRandomAdjective() string {
	rand.Seed(time.Now().UnixNano())
	return adjectives[rand.Intn(len(adjectives))]
}

func getRandomNoun() string {
	rand.Seed(time.Now().UnixNano())
	return nouns[rand.Intn(len(nouns))]
}
