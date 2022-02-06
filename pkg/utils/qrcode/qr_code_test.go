package qrcode

import (
	"encoding/json"
	"github.com/XC-Zero/yinwan/pkg/model"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	now := time.Now()
	s := struct {
		Hello string
		Name  string
		model.BasicModel
	}{
		"FUCK YOU !",
		"♂ 哲学 ♂!",
		model.BasicModel{
			RecID:     nil,
			CreatedAt: now,
			UpdatedAt: &now,
		},
	}
	marshal, err := json.Marshal(s)
	if err != nil {
		panic(err)

	}
	code, err := CreateQRCode(string(marshal), White, Pink, 1024)
	if err != nil {
		panic(err)

	}

	create, err := os.Create(strconv.Itoa(now.Local().Nanosecond()) + ".png")
	if err != nil {
		panic(err)
	}
	_, err = create.Write(code)
	if err != nil {
		panic(err)

	}

}
