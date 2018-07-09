package pacmanconf

type Repository struct {
	Name     string
	Server   []string
	SigLevel []string
	Usage    []string
}

type Config struct {
	RootDir                string
	DBPath                 string
	CacheDir               []string
	HookDir                []string
	GpgDir                 string
	LogFile                string
	HoldPkg                []string
	IgnorePkg              []string
	IgnoreGroup            []string
	Architecture           string
	XferCommand            string
	NoUpgrade              []string
	NoExtract              []string
	CleanMethod            []string
	SigLevel               []string
	LocalFileSigLevel      []string
	RemoteFileSigLevel     []string
	UseSyslog              bool
	Color                  bool
	UseDelta               float64
	TotalDownload          bool
	CheckSpace             bool
	VerbosePkglists        bool
	DisableDownloadTimeout bool
	Repositories           []Repository
}

func (conf *Config) Repository(name string) *Repository {
	for _, repo := range conf.Repositories {
		if repo.Name == name {
			return &repo
		}
	}

	return nil
}
