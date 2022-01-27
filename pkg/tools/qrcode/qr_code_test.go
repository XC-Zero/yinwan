package qrcode

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	s := struct {
		Hello string
		name  string
	}{
		"FUCK YOU !",
		"♂ 哲学 ♂!",
	}
	marshal, err := json.Marshal(s)
	if err != nil {
		panic(err)

	}
	code, err := CreateQRCode(string(marshal), Orange, LightPurple, 256)
	if err != nil {
		panic(err)

	}

	create, err := os.Create(strconv.Itoa(time.Now().Local().Nanosecond()) + ".png")
	if err != nil {
		panic(err)
	}
	_, err = create.Write(code)
	if err != nil {
		panic(err)

	}

}
