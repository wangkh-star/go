package config

// 日志
type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"Prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}
