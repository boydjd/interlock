package version

var (
	Version = "1.0.0"

	// GitCommit will be overwritten automatically by the build system
	GitCommit = "HEAD"
)

func FullVersion() string {
	return Version + " (" + GitCommit + ")"
}
