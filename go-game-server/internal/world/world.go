package world

type World struct {
	WorldId        int // change to uuid
	PlayerCount    int
	PlayerCapacity int
	clientMap      map[Connection]*Player
	playerMap      map[string]*Player
	tickRate       int
	outboundBuffer chan *Message
	inboundBuffer  chan *Message
	workerPool     chan *Worker // concurrent worker pool with max of 100 Workers
	//test messaging with JSON first before protobuff
}

// workerPool and messageBuffer workflow::
/*




 */

// need to figure out message struct and protobuf serialization

type Connection interface {
	sendData(data []byte)
	getplayerId() string
	savePlayerState() error
}
