package gmqtt

type ClientOptions struct {
	// ClientID is the client id for the client.
	ClientID string
	// Username is the username for the client.
	Username string
	// KeepAlive is the keep alive time in seconds for the client.
	// The server will close the client if no there is no packet has been received for 1.5 times the KeepAlive time.
	KeepAlive uint16
	// SessionExpiry is the session expiry interval in seconds.
	// If the client version is v5, this value will be set into CONNACK Session Expiry Interval property.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901082
	SessionExpiry uint32
	// MaxInflight limits the number of QoS 1 and QoS 2 publications that the client is willing to process concurrently.
	// For v3 client, it is default to config.MQTT.MaxInflight.
	// For v5 client, it is the minimum of config.MQTT.MaxInflight and Receive Maximum property in CONNECT packet.
	MaxInflight uint16
	// ReceiveMax limits the number of QoS 1 and QoS 2 publications that the server is willing to process concurrently for the Client.
	// If the client version is v5, this value will be set into Receive Maximum property in CONNACK packet.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901083
	ReceiveMax uint16
	// ClientMaxPacketSize is the maximum packet size that the client is willing to accept.
	// The server will drop the packet if it exceeds ClientMaxPacketSize.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901050
	ClientMaxPacketSize uint32
	// ServerMaxPacketSize is the maximum packet size that the server is willing to accept from the client.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901086
	ServerMaxPacketSize uint32
	// ClientTopicAliasMax is highest value that the client will accept as a Topic Alias sent by the server.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901051
	ClientTopicAliasMax uint16
	// ServerTopicAliasMax is highest value that the server will accept as a Topic Alias sent by the client.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901088
	ServerTopicAliasMax uint16
	// RequestProblemInfo is the value to indicate whether the Reason String or User Properties should be sent in the case of failures.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901053
	RequestProblemInfo bool
	// UserProperties is the user properties provided by the client.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901090
	//UserProperties []*packets.UserProperty
	// WildcardSubAvailable indicates whether the client is permitted to send retained messages.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901091
	RetainAvailable bool
	// WildcardSubAvailable indicates whether the client is permitted to subscribe Wildcard Subscriptions.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901091
	WildcardSubAvailable bool
	// SubIDAvailable indicates whether the client is permitted to set Subscription Identifiers.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901092
	SubIDAvailable bool
	// SharedSubAvailable indicates whether the client is permitted to subscribe Shared Subscriptions.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901093
	SharedSubAvailable bool
	// AuthMethod is the auth method send by the client.
	// Only MQTT v5 client can set this value.
	// See: https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901055
	AuthMethod []byte
}
