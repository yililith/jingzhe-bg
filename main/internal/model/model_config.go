package model

type Conf struct {
	Server   serviceConfig  `mapstructure:"server"`
	Database databaseConfig `mapstructure:"database"`
	// Minio    minioConfig    `mapstructure:"minio"`
}

type serviceConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type databaseConfig struct {
	DBName   string `mapstructure:"db_name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

type minioConfig struct {
	EndPoInt        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"accessKeyID"`
	SecretAccessKey string `mapstructure:"secretAccessKey"`
	BucketName      string `mapstructure:"bucketName"`
	FilePath        string `mapstructure:"filePath"`
	UseSLL          bool   `mapstructure:"useSLL"`
}
