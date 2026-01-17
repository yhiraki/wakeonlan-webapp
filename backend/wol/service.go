package wol

import (
	"fmt"
	"net"
)

// Service defines the interface for Wake-on-LAN operations
type Service interface {
	Wake(macAddr string) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Wake(macAddr string) error {
	packet, err := NewMagicPacket(macAddr)
	if err != nil {
		return err
	}

	// Broadcast address
	addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:9")
	if err != nil {
		return fmt.Errorf("failed to resolve broadcast address: %w", err)
	}

	// Connect to UDP
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return fmt.Errorf("failed to dial UDP: %w", err)
	}
	defer conn.Close()

	// Send packet
	_, err = conn.Write(packet)
	if err != nil {
		return fmt.Errorf("failed to write packet: %w", err)
	}

	return nil
}

// NewMagicPacket constructs a magic packet for the given MAC address
func NewMagicPacket(macAddr string) ([]byte, error) {
	macBytes, err := net.ParseMAC(macAddr)
	if err != nil {
		return nil, fmt.Errorf("invalid MAC address: %w", err)
	}

	if len(macBytes) != 6 {
		return nil, fmt.Errorf("invalid MAC address length: %d", len(macBytes))
	}

	// Magic Packet is 6 bytes of 0xFF followed by 16 repetitions of MAC
	buffer := make([]byte, 6+16*6)
	
	// Fill header
	for i := 0; i < 6; i++ {
		buffer[i] = 0xFF
	}

	// Fill body
	for i := 0; i < 16; i++ {
		copy(buffer[6+i*6:], macBytes)
	}

	return buffer, nil
}
