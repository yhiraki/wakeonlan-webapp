package config

import (
	"fmt"
	"regexp"
	"strings"
)

type Target struct {
	Name string
	MAC  string
}

// macRegex validates standard MAC address format (XX:XX:XX:XX:XX:XX or XX-XX-XX-XX-XX-XX)
var macRegex = regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`)

func ParseTargets(args []string) ([]Target, error) {
	var targets []Target
	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid format '%s', expected name=mac", arg)
		}

		name := strings.TrimSpace(parts[0])
		mac := strings.TrimSpace(parts[1])

		if name == "" {
			return nil, fmt.Errorf("name cannot be empty in '%s'", arg)
		}

		if !macRegex.MatchString(mac) {
			return nil, fmt.Errorf("invalid MAC address '%s'", mac)
		}

		targets = append(targets, Target{
			Name: name,
			MAC:  mac,
		})
	}
	return targets, nil
}
