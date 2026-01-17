package wol

import (
	"bytes"
	"testing"
)

func TestNewMagicPacket(t *testing.T) {
	tests := []struct {
		name    string
		mac     string
		wantErr bool
	}{
		{
			name:    "valid mac colon",
			mac:     "AA:BB:CC:DD:EE:FF",
			wantErr: false,
		},
		{
			name:    "valid mac hyphen",
			mac:     "aa-bb-cc-dd-ee-ff",
			wantErr: false,
		},
		{
			name:    "invalid mac length",
			mac:     "00:11:22",
			wantErr: true,
		},
		{
			name:    "invalid mac chars",
			mac:     "ZZ:ZZ:ZZ:ZZ:ZZ:ZZ",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packet, err := NewMagicPacket(tt.mac)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMagicPacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify packet structure: 6 bytes of 0xFF followed by 16 repetitions of MAC
				if len(packet) != 6+16*6 {
					t.Errorf("NewMagicPacket() length = %d, want %d", len(packet), 102)
				}

				// Check standard header
				header := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
				if !bytes.Equal(packet[:6], header) {
					t.Errorf("NewMagicPacket() header is invalid")
				}
			}
		})
	}
}
