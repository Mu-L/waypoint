package ecs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlatformConfig(t *testing.T) {
	t.Run("empty is fine", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory: 512,
			ALB:    &ALBConfig{},
		}

		require.NoError(t, p.ConfigSet(cfg))
	})

	t.Run("fine if only cert", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory: 512,
			ALB: &ALBConfig{
				CertificateId: "xyz",
			},
		}

		require.NoError(t, p.ConfigSet(cfg))
	})

	t.Run("fine if only listener", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory: 512,
			ALB: &ALBConfig{
				ListenerARN: "xyz",
			},
		}

		require.NoError(t, p.ConfigSet(cfg))
	})

	t.Run("errors if cert and listener are set", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory: 512,
			ALB: &ALBConfig{
				CertificateId: "xyz",
				ListenerARN:   "abc",
			},
		}

		require.Error(t, p.ConfigSet(cfg))
	})

	t.Run("errors if zone_id and fqdn and listener are set", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory: 512,
			ALB: &ALBConfig{
				ZoneId:      "xyz",
				FQDN:        "a.b",
				ListenerARN: "abc",
			},
		}

		require.Error(t, p.ConfigSet(cfg))
	})

	t.Run("errors if zone_id but not fqdn are set", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory: 512,
			ALB: &ALBConfig{
				ZoneId: "xyz",
			},
		}

		require.Error(t, p.ConfigSet(cfg))
	})

	t.Run("errors if fqdn but not zone_id are set", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory: 512,
			ALB: &ALBConfig{
				FQDN: "xyz",
			},
		}

		require.Error(t, p.ConfigSet(cfg))
	})

	t.Run("fine with just zone and fqdn", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory: 512,
			ALB: &ALBConfig{
				ZoneId: "xyz",
				FQDN:   "a.b",
			},
		}

		require.NoError(t, p.ConfigSet(cfg))
	})

	t.Run("errors if internal and listener are set", func(t *testing.T) {
		var p Platform

		i := true
		cfg := &Config{
			Memory: 512,
			ALB: &ALBConfig{
				InternalScheme: &i,
				ListenerARN:    "abc",
			},
		}

		require.Error(t, p.ConfigSet(cfg))
	})

	t.Run("fine with just internal", func(t *testing.T) {
		var p Platform

		i := true
		cfg := &Config{
			Memory: 512,
			ALB: &ALBConfig{
				InternalScheme: &i,
			},
		}

		require.NoError(t, p.ConfigSet(cfg))
	})

	t.Run("allows no memory_reservation", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory: 512,
		}

		require.NoError(t, p.ConfigSet(cfg))
	})

	t.Run("allows memory_reservation same as memory", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory:            512,
			MemoryReservation: 512,
		}

		require.NoError(t, p.ConfigSet(cfg))
	})

	t.Run("allows memory_reservation less than memory", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory:            512,
			MemoryReservation: 256,
		}

		require.NoError(t, p.ConfigSet(cfg))
	})

	t.Run("disallows too small values of memory", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory: 3,
		}

		require.Error(t, p.ConfigSet(cfg))
	})

	t.Run("disallows too small values of memory_reservation", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory:            512,
			MemoryReservation: 3,
		}

		require.Error(t, p.ConfigSet(cfg))
	})

	t.Run("disallows memory_reservation greater than memory", func(t *testing.T) {
		var p Platform

		cfg := &Config{
			Memory:            512,
			MemoryReservation: 513,
		}

		require.Error(t, p.ConfigSet(cfg))
	})
}
