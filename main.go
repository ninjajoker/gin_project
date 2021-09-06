package main

import "gorm.io/gorm"

var (
	GORM   *gorm.DB
	CONFIG Server
)

func main() {
	Viper()
	GORM = Gorm()
	r := setupRouter()
	r.Run(":7070")
}
