package usersconfig

type Configuration struct {
	JwtSecret []byte
}

var Config Configuration
