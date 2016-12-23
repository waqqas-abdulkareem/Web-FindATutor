package app
type DBConfig struct{
	DriverName string
	DriverSourceName string
}

type Configuration struct{
	DBConfig DBConfig
	Environment	 string
}