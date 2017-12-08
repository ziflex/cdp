// Code generated by cdpgen. DO NOT EDIT.

package debugger

import (
	"github.com/mafredri/cdp/protocol/runtime"
)

// BreakpointID Breakpoint identifier.
type BreakpointID string

// CallFrameID Call frame identifier.
type CallFrameID string

// Location Location in the source code.
type Location struct {
	ScriptID     runtime.ScriptID `json:"scriptId"`               // Script identifier as reported in the `Debugger.scriptParsed`.
	LineNumber   int              `json:"lineNumber"`             // Line number in the script (0-based).
	ColumnNumber *int             `json:"columnNumber,omitempty"` // Column number in the script (0-based).
}

// ScriptPosition Location in the source code.
//
// Note: This type is experimental.
type ScriptPosition struct {
	LineNumber   int `json:"lineNumber"`   // No description.
	ColumnNumber int `json:"columnNumber"` // No description.
}

// CallFrame JavaScript call frame. Array of call frames form the call
// stack.
type CallFrame struct {
	CallFrameID      CallFrameID           `json:"callFrameId"`                // Call frame identifier. This identifier is only valid while the virtual machine is paused.
	FunctionName     string                `json:"functionName"`               // Name of the JavaScript function called on this call frame.
	FunctionLocation *Location             `json:"functionLocation,omitempty"` // Location in the source code.
	Location         Location              `json:"location"`                   // Location in the source code.
	URL              string                `json:"url"`                        // JavaScript script name or url.
	ScopeChain       []Scope               `json:"scopeChain"`                 // Scope chain for this call frame.
	This             runtime.RemoteObject  `json:"this"`                       // `this` object for this call frame.
	ReturnValue      *runtime.RemoteObject `json:"returnValue,omitempty"`      // The value being returned, if the function is at return point.
}

// Scope Scope description.
type Scope struct {
	// Type Scope type.
	//
	// Values: "global", "local", "with", "closure", "catch", "block", "script", "eval", "module".
	Type          string               `json:"type"`
	Object        runtime.RemoteObject `json:"object"`                  // Object representing the scope. For `global` and `with` scopes it represents the actual object; for the rest of the scopes, it is artificial transient object enumerating scope variables as its properties.
	Name          *string              `json:"name,omitempty"`          // No description.
	StartLocation *Location            `json:"startLocation,omitempty"` // Location in the source code where scope starts
	EndLocation   *Location            `json:"endLocation,omitempty"`   // Location in the source code where scope ends
}

// SearchMatch Search match for resource.
type SearchMatch struct {
	LineNumber  float64 `json:"lineNumber"`  // Line number in resource content.
	LineContent string  `json:"lineContent"` // Line with match content.
}

// BreakLocation
type BreakLocation struct {
	ScriptID     runtime.ScriptID `json:"scriptId"`               // Script identifier as reported in the `Debugger.scriptParsed`.
	LineNumber   int              `json:"lineNumber"`             // Line number in the script (0-based).
	ColumnNumber *int             `json:"columnNumber,omitempty"` // Column number in the script (0-based).
	// Type
	//
	// Values: "debuggerStatement", "call", "return".
	Type *string `json:"type,omitempty"`
}
