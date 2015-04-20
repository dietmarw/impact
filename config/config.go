package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

// If this environment variable is set, we use the file it
// points to as the name of the user's settings file
var envvar = "IMPACT_CONFIG_FILE"

func SettingsFile() string {
	// Look to see if IMPACT_CONFIG_FILE is set.  If so, read the
	// file it points to.
	if os.Getenv(envvar) != "" {
		return os.Getenv(envvar)
	}

	// Otherwise, find out what platform we are on...
	platform := runtime.GOOS

	datadir := ""
	var err error

	switch platform {
	case "windows":
		// On windows, check to see if APPDATA is defined...
		datadir = os.Getenv("APPDATA")
		if datadir == "" {
			datadir, err = homedir.Expand("~/.config")
		}
	case "linux":
		// On windows, check to see if APPDATA is defined...
		datadir = os.Getenv("XDG_CONFIG_HOME")
		if datadir == "" {
			datadir, err = homedir.Expand("~/.config")
		}
	case "darwin":
		datadir, err = homedir.Expand("~/Library/Preferences")
	default:
		log.Printf("Unknown platform %v", platform)
	}

	if err != nil {
		log.Printf("Error expanding directory: %v", err)
	}

	if datadir == "" {
		datadir = "."
	}

	return path.Join(datadir, "impact", "impactrc")
}

func ReadSettings() (Settings, error) {
	settings := SettingsFile()
	log.Printf("Settings file: %s", settings)

	dir, _ := filepath.Abs(path.Join(os.Getenv("GOPATH"), "src", "github.com", "xogeny",
		"impact", "sample_index.json"))

	return Settings{
		//		Indices: []string{"https://impact.modelica.org/impact_data.json"},
		Indices: []string{"file://" + dir},
	}, nil
}
