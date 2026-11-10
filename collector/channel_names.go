package collector

import (
	"fmt"
	"strings"
)

// ChannelNames represents a list of twitch channels.
type ChannelNames []string

// IsCumulative is required for kingpin interfaces to allow multiple values
func (c ChannelNames) IsCumulative() bool {
	return true
}

// Set sets the value of a ChannelNames. The value may be a single channel
// name, or a comma-separated list of channel names.
func (c *ChannelNames) Set(v string) error {
	for _, name := range strings.Split(v, ",") {
		if name = strings.TrimSpace(name); name != "" {
			*c = append(*c, name)
		}
	}
	return nil
}

// String returns a string representation of the Channels type.
func (c ChannelNames) String() string {
	return fmt.Sprintf("%v", []string(c))
}
