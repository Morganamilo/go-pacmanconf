package pacmanconf

import (
	"strconv"
	"strings"
)

func parse(data string) *Config {
	var repo *Repository
	lines := strings.Split(data, "\n")
	conf := &Config{UseDelta: -1}
	header := ""

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			runes := []rune(line)
			header = string(runes[1 : len(runes)-1])
			continue
		}

		key, value := splitPair(line)

		if header == "options" {
			setOption(conf, key, value)
		} else {
			if repo == nil || repo.Name != header {
				repo = conf.Repository(header)
				if repo == nil {
					conf.Repositories = append(conf.Repositories, Repository{})
					repo = &conf.Repositories[len(conf.Repositories)-1]
					repo.Name = header
				}
			}

			setRepo(repo, key, value)
		}
	}

	return conf
}

func setRepo(repo *Repository, key string, value string) {
	switch key {
	case "Server":
		repo.Server = append(repo.Server, value)
	case "SigLevel":
		repo.SigLevel = append(repo.SigLevel, value)
	case "Usage":
		repo.Usage = append(repo.Usage, value)
	}
}

func setOption(conf *Config, key string, value string) {
	switch key {
	case "RootDir":
		conf.RootDir = value
	case "DBPath":
		conf.DBPath = value
	case "CacheDir":
		conf.CacheDir = append(conf.CacheDir, value)
	case "HookDir":
		conf.HookDir = append(conf.HookDir, value)
	case "GpgDir":
		conf.GpgDir = value
	case "LogFile":
		conf.LogFile = value
	case "HoldPkg":
		conf.HoldPkg = append(conf.HoldPkg, value)
	case "IgnorePkg":
		conf.IgnorePkg = append(conf.IgnorePkg, value)
	case "IgnoreGroup":
		conf.IgnoreGroup = append(conf.IgnoreGroup, value)
	case "Architecture":
		conf.Architecture = value
	case "XferCommand":
		conf.XferCommand = value
	case "NoUpgrade":
		conf.NoUpgrade = append(conf.NoUpgrade, value)
	case "NoExtract":
		conf.NoExtract = append(conf.NoExtract, value)
	case "CleanMethod":
		conf.CleanMethod = append(conf.CleanMethod, value)
	case "SigLevel":
		conf.SigLevel = append(conf.SigLevel, value)
	case "LocalFileSigLevel":
		conf.LocalFileSigLevel = append(conf.LocalFileSigLevel, value)
	case "RemoteFileSigLevel":
		conf.RemoteFileSigLevel = append(conf.RemoteFileSigLevel, value)
	case "UseSyslog":
		conf.UseSyslog = true
	case "Color":
		conf.Color = true
	case "UseDelta":
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			conf.UseDelta = f
		}
	case "TotalDownload":
		conf.TotalDownload = true
	case "CheckSpace":
		conf.CheckSpace = true
	case "VerbosePkgLists":
		conf.VerbosePkglists = true
	case "DisableDownloadTimeout":
		conf.DisableDownloadTimeout = true
	}
}

func splitPair(line string) (string, string) {
	split := strings.SplitN(line, "=", 2)

	key := strings.TrimSpace(split[0])

	if len(split) == 1 {
		return key, ""
	}

	value := strings.TrimSpace(split[1])
	return key, value
}

func Parse(data string) *Config {
	return parse(data)
}

func PacmanConf(args ...string) (*Config, string, error) {
	stdout, stderr, err :=  pacmanconf(args)

	if err != nil {
		return nil, stderr, err
	}

	conf := parse(stdout)

	return conf, stderr, nil
}
