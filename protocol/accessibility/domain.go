// Code generated by cdpgen. DO NOT EDIT.

// Package accessibility implements the Accessibility domain.
package accessibility

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Accessibility domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Accessibility domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// GetPartialAXTree invokes the Accessibility method. Fetches the
// accessibility node and partial accessibility tree for this DOM node,
// if it exists.
func (d *domainClient) GetPartialAXTree(ctx context.Context, args *GetPartialAXTreeArgs) (reply *GetPartialAXTreeReply, err error) {
	reply = new(GetPartialAXTreeReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Accessibility.getPartialAXTree", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Accessibility.getPartialAXTree", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Accessibility", Op: "GetPartialAXTree", Err: err}
	}
	return
}
