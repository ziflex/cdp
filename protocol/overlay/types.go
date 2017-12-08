// Code generated by cdpgen. DO NOT EDIT.

package overlay

import (
	"github.com/mafredri/cdp/protocol/dom"
)

// HighlightConfig Configuration data for the highlighting of page
// elements.
type HighlightConfig struct {
	ShowInfo           *bool     `json:"showInfo,omitempty"`           // Whether the node info tooltip should be shown (default: false).
	ShowRulers         *bool     `json:"showRulers,omitempty"`         // Whether the rulers should be shown (default: false).
	ShowExtensionLines *bool     `json:"showExtensionLines,omitempty"` // Whether the extension lines from node to the rulers should be shown (default: false).
	DisplayAsMaterial  *bool     `json:"displayAsMaterial,omitempty"`  // No description.
	ContentColor       *dom.RGBA `json:"contentColor,omitempty"`       // The content box highlight fill color (default: transparent).
	PaddingColor       *dom.RGBA `json:"paddingColor,omitempty"`       // The padding highlight fill color (default: transparent).
	BorderColor        *dom.RGBA `json:"borderColor,omitempty"`        // The border highlight fill color (default: transparent).
	MarginColor        *dom.RGBA `json:"marginColor,omitempty"`        // The margin highlight fill color (default: transparent).
	EventTargetColor   *dom.RGBA `json:"eventTargetColor,omitempty"`   // The event target element highlight fill color (default: transparent).
	ShapeColor         *dom.RGBA `json:"shapeColor,omitempty"`         // The shape outside fill color (default: transparent).
	ShapeMarginColor   *dom.RGBA `json:"shapeMarginColor,omitempty"`   // The shape margin fill color (default: transparent).
	SelectorList       *string   `json:"selectorList,omitempty"`       // Selectors to highlight relevant nodes.
	CSSGridColor       *dom.RGBA `json:"cssGridColor,omitempty"`       // The grid layout color (default: transparent).
}

// InspectMode
type InspectMode string

// InspectMode as enums.
const (
	InspectModeNotSet               InspectMode = ""
	InspectModeSearchForNode        InspectMode = "searchForNode"
	InspectModeSearchForUAShadowDOM InspectMode = "searchForUAShadowDOM"
	InspectModeNone                 InspectMode = "none"
)

func (e InspectMode) Valid() bool {
	switch e {
	case "searchForNode", "searchForUAShadowDOM", "none":
		return true
	default:
		return false
	}
}

func (e InspectMode) String() string {
	return string(e)
}
