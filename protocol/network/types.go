// Code generated by cdpgen. DO NOT EDIT.

package network

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/mafredri/cdp/protocol/runtime"
	"github.com/mafredri/cdp/protocol/security"
)

// LoaderID Unique loader identifier.
type LoaderID string

// RequestID Unique request identifier.
type RequestID string

// InterceptionID Unique intercepted request identifier.
type InterceptionID string

// ErrorReason Network level fetch failure reason.
type ErrorReason string

// ErrorReason as enums.
const (
	ErrorReasonNotSet               ErrorReason = ""
	ErrorReasonFailed               ErrorReason = "Failed"
	ErrorReasonAborted              ErrorReason = "Aborted"
	ErrorReasonTimedOut             ErrorReason = "TimedOut"
	ErrorReasonAccessDenied         ErrorReason = "AccessDenied"
	ErrorReasonConnectionClosed     ErrorReason = "ConnectionClosed"
	ErrorReasonConnectionReset      ErrorReason = "ConnectionReset"
	ErrorReasonConnectionRefused    ErrorReason = "ConnectionRefused"
	ErrorReasonConnectionAborted    ErrorReason = "ConnectionAborted"
	ErrorReasonConnectionFailed     ErrorReason = "ConnectionFailed"
	ErrorReasonNameNotResolved      ErrorReason = "NameNotResolved"
	ErrorReasonInternetDisconnected ErrorReason = "InternetDisconnected"
	ErrorReasonAddressUnreachable   ErrorReason = "AddressUnreachable"
)

func (e ErrorReason) Valid() bool {
	switch e {
	case "Failed", "Aborted", "TimedOut", "AccessDenied", "ConnectionClosed", "ConnectionReset", "ConnectionRefused", "ConnectionAborted", "ConnectionFailed", "NameNotResolved", "InternetDisconnected", "AddressUnreachable":
		return true
	default:
		return false
	}
}

func (e ErrorReason) String() string {
	return string(e)
}

// TimeSinceEpoch UTC time in seconds, counted from January 1, 1970.
type TimeSinceEpoch float64

// String calls (time.Time).String().
func (t TimeSinceEpoch) String() string {
	return t.Time().String()
}

// Time parses the Unix time with millisecond accuracy.
func (t TimeSinceEpoch) Time() time.Time {
	secs := int64(t)
	// The Unix time in t only has ms accuracy.
	ms := int64((float64(t) - float64(secs)) * 1000000)
	return time.Unix(secs, ms*1000)
}

// MarshalJSON implements json.Marshaler. Encodes to null if t is zero.
func (t TimeSinceEpoch) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	f := float64(t)
	return json.Marshal(&f)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TimeSinceEpoch) UnmarshalJSON(data []byte) error {
	*t = 0
	if len(data) == 0 {
		return nil
	}
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return errors.New("network.TimeSinceEpoch: " + err.Error())
	}
	*t = TimeSinceEpoch(f)
	return nil
}

var _ json.Marshaler = (*TimeSinceEpoch)(nil)
var _ json.Unmarshaler = (*TimeSinceEpoch)(nil)

// MonotonicTime Monotonically increasing time in seconds since an
// arbitrary point in the past.
type MonotonicTime float64

// String calls (time.Time).String().
func (t MonotonicTime) String() string {
	return t.Time().String()
}

// Time parses the Unix time with millisecond accuracy.
func (t MonotonicTime) Time() time.Time {
	secs := int64(t)
	// The Unix time in t only has ms accuracy.
	ms := int64((float64(t) - float64(secs)) * 1000000)
	return time.Unix(secs, ms*1000)
}

// MarshalJSON implements json.Marshaler. Encodes to null if t is zero.
func (t MonotonicTime) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	f := float64(t)
	return json.Marshal(&f)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *MonotonicTime) UnmarshalJSON(data []byte) error {
	*t = 0
	if len(data) == 0 {
		return nil
	}
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return errors.New("network.MonotonicTime: " + err.Error())
	}
	*t = MonotonicTime(f)
	return nil
}

