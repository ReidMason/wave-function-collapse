package tile

import (
	"fmt"
	"strings"

	"github.com/ReidMason/wave-function-collapse/internal/socket"
)

type Tile struct {
	Style   string
	Sockets [4]socket.Socket
}

var Blank = Tile{
	Style: "bg-slate-500",
}

func GetAllTiles() []Tile {
	var cardinalTransition = "from-30% to-70%"
	var diagonalTransition = "from-60% to-90%"

	grass := Tile{
		Sockets: [4]socket.Socket{
			socket.Grass,
			socket.Grass,
			socket.Grass,
			socket.Grass,
		},
		Style: "bg-green-200",
	}

	forest := Tile{
		Sockets: [4]socket.Socket{
			socket.Forest,
			socket.Forest,
			socket.Forest,
			socket.Forest,
		},
		Style: "bg-green-400",
	}

	sand := Tile{
		Sockets: [4]socket.Socket{
			socket.Sand,
			socket.Sand,
			socket.Sand,
			socket.Sand,
		},
		Style: "bg-orange-200",
	}

	water := Tile{
		Sockets: [4]socket.Socket{
			socket.Water,
			socket.Water,
			socket.Water,
			socket.Water,
		},
		Style: "bg-blue-400",
	}

	waterSand := Tile{
		Sockets: [4]socket.Socket{
			socket.WaterT,
			socket.WaterSandE,
			socket.SandT,
			socket.WaterSandW,
		},
		Style: fmt.Sprintf("bg-gradient-to-t from-orange-200 to-blue-400 %s", cardinalTransition),
	}

	waterSandCorner := Tile{
		Sockets: [4]socket.Socket{
			socket.WaterSandCornerN,
			socket.SandT,
			socket.SandT,
			socket.WaterSandCornerW,
		},
		Style: fmt.Sprintf("bg-gradient-to-tl from-orange-200 to-blue-400 %s", diagonalTransition),
	}

	sandWaterCorner := Tile{
		Sockets: [4]socket.Socket{
			socket.SandWaterCornerN,
			socket.WaterT,
			socket.WaterT,
			socket.SandWaterCornerW,
		},
		Style: fmt.Sprintf("bg-gradient-to-tl from-blue-400 to-orange-200 %s", diagonalTransition),
	}

	sandGrass := Tile{
		Sockets: [4]socket.Socket{
			socket.SandT,
			socket.SandGrassE,
			socket.GrassT,
			socket.SandGrassW,
		},
		Style: fmt.Sprintf("bg-gradient-to-t from-green-200 to-orange-200 %s", cardinalTransition),
	}

	sandGrassCorner := Tile{
		Sockets: [4]socket.Socket{
			socket.SandGrassCornerN,
			socket.GrassT,
			socket.GrassT,
			socket.SandGrassCornerW,
		},
		Style: fmt.Sprintf("bg-gradient-to-tl from-green-200 to-orange-200 %s", diagonalTransition),
	}

	grassSandCorner := Tile{
		Sockets: [4]socket.Socket{
			socket.GrassSandCornerN,
			socket.SandT,
			socket.SandT,
			socket.GrassSandCornerW,
		},
		Style: fmt.Sprintf("bg-gradient-to-tl from-orange-200 to-green-200 %s", diagonalTransition),
	}

	forestGrass := Tile{
		Sockets: [4]socket.Socket{
			socket.ForestT,
			socket.ForestGrassE,
			socket.GrassT,
			socket.ForestGrassW,
		},
		Style: fmt.Sprintf("bg-gradient-to-t from-green-200 to-green-400 %s", cardinalTransition),
	}

	forestGrassCorner := Tile{
		Sockets: [4]socket.Socket{
			socket.ForestGrassCornerN,
			socket.GrassT,
			socket.GrassT,
			socket.ForestGrassCornerW,
		},
		Style: fmt.Sprintf("bg-gradient-to-tl from-green-200 to-green-400 %s", diagonalTransition),
	}

	grassForestCorner := Tile{
		Sockets: [4]socket.Socket{
			socket.GrassForestCornerN,
			socket.ForestT,
			socket.ForestT,
			socket.GrassForestCornerW,
		},
		Style: fmt.Sprintf("bg-gradient-to-tl from-green-400 to-green-200 %s", diagonalTransition),
	}

	tiles := []Tile{grass, forest, sand, water}

	rotatableTiles := []Tile{waterSand, waterSandCorner, sandWaterCorner, sandGrass, sandGrassCorner, grassSandCorner, forestGrass, forestGrassCorner, grassForestCorner}
	for _, tile := range rotatableTiles {
		for i := 0; i < 4; i++ {
			tiles = append(tiles, rotate(tile, i))
		}
	}

	return tiles
}

func rotate(tile Tile, rotations int) Tile {
	totalSockets := len(tile.Sockets)
	newSockets := [4]socket.Socket{}
	for i := 0; i < totalSockets; i++ {
		newSockets[i] = tile.Sockets[(i-rotations+totalSockets)%totalSockets]
	}
	tile.Sockets = newSockets
	tile.Style = getRotateClass(rotations, tile.Style)
	return tile
}

func getRotateClass(rotations int, style string) string {
	if strings.Contains(style, "bg-gradient-to-tl") {
		return rotateDiagonalGradientClass(rotations, style)
	} else if strings.Contains(style, "bg-gradient-to-t") {
		return rotateTopGradientClass(rotations, style)
	}

	return style
}

func rotateDiagonalGradientClass(rotations int, style string) string {
	rotationClasses := []string{"bg-gradient-to-tl", "bg-gradient-to-bl", "bg-gradient-to-br", "bg-gradient-to-tr"}
	count := len(rotationClasses)
	newRotation := rotationClasses[(0-rotations+count)%count]

	return strings.Replace(style, rotationClasses[0], newRotation, 1)
}

func rotateTopGradientClass(rotations int, style string) string {
	rotationClasses := []string{"bg-gradient-to-t", "bg-gradient-to-l", "bg-gradient-to-b", "bg-gradient-to-r"}
	count := len(rotationClasses)
	newRotation := rotationClasses[(0-rotations+count)%count]

	return strings.Replace(style, rotationClasses[0], newRotation, 1)
}
