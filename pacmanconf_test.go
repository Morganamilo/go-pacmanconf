package pacmanconf

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, field string, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%s expected: %s got %s", field, a, b)
	}
}

func TestPacmanconf(t *testing.T) {
	conf, stderr, err := ParseFile("testdata/pacman.conf")
	if err != nil {
		t.Fatalf("%s %s", stderr, err)
	}

	repo1 := Repository{
		Name:     "repo1",
		Servers:  []string{"foo", "bar"},
		SigLevel: []string{"PackageNever", "DatabaseNever"},
		Usage:    []string{"Sync", "Search"},
	}

	custom := Repository{
		Name:    "custom",
		Servers: []string{"custom"},
		Usage:   []string{"All"},
	}

	expect(t, "RootDir", "/path/to/root/dir", conf.RootDir)
	expect(t, "DBPath", "/path/to/db/dir", conf.DBPath)
	expect(t, "CacheDir", []string{"/path/to/cache/dir", "/path/to/custom"}, conf.CacheDir)
	expect(t, "HookDir", []string{"/path/to/hook/dir"}, conf.HookDir)
	expect(t, "GPGDir", "/path/to/gpg/dir", conf.GPGDir)
	expect(t, "LogFile", "/path/to/log/file", conf.LogFile)
	expect(t, "HoldPkg", []string{"hold", "package"}, conf.HoldPkg)
	expect(t, "IgnorePkg", []string{"ignore", "package"}, conf.IgnorePkg)
	expect(t, "IgnoreGroup", []string{"ignore", "group"}, conf.IgnoreGroup)
	expect(t, "Architecture", []string{"neo-classical1", "neo-classical2"}, conf.Architecture)
	expect(t, "XferCommand", "/path/to/command %u", conf.XferCommand)
	expect(t, "NoUpgrade", []string{"no", "upgrade"}, conf.NoUpgrade)
	expect(t, "NoExtract", []string{"no", "extract"}, conf.NoExtract)
	expect(t, "CleanMethod", []string{"KeepInstalled", "KeepCurrent"}, conf.CleanMethod)
	expect(t, "SigLevel", []string{"PackageRequired", "PackageTrustedOnly", "DatabaseRequired", "DatabaseTrustedOnly"}, conf.SigLevel)
	expect(t, "LocalFileSigLevel", []string{"PackageOptional", "PackageTrustedOnly"}, conf.LocalFileSigLevel)
	expect(t, "RemoteFileSigLevel", []string{"PackageNever"}, conf.RemoteFileSigLevel)
	expect(t, "Color", true, conf.Color)
	expect(t, "UseDelta", 1.1, conf.UseDelta)
	expect(t, "TotalDownload", true, conf.TotalDownload)
	expect(t, "CheckSpace", true, conf.CheckSpace)
	expect(t, "VerbosePkgLists", true, conf.VerbosePkgLists)
	//expect(t, "DisableDownloadTimeout", true, conf.DisableDownloadTimeout)
	//pacman-conf bug: does not output this
	expect(t, "Repositories", []Repository{repo1, custom}, conf.Repos)
}
