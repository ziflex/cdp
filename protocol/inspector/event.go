// Code generated by cdpgen. DO NOT EDIT.

package inspector

import (
	"github.com/mafredri/cdp/rpcc"
)

// DetachedClient is a client for Detached events. Fired when remote
// debugging connection is about to be terminated. Contains detach
// reason.
type DetachedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*DetachedReply, error)
	rpcc.Stream
}

// DetachedReply is the reply for Detached events.
type DetachedReply struct {
	Reason string `json:"reason"` // The reason why connection has been terminated.
}

// TargetCrashedClient is a client for TargetCrashed events. Fired
// when debugging target has crashed
type TargetCrashedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*TargetCrashedReply, error)
	rpcc.Stream
}

// TargetCrashedReply is the reply for TargetCrashed events.
type TargetCrashedReply struct {
}
