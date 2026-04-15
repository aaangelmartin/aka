// Package buildinfo holds values injected at build time via -ldflags.
package buildinfo

// Version is the semver tag (or "dev" when built from a working tree).
var Version = "dev"

// Commit is the short git hash ("none" if unavailable).
var Commit = "none"

// Date is the ISO 8601 build timestamp (UTC).
var Date = "unknown"
