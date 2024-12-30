package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pion/webrtc/v4"
)

func SignalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var offer webrtc.SessionDescription
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		http.Error(w, "Failed to parse offer", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received SDP offer: %s\n", offer.SDP)

	peerConnection, err := webrtc.NewPeerConnection(webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs:       []string{"turn:global.relay.metered.ca:80"},
				Username:   "",
				Credential: "",
			},
			{
				URLs:       []string{"turn:global.relay.metered.ca:80?transport=tcp"},
				Username:   "",
				Credential: "",
			},
			{
				URLs:       []string{"turn:global.relay.metered.ca:443"},
				Username:   "",
				Credential: "",
			},
			{
				URLs:       []string{"turns:global.relay.metered.ca:443?transport=tcp"},
				Username:   "",
				Credential: "",
			},
		},
	})
	if err != nil {
		http.Error(w, "Failed to create peer connection", http.StatusInternalServerError)
		return
	}

	videoTrack, err := webrtc.NewTrackLocalStaticSample(
		webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeH264},
		"video", "pion",
	)
	if err != nil {
		http.Error(w, "Failed to create video track", http.StatusInternalServerError)
		return
	}

	if _, err := peerConnection.AddTrack(videoTrack); err != nil {
		http.Error(w, "Failed to add video track to peer connection", http.StatusInternalServerError)
		return
	}

	if err := peerConnection.SetRemoteDescription(offer); err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to set remote description", http.StatusInternalServerError)
		return
	}

	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		http.Error(w, "Failed to create answer", http.StatusInternalServerError)
		return
	}
	if err := peerConnection.SetLocalDescription(answer); err != nil {
		http.Error(w, "Failed to set local description", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answer)

	go SendVideoFrames(videoTrack)

	AddPeerConnection(peerConnection)
}