var _ json.Marshaler = (*MonotonicTime)(nil)
var _ json.Unmarshaler = (*MonotonicTime)(nil)

// Headers Request / response headers as keys / values of JSON object.
type Headers []byte

// MarshalJSON copies behavior of json.RawMessage.
func (h Headers) MarshalJSON() ([]byte, error) {
	if h == nil {
		return []byte("null"), nil
	}
	return h, nil
}

// UnmarshalJSON copies behavior of json.RawMessage.
func (h *Headers) UnmarshalJSON(data []byte) error {
	if h == nil {
		return errors.New("network.Headers: UnmarshalJSON on nil pointer")
	}
	*h = append((*h)[0:0], data...)
	return nil
}

var _ json.Marshaler = (*Headers)(nil)
var _ json.Unmarshaler = (*Headers)(nil)

// ConnectionType The underlying connection technology that the
// browser is supposedly using.
type ConnectionType string

// ConnectionType as enums.
const (
	ConnectionTypeNotSet     ConnectionType = ""
	ConnectionTypeNone       ConnectionType = "none"
	ConnectionTypeCellular2g ConnectionType = "cellular2g"
	ConnectionTypeCellular3g ConnectionType = "cellular3g"
	ConnectionTypeCellular4g ConnectionType = "cellular4g"
	ConnectionTypeBluetooth  ConnectionType = "bluetooth"
	ConnectionTypeEthernet   ConnectionType = "ethernet"
	ConnectionTypeWifi       ConnectionType = "wifi"
	ConnectionTypeWimax      ConnectionType = "wimax"
	ConnectionTypeOther      ConnectionType = "other"
)

func (e ConnectionType) Valid() bool {
	switch e {
	case "none", "cellular2g", "cellular3g", "cellular4g", "bluetooth", "ethernet", "wifi", "wimax", "other":
		return true
	default:
		return false
	}
}

func (e ConnectionType) String() string {
	return string(e)
}

// CookieSameSite Represents the cookie's 'SameSite' status:
// https://tools.ietf.org/html/draft-west-first-party-cookies
type CookieSameSite string

// CookieSameSite as enums.
const (
	CookieSameSiteNotSet CookieSameSite = ""
	CookieSameSiteStrict CookieSameSite = "Strict"
	CookieSameSiteLax    CookieSameSite = "Lax"
)

func (e CookieSameSite) Valid() bool {
	switch e {
	case "Strict", "Lax":
		return true
	default:
		return false
	}
}

func (e CookieSameSite) String() string {
	return string(e)
}

// ResourceTiming Timing information for the request.
type ResourceTiming struct {
	RequestTime  float64 `json:"requestTime"`  // Timing's requestTime is a baseline in seconds, while the other numbers are ticks in milliseconds relatively to this requestTime.
	ProxyStart   float64 `json:"proxyStart"`   // Started resolving proxy.
	ProxyEnd     float64 `json:"proxyEnd"`     // Finished resolving proxy.
	DNSStart     float64 `json:"dnsStart"`     // Started DNS address resolve.
	DNSEnd       float64 `json:"dnsEnd"`       // Finished DNS address resolve.
	ConnectStart float64 `json:"connectStart"` // Started connecting to the remote host.
	ConnectEnd   float64 `json:"connectEnd"`   // Connected to the remote host.
	SslStart     float64 `json:"sslStart"`     // Started SSL handshake.
	SslEnd       float64 `json:"sslEnd"`       // Finished SSL handshake.
	// WorkerStart Started running ServiceWorker.
	//
	// Note: This property is experimental.
	WorkerStart float64 `json:"workerStart"`
	// WorkerReady Finished Starting ServiceWorker.
	//
	// Note: This property is experimental.
	WorkerReady float64 `json:"workerReady"`
	SendStart   float64 `json:"sendStart"` // Started sending request.
	SendEnd     float64 `json:"sendEnd"`   // Finished sending request.
	// PushStart Time the server started pushing request.
	//
	// Note: This property is experimental.
	PushStart float64 `json:"pushStart"`
	// PushEnd Time the server finished pushing request.
	//
	// Note: This property is experimental.
	PushEnd           float64 `json:"pushEnd"`
	ReceiveHeadersEnd float64 `json:"receiveHeadersEnd"` // Finished receiving response headers.
}

