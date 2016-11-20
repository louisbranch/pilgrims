package pilgrims

import (
	"fmt"
	"math/rand"
	"time"
)

type Board [19]Tile

func NewBoard() (b Board) {
	const a = 97

	terrains := terrains()
	chits := chits()
	c := 0

	for i, t := range b {
		t.Letter = byte(a + i)
		t.Terrain = terrains[i]
		if t.Terrain != Desert {
			t.Chit = chits[c]
			c++
		} else {
			t.Chit = 7
		}
		b[i] = t
	}

	return
}

func (b Board) String() string {
	acc := ""

	for _, t := range b {
		acc += fmt.Sprintf("%c: %d - %s | ", t.Letter, t.Chit, terrain(t.Terrain))
	}

	return acc
}

func terrain(t byte) string {
	switch t {
	case Desert:
		return "Desert"
	case Hills:
		return "Hills"
	case Pasture:
		return "Pasture"
	case Mountains:
		return "Mountains"
	case Fields:
		return "Fields"
	case Forest:
		return "Forest"
	}

	return "Unknown"
}

func chits() [18]byte {
	const size = 18
	rand.Seed(time.Now().UTC().UnixNano())
	n := [size]byte{2, 3, 3, 4, 4, 5, 5, 6, 6, 8, 8, 9, 9, 10, 10, 11, 11, 12}
	m := [size]byte{}
	perm := rand.Perm(size)

	for i := range m {
		m[i] = n[perm[i]]
	}

	return m
}

func terrains() [19]byte {
	const size = 19
	rand.Seed(time.Now().UTC().UnixNano())
	n := [size]byte{
		Desert,
		Hills, Hills, Hills,
		Mountains, Mountains, Mountains,
		Pasture, Pasture, Pasture, Pasture,
		Fields, Fields, Fields, Fields,
		Forest, Forest, Forest, Forest,
	}

	m := [size]byte{}
	perm := rand.Perm(size)

	for i := range m {
		m[i] = n[perm[i]]
	}

	return m
}
