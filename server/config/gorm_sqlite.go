package config

type Sqlite struct {
	Path   string `mapstructure:"path" json:"path" yaml:"path"`         // 数据库路径
	Config string `mapstructure:"config" json:"config" yaml:"config"`   // 高级配置
	Dbname string `mapstructure:"db-name" json:"dbname" yaml:"db-name"` // 数据库名
}
