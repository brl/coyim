package events

import (
	"time"

	"github.com/twstrike/coyim/session/access"
	"github.com/twstrike/coyim/xmpp/data"
)

// Event represents a Session event
type Event struct {
	Type    EventType
	Session access.Session
}

// EventType represents the type of Session event
type EventType int

// Session event types
const (
	Disconnected EventType = iota
	Connecting
	Connected
	ConnectionLost

	RosterReceived
	Ping
	PongReceived
)

// Peer represents an event associated to a peer
type Peer struct {
	Session access.Session
	Type    PeerType
	From    string
}

// Notification represents a notification event
type Notification struct {
	Session      access.Session
	Peer         string
	Notification string
}

// PeerType represents the type of Peer event
type PeerType int

// Peer types
const (
	IQReceived PeerType = iota

	OTREnded
	OTRNewKeys
	OTRRenewedKeys

	SubscriptionRequest
	Subscribed
	Unsubscribe
)

// Presence represents a presence event
type Presence struct {
	Session access.Session
	*data.ClientPresence
	Gone bool
}

// Message represents a message event
type Message struct {
	Session   access.Session
	From      string
	Resource  string
	When      time.Time
	Body      []byte
	Encrypted bool
}

// LogLevel is the current log level
type LogLevel int

// The different available log levels
const (
	Info LogLevel = iota
	Warn
	Alert
)

// Log contains information one specific log event
type Log struct {
	Level   LogLevel
	Message string
}
