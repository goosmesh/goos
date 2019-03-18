package redis

import (
	"fmt"
	"github.com/coredns/coredns/plugin"
	"testing"
)

func TestRedisZone(t *testing.T) {
	zones := []string{"example.net.", "example.com."}

	zone := plugin.Zones(zones).Matches("host3.example.net.")

	fmt.Println(zones)
	fmt.Println(zone)

}