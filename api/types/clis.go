package types

type BapiClisT struct {
	TaskID       string
	Verbose      int
	SaveLog      bool
	LogDir       string
	HelpFlags    bool
	Version      string
	Proxy        string
	Retries      int
	Query        string
	Format       string
	Outfn        string
	Email        string
	Thread       int
	From         int
	Size         int
	RemoteName   bool
	Timeout      int
	RetSleepTime int
	PrettyJSON   bool
	Indent       int
	SortKeys     bool
	Extra        string
	XML2json     bool
}

type NcbiClisT struct {
	NcbiDB        string
	NcbiRetmax    int
	NcbiXMLToJSON string
	NcbiXMLPaths  []string
	NcbiKeywords  string
}
