package strategy

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"testing"
)

func TestMd5(t *testing.T) {
	str := "99999999999999999999"
	md5Str1 := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	fmt.Println()
	//hex.EncodeToString
	hasher := md5.New()
	hasher.Write([]byte(str))
	md5Str2 := hex.EncodeToString(hasher.Sum(nil))

	if md5Str1 != md5Str2 {
		t.Fatal("Algorithm issue ...")
	}
}

func TestStrategy(t *testing.T) {
	outputStrategy := GetOutputStrategy("console")
	outputStrategy.Draw()

	outputStrategy = GetOutputStrategy("img")
	outputStrategy.Draw()

	filename := fmt.Sprintf("%x.jpg", md5.Sum([]byte("img")))
	err := os.Remove(filename)
	if err != nil {
		t.Fatal("It didn't draw jpg image")
	}
}
