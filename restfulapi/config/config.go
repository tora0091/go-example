package config

const HOSTNAME = "localhost"
const PORT = "8080"

const DATABASE_DRIVER = "sqlite3"
const DATA_SOURCE_PATH = "data/"
const DATA_SOURCE_NAME = "restapi-sample.db"

func GetServerAddr() string {
	return HOSTNAME + ":" + PORT
}

func GetDatabaseDriver() string {
	return DATABASE_DRIVER
}

func GetDataSourceName() string {
	return DATA_SOURCE_PATH + DATA_SOURCE_NAME
}
