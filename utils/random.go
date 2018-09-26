package utils

func GetRandomPngFileName() string {
	uuid := GetUUID()
	return uuid + ".png"
}
