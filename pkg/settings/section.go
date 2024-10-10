package settings

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Mysql MySQLSetting `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis RedisSetting `mapstructure:"redis"`
}

type ServerSetting struct {
	Port int `mapstructure:"port"`
}

type MySQLSetting struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname string `mapstructure:"dbname"`
	MaxIdleConns int `mapstructure:"MaxIdleConns"`
	MaxOpenConns int `mapstructure:"MaxOpenConns"`
	ConnMaxLifeTime int `mapstructure:"ConnMaxLifeTime"`
}

type LoggerSetting struct {
	Log_level string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_backups int `mapstructure:"max_backups"`
	Max_age int `mapstructure:"max_age"`
	Max_size int `mapstructure:"max_size"`
	Compress bool `mapstructure:"compress"`
}

type RedisSetting struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int `mapstructure:"database"`
	PoolSize int `mapstructure:"pool_size"`
}