// ResourcePriority Loading priority of a resource request.
type ResourcePriority string

// ResourcePriority as enums.
const (
	ResourcePriorityNotSet   ResourcePriority = ""
	ResourcePriorityVeryLow  ResourcePriority = "VeryLow"
	ResourcePriorityLow      ResourcePriority = "Low"
	ResourcePriorityMedium   ResourcePriority = "Medium"
	ResourcePriorityHigh     ResourcePriority = "High"
	ResourcePriorityVeryHigh ResourcePriority = "VeryHigh"
)

func (e ResourcePriority) Valid() bool {
	switch e {
	case "VeryLow", "Low", "Medium", "High", "VeryHigh":
		return true
	default:
		return false
	}
}

func (e ResourcePriority) String() string {
	return string(e)
}

// Request HTTP request data.
type Request struct {
	URL              string                     `json:"url"`                        // Request URL.
	Method           string                     `json:"method"`                     // HTTP request method.
	Headers          Headers                    `json:"headers"`                    // HTTP request headers.
	PostData         *string                    `json:"postData,omitempty"`         // HTTP POST request data.
	MixedContentType *security.MixedContentType `json:"mixedContentType,omitempty"` // The mixed content type of the request.
	InitialPriority  ResourcePriority           `json:"initialPriority"`            // Priority of the resource request at the time request is sent.
	// ReferrerPolicy The referrer policy of the request, as defined in
	// https://www.w3.org/TR/referrer-policy/
	//
	// Values: "unsafe-url", "no-referrer-when-downgrade", "no-referrer", "origin", "origin-when-cross-origin", "same-origin", "strict-origin", "strict-origin-when-cross-origin".
	ReferrerPolicy string `json:"referrerPolicy"`
	IsLinkPreload  *bool  `json:"isLinkPreload,omitempty"` // Whether is loaded via link preload.
}

// SignedCertificateTimestamp Details of a signed certificate
// timestamp (SCT).
type SignedCertificateTimestamp struct {
	Status             string         `json:"status"`             // Validation status.
	Origin             string         `json:"origin"`             // Origin.
	LogDescription     string         `json:"logDescription"`     // Log name / description.
	LogID              string         `json:"logId"`              // Log ID.
	Timestamp          TimeSinceEpoch `json:"timestamp"`          // Issuance date.
	HashAlgorithm      string         `json:"hashAlgorithm"`      // Hash algorithm.
	SignatureAlgorithm string         `json:"signatureAlgorithm"` // Signature algorithm.
	SignatureData      string         `json:"signatureData"`      // Signature data.
}

// SecurityDetails Security details about a request.
type SecurityDetails struct {
	Protocol                       string                       `json:"protocol"`                       // Protocol name (e.g. "TLS 1.2" or "QUIC").
	KeyExchange                    string                       `json:"keyExchange"`                    // Key Exchange used by the connection, or the empty string if not applicable.
	KeyExchangeGroup               *string                      `json:"keyExchangeGroup,omitempty"`     // (EC)DH group used by the connection, if applicable.
	Cipher                         string                       `json:"cipher"`                         // Cipher name.
	Mac                            *string                      `json:"mac,omitempty"`                  // TLS MAC. Note that AEAD ciphers do not have separate MACs.
	CertificateID                  security.CertificateID       `json:"certificateId"`                  // Certificate ID value.
	SubjectName                    string                       `json:"subjectName"`                    // Certificate subject name.
	SanList                        []string                     `json:"sanList"`                        // Subject Alternative Name (SAN) DNS names and IP addresses.
	Issuer                         string                       `json:"issuer"`                         // Name of the issuing CA.
	ValidFrom                      TimeSinceEpoch               `json:"validFrom"`                      // Certificate valid from date.
	ValidTo                        TimeSinceEpoch               `json:"validTo"`                        // Certificate valid to (expiration) date
	SignedCertificateTimestampList []SignedCertificateTimestamp `json:"signedCertificateTimestampList"` // List of signed certificate timestamps (SCTs).
}

