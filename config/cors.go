package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://be-latihan-production-89a4.up.railway.app",
	"my-fe-one.vercel.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}