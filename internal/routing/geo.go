package routing

import (
	"net"
	"log"
)
var downloadDBString = "https://download.maxmind.com/geoip/databases/GeoLite2-Country/download?suffix=tar.gz"
var dbSHA256 = "https://download.maxmind.com/geoip/databases/GeoLite2-Country/download?suffix=tar.gz.sha256"
func InitializeGeoDB(dbPath string) {
	
}