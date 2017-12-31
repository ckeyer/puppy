package config

type Config struct {
	PreRun      *Action       `json:"preRun" yaml:"preRun"`
	PreStop     *Action       `json:"preStop" yaml:"preStop"`
	Apps        []*App        `json:"apps" yaml:"apps"`
	SignalBinds []*SignalBind `json:"signalBinds" yaml:"signalBinds"`
}

type App struct {
	Name        string            `json:"name" yaml:"name"`
	Seq         int               `json:"seq" yaml:"seq"`
	Env         map[string]string `json:"env" yaml:"env"`
	PreRun      *Action           `json:"preRun" yaml:"preRun"`
	PostStop    *Action           `json:"postStop" yaml:"postStop"`
	HealthCheck HealthCheck       `json:"healthCheck" yaml:"healthCheck"`
}

type Action struct {
	Cmd  *[]string   `json:"cmd" yaml:"cmd"`
	Http *HTTPAction `json:"http" yaml:"http"`
}

type HTTPAction struct {
	URL     string `json:"url" yaml:"url"`
	Method  string `json:"method" yaml:"method"`
	CAFile  string `json:"caFile" yaml:"caFile"`
	Cert    string `json:"cert" yaml:"cert"`
	CertKey string `json:"certKey" yaml:"certKey"`
}

type SignalBind struct {
	Action `json:",inline" yaml:",inline"`
	Signal int `json:"signal" yaml:"signal"`
}

type HealthCheck struct {
	Action   `json:",inline" yaml:",inline"`
	Interval string `json:"interval" yaml:"interval"`
	Timeout  string `json:"timeout" yaml:"timeout"`
	Retries  int    `json:"retries" yaml:"retries"`
}

func init() {

}
