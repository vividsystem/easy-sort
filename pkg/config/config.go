package config

type Config struct {
	Format map[string]string `yaml:"format" json:"format"`
	Directories map[string]Directory `yaml:"directories" json:"directories"`
}

type Directory struct {
	Filetypes ModuleList `yaml:"filetypes" json:"filetypes"`
	Filenames ModuleList `yaml:"filenames" json:"filenames"`
}

type ModuleList map[string]interface{} 
