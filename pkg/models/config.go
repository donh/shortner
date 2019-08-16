package models

type databaseStruct struct {
	Account  string `yaml:"account"`
	Database string `yaml:"database"`
	Hostname string `yaml:"hostname"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

// ConfigStruct is the format of config/config.yml
type ConfigStruct struct {
	Database databaseStruct `yaml:"database"`
	Port     int            `yaml:"port"`
}
