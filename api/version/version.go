package version

import (
	"runtime"
)

// Version : The main version number that is being run at the moment.
const Version = "0.1.0"

// GoVersion : Current GoLang Version
var GoVersion = runtime.Version()
