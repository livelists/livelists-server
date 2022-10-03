package apiprotobuf

type LoggerConfig struct {
	JSON   bool   `yaml:"json"`
	Level  string `yaml:"level"`
	Sample bool   `yaml:"sample,omitempty"`
}
