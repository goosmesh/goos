package alg

import (
	"fmt"
	"testing"
)

func TestRsaEncrypt(t *testing.T) {
	if data, e := RsaEncrypt("hello"); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(data)
	}
}

func TestRsaDecrypt(t *testing.T) {
	if data, e := RsaEncrypt("hello"); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(data)
		if data, e := RsaDecrypt(data); e != nil {
			fmt.Println(e)
		} else {
			fmt.Println(data)
		}
	}

}

func TestRsaDecryptConfig(t *testing.T) {
	data := "ETsjrshbw3tRDx8cyQGmStrbjnvCplfSThN8GWKpowve+VOg4M0dklU000ajR79/OISYBvlxIPYcl5NVy9RqRwc/Lj9sjjgiFx0G42ll0QroTdynkedgcTaRcfSOVnzokKHntbyeDMacwK71DX/rHii6y3yKGBWqeKSelcVsOr1Q3sZes1XBBubnHXIawOw5b0f/ANeVT5VK1vnl2CQSXx5JkX+LlmnCraDgLYmWkGw3WW6pnoSsPq7LgzPRY3nqRDv5vcgig5uHL4FKQbo5oe7ZEDEIySNad4O5Xt8dubGLI/wFHUcjnUZofkPeA6RnkinFodMAFI5ivmzXcEJADg=="

	if data, e := RsaDecrypt(data); e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(data)
	}

}