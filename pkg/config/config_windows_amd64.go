package config

import (
	"path/filepath"
)

func defaultVstPaths() []string {
	programFiles := "C:\\Program Files"
	programFilesX86 := "C:\\Program Files (x86)"

	paths := []string{
		filepath.Join(programFilesX86, "Cakewalk\\VstPlugins"),
		filepath.Join(programFilesX86, "Steinberg\\VstPlugins"),
		filepath.Join(programFilesX86, "Common Files\\VST3"),
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
