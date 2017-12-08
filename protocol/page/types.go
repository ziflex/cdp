// Code generated by cdpgen. DO NOT EDIT.

package page

import (
	"github.com/mafredri/cdp/protocol/network"
)

// FrameResourceTree Information about the Frame hierarchy along with
// their cached resources.
//
// Note: This type is experimental.
type FrameResourceTree struct {
	Frame       Frame               `json:"frame"`                 // Frame information for this tree item.
	ChildFrames []FrameResourceTree `json:"childFrames,omitempty"` // Child frames.
	Resources   []FrameResource     `json:"resources"`             // Information about frame resources.
}

// FrameTree Information about the Frame hierarchy.
type FrameTree struct {
	Frame       Frame       `json:"frame"`                 // Frame information for this tree item.
	ChildFrames []FrameTree `json:"childFrames,omitempty"` // Child frames.
}

// ScriptIdentifier Unique script identifier.
type ScriptIdentifier string

// TransitionType Transition type.
type TransitionType string

// TransitionType as enums.
const (
	TransitionTypeNotSet           TransitionType = ""
	TransitionTypeLink             TransitionType = "link"
	TransitionTypeTyped            TransitionType = "typed"
	TransitionTypeAutoBookmark     TransitionType = "auto_bookmark"
	TransitionTypeAutoSubframe     TransitionType = "auto_subframe"
	TransitionTypeManualSubframe   TransitionType = "manual_subframe"
	TransitionTypeGenerated        TransitionType = "generated"
	TransitionTypeAutoToplevel     TransitionType = "auto_toplevel"
	TransitionTypeFormSubmit       TransitionType = "form_submit"
	TransitionTypeReload           TransitionType = "reload"
	TransitionTypeKeyword          TransitionType = "keyword"
	TransitionTypeKeywordGenerated TransitionType = "keyword_generated"
	TransitionTypeOther            TransitionType = "other"
)

func (e TransitionType) Valid() bool {
	switch e {
	case "link", "typed", "auto_bookmark", "auto_subframe", "manual_subframe", "generated", "auto_toplevel", "form_submit", "reload", "keyword", "keyword_generated", "other":
		return true
	default:
		return false
	}
}

func (e TransitionType) String() string {
	return string(e)
}

// NavigationEntry Navigation history entry.
type NavigationEntry struct {
	ID             int            `json:"id"`             // Unique id of the navigation history entry.
	URL            string         `json:"url"`            // URL of the navigation history entry.
	UserTypedURL   string         `json:"userTypedURL"`   // URL that the user typed in the url bar.
	Title          string         `json:"title"`          // Title of the navigation history entry.
	TransitionType TransitionType `json:"transitionType"` // Transition type.
}

// ScreencastFrameMetadata Screencast frame metadata.
//
// Note: This type is experimental.
type ScreencastFrameMetadata struct {
	OffsetTop       float64                `json:"offsetTop"`           // Top offset in DIP.
	PageScaleFactor float64                `json:"pageScaleFactor"`     // Page scale factor.
	DeviceWidth     float64                `json:"deviceWidth"`         // Device screen width in DIP.
	DeviceHeight    float64                `json:"deviceHeight"`        // Device screen height in DIP.
	ScrollOffsetX   float64                `json:"scrollOffsetX"`       // Position of horizontal scroll in CSS pixels.
	ScrollOffsetY   float64                `json:"scrollOffsetY"`       // Position of vertical scroll in CSS pixels.
	Timestamp       network.TimeSinceEpoch `json:"timestamp,omitempty"` // Frame swap timestamp.
}

// DialogType Javascript dialog type.
type DialogType string

// DialogType as enums.
const (
	DialogTypeNotSet       DialogType = ""
	DialogTypeAlert        DialogType = "alert"
	DialogTypeConfirm      DialogType = "confirm"
	DialogTypePrompt       DialogType = "prompt"
	DialogTypeBeforeunload DialogType = "beforeunload"
)

func (e DialogType) Valid() bool {
	switch e {
	case "alert", "confirm", "prompt", "beforeunload":
		return true
	default:
		return false
	}
}

func (e DialogType) String() string {
	return string(e)
}

// AppManifestError Error while paring app manifest.
type AppManifestError struct {
	Message  string `json:"message"`  // Error message.
	Critical int    `json:"critical"` // If criticial, this is a non-recoverable parse error.
	Line     int    `json:"line"`     // Error line.
	Column   int    `json:"column"`   // Error column.
}

// LayoutViewport Layout viewport position and dimensions.
type LayoutViewport struct {
	PageX        int `json:"pageX"`        // Horizontal offset relative to the document (CSS pixels).
	PageY        int `json:"pageY"`        // Vertical offset relative to the document (CSS pixels).
	ClientWidth  int `json:"clientWidth"`  // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight int `json:"clientHeight"` // Height (CSS pixels), excludes scrollbar if present.
}

// VisualViewport Visual viewport position, dimensions, and scale.
type VisualViewport struct {
	OffsetX      float64 `json:"offsetX"`      // Horizontal offset relative to the layout viewport (CSS pixels).
	OffsetY      float64 `json:"offsetY"`      // Vertical offset relative to the layout viewport (CSS pixels).
	PageX        float64 `json:"pageX"`        // Horizontal offset relative to the document (CSS pixels).
	PageY        float64 `json:"pageY"`        // Vertical offset relative to the document (CSS pixels).
	ClientWidth  float64 `json:"clientWidth"`  // Width (CSS pixels), excludes scrollbar if present.
	ClientHeight float64 `json:"clientHeight"` // Height (CSS pixels), excludes scrollbar if present.
	Scale        float64 `json:"scale"`        // Scale relative to the ideal viewport (size at width=device-width).
}

// Viewport Viewport for capturing screenshot.
type Viewport struct {
	X      float64 `json:"x"`      // X offset in CSS pixels.
	Y      float64 `json:"y"`      // Y offset in CSS pixels
	Width  float64 `json:"width"`  // Rectangle width in CSS pixels
	Height float64 `json:"height"` // Rectangle height in CSS pixels
	Scale  float64 `json:"scale"`  // Page scale factor.
}
