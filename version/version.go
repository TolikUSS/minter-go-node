package version

// Version components
const (
	Maj = "0"
	Min = "20"
	Fix = "2"

	AppVer = 4
)

var (
	// Must be a string because scripts like dist.sh read this file.
	Version = "0.20.2"

	// GitCommit is the current HEAD set using ldflags.
	GitCommit string
)

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit
	}
}
