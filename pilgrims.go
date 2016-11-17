package pilgrims

import "log"

type Player struct {
	Name        string
	Color       int
	Settlements int
	Cities      int
	Roads       int
	Resources   []Resource
	Cards       []DevelopmentCards
}

type Settlement struct{}
type City struct{}
type Road struct{}

type Point int
type SpecialCards int
type DevelopmentCards int

type GameMode int

const (
	Beginners GameMode = iota
	Experienced
)

type Terrain int

const (
	Desert Terrain = iota
	Hills
	Pasture
	Moutains
	Fields
	Forest
)

type Resource int

const (
	Nothing Resource = iota
	Brick
	Wool
	Ore
	Grain
	Lumber
)

type Tile struct {
	Number  int
	Terrain Terrain
}

type Game struct {
	actions chan Action
}

func (g *Game) Join(*Player) error {
	return nil
}

func (g *Game) Start() {
	go func() {
		for a := range g.actions {
			log.Println(a)
		}
	}()
}

type Action struct{}
