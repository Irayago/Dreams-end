package world

type World struct {
	WorldId        int // change to uuid
	PlayerCount    int
	PlayerCapacity int
	clientMap      map[Connection]*Player
	playerMap      map[string]*Player
	tickRate       int
	outboundBuffer chan MessageInterface
	inboundBuffer  chan MessageInterface
	workerPool     chan *Worker // concurrent worker pool with max of 100 Workers
	worldStore     WorldStore
	//test messaging with JSON first before protobuff
}

// workerPool and messageBuffer workflow::
/*




 */

// need to figure out message struct and protobuf serialization

type WorldStore interface {
	SavePlayerState(player *Player) error
	LoadPlayerState(playerId string) (*Player, error)
}

type MessageInterface interface {
	OpCode() int
	Payload() any
}

type Connection interface {
	SendData(data []byte)
	GetPlayerName() string
	SavePlayerState() error
}
