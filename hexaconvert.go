package tgr

import (
	"fmt"
	"strconv"
)

func HexaConvert(n string) string {

	num, err := strconv.ParseInt(n, 16, 64)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(num)

}
