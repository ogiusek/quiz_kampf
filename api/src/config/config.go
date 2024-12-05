package config

type configuration struct {
	Port             int
	ConnectionString string
}

var Config configuration = configuration{
	Port:             5000,
	ConnectionString: "host=localhost port=5432 user=username password=password dbname=database",
}