// BlockedReason The reason why request was blocked.
type BlockedReason string

// BlockedReason as enums.
const (
	BlockedReasonNotSet            BlockedReason = ""
	BlockedReasonCsp               BlockedReason = "csp"
	BlockedReasonMixedContent      BlockedReason = "mixed-content"
	BlockedReasonOrigin            BlockedReason = "origin"
	BlockedReasonInspector         BlockedReason = "inspector"
	BlockedReasonSubresourceFilter BlockedReason = "subresource-filter"
	BlockedReasonOther             BlockedReason = "other"
)

func (e BlockedReason) Valid() bool {
	switch e {
	case "csp", "mixed-content", "origin", "inspector", "subresource-filter", "other":
		return true
	default:
		return false
	}
}

func (e BlockedReason) String() string {
	return string(e)
}

// Response HTTP response data.
type Response struct {
	URL                string           `json:"url"`                          // Response URL. This URL can be different from CachedResource.url in case of redirect.
	Status             int              `json:"status"`                       // HTTP response status code.
	StatusText         string           `json:"statusText"`                   // HTTP response status text.
	Headers            Headers          `json:"headers"`                      // HTTP response headers.
	HeadersText        *string          `json:"headersText,omitempty"`        // HTTP response headers text.
	MimeType           string           `json:"mimeType"`                     // Resource mimeType as determined by the browser.
	RequestHeaders     Headers          `json:"requestHeaders,omitempty"`     // Refined HTTP request headers that were actually transmitted over the network.
	RequestHeadersText *string          `json:"requestHeadersText,omitempty"` // HTTP request headers text.
	ConnectionReused   bool             `json:"connectionReused"`             // Specifies whether physical connection was actually reused for this request.
	ConnectionID       float64          `json:"connectionId"`                 // Physical connection id that was actually used for this request.
	RemoteIPAddress    *string          `json:"remoteIPAddress,omitempty"`    // Remote IP address.
	RemotePort         *int             `json:"remotePort,omitempty"`         // Remote port.
	FromDiskCache      *bool            `json:"fromDiskCache,omitempty"`      // Specifies that the request was served from the disk cache.
	FromServiceWorker  *bool            `json:"fromServiceWorker,omitempty"`  // Specifies that the request was served from the ServiceWorker.
	EncodedDataLength  float64          `json:"encodedDataLength"`            // Total number of bytes received for this request so far.
	Timing             *ResourceTiming  `json:"timing,omitempty"`             // Timing information for the given request.
	Protocol           *string          `json:"protocol,omitempty"`           // Protocol used to fetch this request.
	SecurityState      security.State   `json:"securityState"`                // Security state of the request resource.
	SecurityDetails    *SecurityDetails `json:"securityDetails,omitempty"`    // Security details for the request.
}

// WebSocketRequest WebSocket request data.
type WebSocketRequest struct {
	Headers Headers `json:"headers"` // HTTP request headers.
}

// WebSocketResponse WebSocket response data.
type WebSocketResponse struct {
	Status             int     `json:"status"`                       // HTTP response status code.
	StatusText         string  `json:"statusText"`                   // HTTP response status text.
	Headers            Headers `json:"headers"`                      // HTTP response headers.
	HeadersText        *string `json:"headersText,omitempty"`        // HTTP response headers text.
	RequestHeaders     Headers `json:"requestHeaders,omitempty"`     // HTTP request headers.
	RequestHeadersText *string `json:"requestHeadersText,omitempty"` // HTTP request headers text.
}

// WebSocketFrame WebSocket frame data.
type WebSocketFrame struct {
	Opcode      float64 `json:"opcode"`      // WebSocket frame opcode.
	Mask        bool    `json:"mask"`        // WebSocke frame mask.
	PayloadData string  `json:"payloadData"` // WebSocke frame payload data.
}

