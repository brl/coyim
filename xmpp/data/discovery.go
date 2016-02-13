package data

import "encoding/xml"

// DiscoveryReply contains the deserialized information about a discovery reply
type DiscoveryReply struct {
	XMLName    xml.Name            `xml:"http://jabber.org/protocol/disco#info query"`
	Node       string              `xml:"node"`
	Identities []DiscoveryIdentity `xml:"identity"`
	Features   []DiscoveryFeature  `xml:"feature"`
	Forms      []Form              `xml:"jabber:x:data x"`
}

// DiscoveryIdentity contains identity information for a specific discovery
type DiscoveryIdentity struct {
	XMLName  xml.Name `xml:"http://jabber.org/protocol/disco#info identity"`
	Lang     string   `xml:"lang,attr,omitempty"`
	Category string   `xml:"category,attr"`
	Type     string   `xml:"type,attr"`
	Name     string   `xml:"name,attr"`
}

// DiscoveryFeature contains information about a specific discovery feature
type DiscoveryFeature struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/disco#info feature"`
	Var     string   `xml:"var,attr"`
}
