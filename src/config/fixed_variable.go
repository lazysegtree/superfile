package variable

import "github.com/adrg/xdg"

var HomeDir = xdg.Home
var SuperFileMainDir = xdg.ConfigHome + "/superfile"
var SuperFileCacheDir = xdg.CacheHome + "/superfile"
var SuperFileDataDir = xdg.DataHome + "/superfile"
var SuperFileStateDir = xdg.StateHome + "/superfile"

const (
	CurrentVersion      string = "v1.1.7.1 [lazysegtree custom build]"
	LatestVersionURL    string = "https://api.github.com/repos/yorukot/superfile/releases/latest"
	LatestVersionGithub string = "github.com/yorukot/superfile/releases/latest"
)

var (
	ThemeFolder      string = SuperFileMainDir + "/theme"
	LastCheckVersion string = SuperFileDataDir + "/lastCheckVersion"
	ThemeFileVersion string = SuperFileDataDir + "/themeFileVersion"
	FirstUseCheck    string = SuperFileDataDir + "/firstUseCheck"
	PinnedFile       string = SuperFileDataDir + "/pinned.json"
	ConfigFile       string = SuperFileMainDir + "/config.toml"
	HotkeysFile      string = SuperFileMainDir + "/hotkeys.toml"
	ToggleDotFile    string = SuperFileDataDir + "/toggleDotFile"
	ToggleFooter    string = SuperFileDataDir + "/toggleFooter"
	LogFile          string = SuperFileStateDir + "/superfile.log"
	FixHotkeys       bool   = false
	FixConfigFile    bool   = false
	LastDir          string = ""
	PrintLastDir     bool   = false
)

const (
	TrashDirectory      string = "/Trash"
	TrashDirectoryFiles string = "/Trash/files"
	TrashDirectoryInfo  string = "/Trash/info"
)
