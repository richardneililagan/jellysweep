package arr

import "github.com/jon4hz/jellysweep/internal/config"

// Settings holds the global Jellysweep settings shared by every Sonarr and
// Radarr instance. Build it with NewSettings.
type Settings struct {
	DryRun      bool
	CleanupMode config.CleanupMode
	KeepCount   int
}

// NewSettings builds the arr Settings from the Jellysweep config, with the
// config defaults applied.
func NewSettings(cfg *config.Config) Settings {
	return Settings{
		DryRun:      cfg.DryRun,
		CleanupMode: cfg.CleanupMode,
		KeepCount:   cfg.KeepCount,
	}.WithDefaults()
}

// WithDefaults returns s with the config defaults applied to unset fields.
// Clients call it on construction, so Settings built without NewSettings are
// safe too: a zero KeepCount in a keep mode would otherwise delete every
// regular episode.
func (s Settings) WithDefaults() Settings {
	// Delegate to the config getters so the defaults have a single source.
	defaults := config.Config{CleanupMode: s.CleanupMode, KeepCount: s.KeepCount}
	s.CleanupMode = defaults.GetCleanupMode()
	s.KeepCount = defaults.GetKeepCount()
	return s
}
