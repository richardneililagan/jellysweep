package arr

import (
	"testing"

	"github.com/jon4hz/jellysweep/internal/config"
	"github.com/stretchr/testify/require"
)

func TestNewSettings(t *testing.T) {
	got := NewSettings(&config.Config{
		DryRun:      true,
		CleanupMode: config.CleanupModeKeepSeasons,
		KeepCount:   3,
	})
	require.Equal(t, Settings{DryRun: true, CleanupMode: config.CleanupModeKeepSeasons, KeepCount: 3}, got)
}

func TestNewSettingsAppliesDefaults(t *testing.T) {
	got := NewSettings(&config.Config{})
	require.Equal(t, Settings{CleanupMode: config.CleanupModeAll, KeepCount: 1}, got)
}

func TestWithDefaultsKeepsSetValues(t *testing.T) {
	got := Settings{CleanupMode: config.CleanupModeKeepEpisodes}.WithDefaults()
	require.Equal(t, Settings{CleanupMode: config.CleanupModeKeepEpisodes, KeepCount: 1}, got)
}
