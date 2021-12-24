package dns

type Server struct {
	Address string
	Name string
	Country string
	Location string
}

func GetServers() []Server {
	return []Server{
		Server{ Name: "Google", Address: "8.8.8.8", Country: "us", Location: "Mountain View CA, USA"},
		Server{ Name: "OpenDNS", Address: "208.67.222.220", Country: "us", Location: "Holtsville NY, USA"},
		Server{ Name: "Quad9", Address: "9.9.9.9", Country: "us", Location: "Mountain View"},
		Server{ Name: "CloudFlare", Address: "1.1.1.1", Country: "au", Location: "Research, Australia"},
		Server{ Name: "YANDEX", Address: "77.88.8.8", Country: "ru", Location: "St Petersburg, Russia"},
		Server{ Name: "ShenZhen Sunrise Technology", Address: "202.46.34.74", Country: "cn", Location: "Guangzhou, China"},
	}
}
