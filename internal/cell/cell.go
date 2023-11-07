package cell

import (
	"math/rand"

	"github.com/ReidMason/wave-function-collapse/internal/socket"
	"github.com/ReidMason/wave-function-collapse/internal/tile"
)

type Cell struct {
	North         *Cell
	South         *Cell
	East          *Cell
	West          *Cell
	r             *rand.Rand
	possibilities []*tile.Tile
	Tile          tile.Tile
	entropy       int
	collapsed     bool
}

func New(r *rand.Rand, possibilities []*tile.Tile) *Cell {

	cell := Cell{
		Tile:      tile.Blank,
		collapsed: false,
		r:         r,
	}

	cell.setPossibilities(possibilities)

	return &cell
}

func (c *Cell) setPossibilities(possibilities []*tile.Tile) {
	c.possibilities = possibilities
	c.entropy = len(c.possibilities)
}

func (c Cell) Entropy() int    { return c.entropy }
func (c Cell) Collapsed() bool { return c.collapsed }

func (c *Cell) Collapse() {
	// idx := c.r.Intn(c.entropy)
	idx := c.r.Intn(len(c.possibilities))
	c.Tile = *c.possibilities[idx]

	c.setPossibilities([]*tile.Tile{&c.Tile})
	c.collapsed = true
	c.North.constrain(c, 0, 2)
	c.South.constrain(c, 2, 0)
	c.East.constrain(c, 1, 3)
	c.West.constrain(c, 3, 1)
}

func (c *Cell) constrain(neighbour *Cell, dir1, dir2 int) {
	if c == nil || c.Collapsed() || neighbour == nil {
		return
	}

	neighbourPossibilities := neighbour.possibilities
	if len(neighbourPossibilities) == 0 {
		return
	}

	constrained := false
	for _, possiblity := range c.possibilities {
		canConnect := false
		for _, neighbourPossibility := range neighbourPossibilities {
			if socket.CanConnect(possiblity.Sockets[dir2], neighbourPossibility.Sockets[dir1]) {
				canConnect = true
				break
			}
		}

		if !canConnect {
			c.filterPossibilties(possiblity)
			constrained = true
		}
	}

	if constrained {
		c.North.constrain(c, 0, 2)
		c.South.constrain(c, 2, 0)
		c.East.constrain(c, 1, 3)
		c.West.constrain(c, 3, 1)
	}
}

func (c *Cell) filterPossibilties(targetTile *tile.Tile) {
	newPossibilities := make([]*tile.Tile, 0, len(c.possibilities))
	for _, val := range c.possibilities {
		if val != targetTile {
			newPossibilities = append(newPossibilities, val)
		}
	}
	c.setPossibilities(newPossibilities)
}
