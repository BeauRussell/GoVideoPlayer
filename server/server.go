package server

import (
	"sync"

	"github.com/pion/webrtc/v4"
)

var (
	peerConnections []*webrtc.PeerConnection
	mutex           sync.Mutex
)

func AddPeerConnection(pc *webrtc.PeerConnection) {
	mutex.Lock()
	defer mutex.Unlock()
	peerConnections = append(peerConnections, pc)
}
