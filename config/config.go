package config

func Get() Config {
	return Config {
		SqlDbConfig {
			Name: "badass-city",
			UserName: "root",
			Password: "devpassword",
		},
		KiteServiceConfig {
			Name: "BaddassCity.user",
			Version: "0.0.1",
			Port: 3635,
		},
	}
}
