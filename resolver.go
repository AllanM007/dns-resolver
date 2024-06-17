package main

type Mapper struct {
	Domain string
	IPV4   string
	IPV6   string
}

type Header struct {
	Id      uint16
	Flags   uint16
	QdCount uint16
	AnCount uint16
	NsCount uint16
	ArCount uint16
}

func GetDomain(ip string) (string, error) {
	return "", nil
}

func GetIp(domain string) (string, error) {
	return "", nil
}
