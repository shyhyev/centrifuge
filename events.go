package centrifuge

import (
	"context"
)

// ConnectEvent contains fields related to connecting event.
type ConnectEvent struct {
	// ClientID that was generated by library for client connection.
	ClientID string
	// Token received from client as part of Connect Command.
	Token string
	// Data received from client as part of Connect Command.
	Data []byte
}

// ConnectReply contains fields determining the reaction on auth event.
type ConnectReply struct {
	// Context allows to return modified context.
	Context context.Context
	// Error for connect command reply.
	Error *Error
	// Disconnect client.
	Disconnect *Disconnect
	// Credentials should be set if app wants to authenticate connection.
	// This field still optional as auth could be provided through HTTP middleware
	// or via JWT token.
	Credentials *Credentials
	// Data allows to set custom data in connect reply.
	Data []byte
	// Channels slice contains channels to subscribe connection to on server-side.
	Channels []string
}

// ConnectingHandler called when new client authenticates on server.
type ConnectingHandler func(context.Context, TransportInfo, ConnectEvent) ConnectReply

// ConnectedHandler called when new client connects to server.
type ConnectedHandler func(context.Context, *Client)

// RefreshEvent contains fields related to refresh event.
type RefreshEvent struct{}

// RefreshReply contains fields determining the reaction on refresh event.
type RefreshReply struct {
	// Expired when set mean that connection must be closed with DisconnectExpired reason.
	Expired bool
	// ExpireAt defines time in future when connection should expire,
	// zero value means no expiration.
	ExpireAt int64
	// Info allows to modify connection information, zero value means no modification.
	Info []byte
}

// RefreshHandler called when it's time to validate client connection and
// update it's expiration time.
type RefreshHandler func(context.Context, *Client, RefreshEvent) RefreshReply

// DisconnectEvent contains fields related to disconnect event.
type DisconnectEvent struct {
	Disconnect *Disconnect
}

// DisconnectReply contains fields determining the reaction on disconnect event.
type DisconnectReply struct{}

// DisconnectHandler called when client disconnects from server.
type DisconnectHandler func(DisconnectEvent) DisconnectReply

// SubscribeEvent contains fields related to subscribe event.
type SubscribeEvent struct {
	Channel string
}

// SubscribeReply contains fields determining the reaction on subscribe event.
type SubscribeReply struct {
	// Error to return, nil value means no error.
	Error *Error
	// Disconnect client, nil value means no disconnect.
	Disconnect *Disconnect
	// ExpireAt defines time in future when subscription should expire,
	// zero value means no expiration.
	ExpireAt int64
	// ChannelInfo defines custom channel information, zero value means no channel information.
	ChannelInfo []byte
}

// SubscribeHandler called when client wants to subscribe on channel.
type SubscribeHandler func(SubscribeEvent) SubscribeReply

// UnsubscribeEvent contains fields related to unsubscribe event.
type UnsubscribeEvent struct {
	Channel string
}

// UnsubscribeReply contains fields determining the reaction on unsubscribe event.
type UnsubscribeReply struct {
}

// UnsubscribeHandler called when client unsubscribed from channel.
type UnsubscribeHandler func(UnsubscribeEvent) UnsubscribeReply

// PublishEvent contains fields related to publish event.
type PublishEvent struct {
	Channel string
	Data    []byte
	Info    *ClientInfo
}

// PublishReply contains fields determining the reaction on publish event.
type PublishReply struct {
	// Error to return, nil value means no error.
	Error *Error
	// Disconnect client, nil value means no disconnect.
	Disconnect *Disconnect
	// Data is modified data to publish, zero value means no modification
	// of original data published by client.
	Data []byte
}

// PublishHandler called when client publishes into channel.
type PublishHandler func(PublishEvent) PublishReply

// SubRefreshEvent contains fields related to subscription refresh event.
type SubRefreshEvent struct {
	Channel string
}

// SubRefreshReply contains fields determining the reaction on
// subscription refresh event.
type SubRefreshReply struct {
	Expired  bool
	ExpireAt int64
	Info     []byte
}

// SubRefreshHandler called when it's time to validate client subscription to channel and
// update it's state if needed.
type SubRefreshHandler func(SubRefreshEvent) SubRefreshReply

// RPCEvent contains fields related to rpc request.
type RPCEvent struct {
	Data []byte
}

// RPCReply contains fields determining the reaction on rpc request.
type RPCReply struct {
	// Error to return, nil value means no error.
	Error *Error
	// Disconnect client, nil value means no disconnect.
	Disconnect *Disconnect
	// Data to return in RPC reply to client.
	Data []byte
}

// RPCHandler must handle incoming command from client.
type RPCHandler func(RPCEvent) RPCReply

// MessageEvent contains fields related to message request.
type MessageEvent struct {
	Data []byte
}

// MessageReply contains fields determining the reaction on message request.
type MessageReply struct {
	Disconnect *Disconnect
}

// MessageHandler must handle incoming async message from client.
type MessageHandler func(MessageEvent) MessageReply
