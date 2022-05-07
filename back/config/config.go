package config

type AppConfig struct {
	ServiceName  string `envconfig:"SERVICE_NAME"         default:"dso-back"`
	BindIP       string `envconfig:"BIND_IP"              required:"true"`
	BindPort     int    `envconfig:"BIND_PORT"            required:"true"`
	SaveLocation string `envconfig:"SAVE_LOCATION"        required:"true"`
}

func (cfg AppConfig) Validate() error {
	return nil
}
