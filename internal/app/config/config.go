package config

type Config struct {
	Env         string
	PublicUrl   string
	HTTP        HTTPConfig
	DB          DBConfig
	EmailSender EmailSenderConfig
}

type HTTPConfig struct {
	Port int
}

type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

type EmailSenderConfig struct {
	EmailFrom string
	Password  string
}