// Initiator Information about the request initiator.
type Initiator struct {
	// Type Type of this initiator.
	//
	// Values: "parser", "script", "preload", "other".
	Type       string              `json:"type"`
	Stack      *runtime.StackTrace `json:"stack,omitempty"`      // Initiator JavaScript stack trace, set for Script only.
	URL        *string             `json:"url,omitempty"`        // Initiator URL, set for Parser type or for Script type (when script is importing module).
	LineNumber *float64            `json:"lineNumber,omitempty"` // Initiator line number, set for Parser type or for Script type (when script is importing module) (0-based).
}

// Cookie Cookie object
type Cookie struct {
	Name     string         `json:"name"`               // Cookie name.
	Value    string         `json:"value"`              // Cookie value.
	Domain   string         `json:"domain"`             // Cookie domain.
	Path     string         `json:"path"`               // Cookie path.
	Expires  float64        `json:"expires"`            // Cookie expiration date as the number of seconds since the UNIX epoch.
	Size     int            `json:"size"`               // Cookie size.
	HTTPOnly bool           `json:"httpOnly"`           // True if cookie is http-only.
	Secure   bool           `json:"secure"`             // True if cookie is secure.
	Session  bool           `json:"session"`            // True in case of session cookie.
	SameSite CookieSameSite `json:"sameSite,omitempty"` // Cookie SameSite type.
}

// CookieParam Cookie parameter object
type CookieParam struct {
	Name     string         `json:"name"`               // Cookie name.
	Value    string         `json:"value"`              // Cookie value.
	URL      *string        `json:"url,omitempty"`      // The request-URI to associate with the setting of the cookie. This value can affect the default domain and path values of the created cookie.
	Domain   *string        `json:"domain,omitempty"`   // Cookie domain.
	Path     *string        `json:"path,omitempty"`     // Cookie path.
	Secure   *bool          `json:"secure,omitempty"`   // True if cookie is secure.
	HTTPOnly *bool          `json:"httpOnly,omitempty"` // True if cookie is http-only.
	SameSite CookieSameSite `json:"sameSite,omitempty"` // Cookie SameSite type.
	Expires  TimeSinceEpoch `json:"expires,omitempty"`  // Cookie expiration date, session cookie if not set
}

// AuthChallenge Authorization challenge for HTTP status code 401 or
// 407.
//
// Note: This type is experimental.
type AuthChallenge struct {
	// Source Source of the authentication challenge.
	//
	// Values: "Server", "Proxy".
	Source *string `json:"source,omitempty"`
	Origin string  `json:"origin"` // Origin of the challenger.
	Scheme string  `json:"scheme"` // The authentication scheme used, such as basic or digest
	Realm  string  `json:"realm"`  // The realm of the challenge. May be empty.
}

// AuthChallengeResponse Response to an AuthChallenge.
//
// Note: This type is experimental.
type AuthChallengeResponse struct {
	// Response The decision on what to do in response to the authorization
	// challenge. Default means deferring to the default behavior of the
	// net stack, which will likely either the Cancel authentication or
	// display a popup dialog box.
	//
	// Values: "Default", "CancelAuth", "ProvideCredentials".
	Response string  `json:"response"`
	Username *string `json:"username,omitempty"` // The username to provide, possibly empty. Should only be set if response is ProvideCredentials.
	Password *string `json:"password,omitempty"` // The password to provide, possibly empty. Should only be set if response is ProvideCredentials.
}

// InterceptionStage Stages of the interception to begin intercepting.
// Request will intercept before the request is sent. Response will
// intercept after the response is received.
//
// Note: This type is experimental.
type InterceptionStage string

// InterceptionStage as enums.
const (
	InterceptionStageNotSet          InterceptionStage = ""
	InterceptionStageRequest         InterceptionStage = "Request"
	InterceptionStageHeadersReceived InterceptionStage = "HeadersReceived"
)

func (e InterceptionStage) Valid() bool {
	switch e {
	case "Request", "HeadersReceived":
		return true
	default:
		return false
	}
}

func (e InterceptionStage) String() string {
	return string(e)
}
