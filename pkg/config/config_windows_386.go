package config

import (
	"path/filepath"
)

func defaultVstPaths() []string {
	programFiles := "C:\\Program Files"

	paths := []string{
		filepath.Join(programFiles, "Cakewalk\\VstPlugins"),
		filepath.Join(programFiles, "Steinberg\\VstPlugins"),
		filepath.Join(programFiles, "Common Files\\VST3"),
	}
	return paths
}

func defaultSamplePaths() []string {
	paths := make([]string, 5)
	return paths
}

func defaultKontaktLibraryPaths() []string {
	paths := make([]string, 5)
	// @todo Set default Kontakt library directories.
	return paths
}
