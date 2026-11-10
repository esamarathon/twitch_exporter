package collector

import "testing"

func TestChannelNamesSet(t *testing.T) {
	tests := []struct {
		name     string
		values   []string
		expected ChannelNames
	}{
		{
			name:     "single value",
			values:   []string{"channel"},
			expected: ChannelNames{"channel"},
		},
		{
			name:     "repeated flags",
			values:   []string{"one", "two"},
			expected: ChannelNames{"one", "two"},
		},
		{
			name:     "comma-separated values",
			values:   []string{"one,two,three"},
			expected: ChannelNames{"one", "two", "three"},
		},
		{
			name:     "mixed single and comma-separated values",
			values:   []string{"one,two", "three,four"},
			expected: ChannelNames{"one", "two", "three", "four"},
		},
		{
			name:     "spaces around names are trimmed",
			values:   []string{" one , two "},
			expected: ChannelNames{"one", "two"},
		},
		{
			name:     "empty names are ignored",
			values:   []string{"one,,two, "},
			expected: ChannelNames{"one", "two"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var channels ChannelNames
			for _, v := range tt.values {
				if err := channels.Set(v); err != nil {
					t.Fatalf("Set(%q) returned error: %v", v, err)
				}
			}
			if len(channels) != len(tt.expected) {
				t.Fatalf("got %v, want %v", channels, tt.expected)
			}
			for i, name := range tt.expected {
				if channels[i] != name {
					t.Fatalf("got %v, want %v", channels, tt.expected)
				}
			}
		})
	}
}
