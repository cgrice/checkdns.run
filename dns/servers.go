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
		Server{ Name: "Liquid Telecommunications Ltd", Address: "5.11.11.5", Country: "za", Location: "Cullinan, South Africa"},
		Server{ Name: "Deutsche Telekom AG", Address: "195.243.214.4", Country: "de", Location: "Oberhausen, Germany"},
		Server{ Name: "Claro S.A", Address: "200.248.178.54", Country: "br", Location: "Santa Cruz do Sul, Brazil"},
		Server{ Name: "ICONZ Ltd", Address: "210.48.77.68", Country: "nz", Location: "Auckland, NZ"},
		Server{ Name: "LG Dacom Corporation", Address: "164.124.101.2", Country: "kr", Location: "Seoul, South Korea"},
		Server{ Name: "Mahanagar Telephone Nigam Limited", Address: "203.94.227.70", Country: "in", Location: "Mumbai, India"},
		Server{ Name: "Cogecodata", Address: "66.199.45.225", Country: "ca", Location: "Toronto, Canada"},
		Server{ Name: "Online S.A.S.", Address: "163.172.107.158", Country: "fr", Location: "Paris, France"},
		Server{ Name: "Linode LLC", Address: "172.104.49.100", Country: "sg", Location: "Singapore"},
		Server{ Name: "Ancar B Technologies Ltd", Address: "194.145.240.6", Country: "gb", Location: "Manchester, United Kingdom"},
	}
}
