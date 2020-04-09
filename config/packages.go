package config

//Insert your package library here
func GetPackages() []string {
	return []string{
		"github.com/gin-gonic/gin",
		"github.com/ulule/deepcopier",
		"github.com/go-sql-driver/mysql",
		"github.com/jinzhu/gorm",
		"golang.org/x/crypto/bcrypt",
		"github.com/dgrijalva/jwt-go",
		"github.com/joho/godotenv"}
}