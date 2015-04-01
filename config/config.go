package config

func Get() Config {
	return Config {
		SqlDbConfig {
			Name: "BadassCity.Users",
			UserName: "root",
			Password: "",
			Host: "http://localhost",
		},
		KiteServiceConfig {
			Name: "BaddassCity.user",
			Version: "0.0.1",
			Port: 3635,
		},
	}
}
