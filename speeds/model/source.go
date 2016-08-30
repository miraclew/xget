package model

const (
    PROTOCOL_HTTP = 1
    PROTOCOL_E2DK = 2
    PROTOCOL_BT = 3
)

// Source represents an file source on internet.
type Source struct {
    Protocol int
    URI string
    SupportRandAccess bool
}
