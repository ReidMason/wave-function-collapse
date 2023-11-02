package socket

type Socket int

const (
	Grass = iota
	Forest
	Water
	Sand

	WaterT
	SandT
	GrassT
	ForestT

	WaterSandW
	WaterSandE
	WaterSandCornerN
	WaterSandCornerW
	SandWaterCornerN
	SandWaterCornerW

	SandGrassW
	SandGrassE
	SandGrassCornerN
	SandGrassCornerW
	GrassSandCornerN
	GrassSandCornerW

	ForestGrassW
	ForestGrassE
	ForestGrassCornerN
	ForestGrassCornerW
	GrassForestCornerN
	GrassForestCornerW
)

var SocketConstraints = map[Socket]map[Socket]bool{
	Grass:  {Grass: true, GrassT: true},
	Forest: {Forest: true, ForestT: true},
	Water:  {Water: true, WaterT: true},
	Sand:   {Sand: true, SandT: true},

	WaterSandW:       {WaterSandE: true},
	WaterSandCornerN: {WaterSandW: true},
	WaterSandCornerW: {WaterSandE: true},
	SandWaterCornerN: {WaterSandE: true},
	SandWaterCornerW: {WaterSandW: true},

	SandGrassW:       {SandGrassE: true},
	SandGrassCornerN: {SandGrassW: true},
	SandGrassCornerW: {SandGrassE: true},
	GrassSandCornerN: {SandGrassE: true},
	GrassSandCornerW: {SandGrassW: true},

	ForestGrassW:       {ForestGrassE: true},
	ForestGrassCornerN: {ForestGrassW: true},
	ForestGrassCornerW: {ForestGrassE: true},
	GrassForestCornerN: {ForestGrassE: true},
	GrassForestCornerW: {ForestGrassW: true},
}

var SocketConstraintsArr [][]bool

func ConvertSocketConstraints() [][]bool {
	total := 26
	newConstraints := make([][]bool, total)

	for socket, hashmap := range SocketConstraints {
		constraints := make([]bool, total)
		for x, y := range hashmap {
			constraints[x] = y
		}
		newConstraints[socket] = constraints
	}

	SocketConstraintsArr = newConstraints
	return newConstraints
}

func CanConnect(socket1, socket2 Socket) bool {
	compatibleSockets := SocketConstraintsArr[socket1]
	if len(compatibleSockets) > 0 {
		foundSocket := compatibleSockets[socket2]
		if foundSocket {
			return true
		}
	}

	compatibleSockets = SocketConstraintsArr[socket2]
	if len(compatibleSockets) > 0 {
		foundSocket := compatibleSockets[socket1]
		if foundSocket {
			return true
		}
	}

	return false
}

func CanConnectOld(socket1, socket2 Socket) bool {
	compatibleSockets, found := SocketConstraints[socket1]
	if found {
		_, found = compatibleSockets[socket2]
		if found {
			return true
		}
	}

	compatibleSockets, found = SocketConstraints[socket2]
	if found {
		_, found = compatibleSockets[socket1]
		if found {
			return true
		}
	}

	return false
}
