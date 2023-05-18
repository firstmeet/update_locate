package main

import (
	"net"

	"github.com/oschwald/geoip2-golang"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type BOT struct {
	ID                int      `gorm:"id" json:"id"`
	IP                string   `gorm:"ip" json:"ip"`
	Local             string   `gorm:"local" json:"local"`
	Mac               string   `gorm:"index:mac,unique" json:"mac"`
	Message           string   `gorm:"message" json:"message"`
	Describe          string   `gorm:"describe" json:"describe"`
	Mark              string   `gorm:"mark" json:"mark"`
	Connected         int64    `gorm:"connected" json:"connected"`
	Flags             []string `gorm:"-" json:"flags"`
	FlagString        string   `gorm:"flag" json:"-"`
	Locate            string   `gorm:"locate" json:"locate"`
	Address           string   `gorm:"address" json:"address"`
	Online            int      `gorm:"online" json:"online"`
	Uptime            int      `gorm:"uptime" json:"uptime"`
	Version           string   `gorm:"version" json:"version"`
	User              string   `gorm:"user" json:"user"`
	Netflow           string   `gorm:"netflow" json:"netflow"`
	ShadowsocksServe  string   `json:"shadowsocksServe"`
	Socks5Serve       string   `json:"socks5Serve"`
	TcpForwardServe   string   `json:"tcpForwardServe"`
	UpdForwardServe   string   `json:"updForwardServe"`
	ReverseProxyServe string   `json:"reverseProxyServe"`
}

func main() {
	Database, err := gorm.Open(sqlite.Open("zone.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	Region, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		panic(err)
	}
	var bots []BOT
	Database.Model(&BOT{}).Find(&bots)
	for _, bot := range bots {
		record, err := Region.Country(net.ParseIP(bot.IP))
		if err == nil {
			Database.Model(&BOT{}).Where("id = ?", bot.ID).Updates(BOT{
				Locate:  record.Country.Names["en"],
				Address: record.Country.Names["en"],
			})
		}

	}
}
