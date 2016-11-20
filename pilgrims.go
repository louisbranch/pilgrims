package pilgrims

type Players [4]Player

type GameState struct {
	Players
	Cards
	Board
}

type Cards struct {
	Resources
	DevelopmentCards
	SpecialCards
}

type SpecialCards struct {
	LongestRoad byte
	LargestArmy byte
}

type Resources struct {
	Brick  byte
	Wool   byte
	Ore    byte
	Grain  byte
	Lumber byte
}

type DevelopmentCards struct {
	Knight       byte
	VictoryPoint byte
	RoadBuild    byte
	Monopoly     byte
	YearOfPlenty byte
}

type Player struct {
	Name         [10]byte
	Settlements  byte
	Cities       byte
	Roads        byte
	Resources    Resources
	VisibleCards DevelopmentCards
	HiddenCards  DevelopmentCards
}

const (
	Desert byte = iota
	Hills
	Pasture
	Mountains
	Fields
	Forest
)

const (
	Nothing byte = iota
	Brick
	Wool
	Ore
	Grain
	Lumber
)

type Tile struct {
	Letter   byte
	Chit     byte
	Terrain  byte
	Edges    [6]byte
	Vertices [6]byte
}
