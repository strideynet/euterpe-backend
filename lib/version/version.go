package version

// These values are set as part of the build process using the -X ldflag

// CommitHash is the last git commit, followed by -dirty if there were uncommitted changes
var CommitHash = "not_set"

// BuildTime is the time that the build ran
var BuildTime = "not_set"
