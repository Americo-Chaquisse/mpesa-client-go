package mpesa

type Config struct {
	PublicKey           string
	ApiKey              string
	Host                string
	origin              string
	ServiceProviderCode string
}

type Client struct {
	Config Config
}

func (config *Config) SetDefaults() {
	if config.Host == "" {
		config.Host = "https://api.sandbox.vm.co.mz:18352"
	}
	if config.origin == "" {
		config.origin = "developer.mpesa.vm.co.mz"
	}
}
