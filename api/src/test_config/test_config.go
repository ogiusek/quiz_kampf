package testconfig

type Configuration struct {
	Api string
}

var Config = Configuration{
	Api: "http://localhost:5000",
}
