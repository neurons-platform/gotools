package redis

import "testing"

func TestInitRedisClient(t *testing.T) {
	InitRedisClient("172.30.42.4:7001", "M5zLKFfPRbTdMWtJ4Cz8")
	Rc.Set("name", "neurons-platform", 0)
}
