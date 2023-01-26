package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

type TestData struct {
	in  string
	out string
}

func TestSha256(t *testing.T) {
	da := []TestData{
		{in: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", out: ""},
		{"ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb", "a"},
		{"fb8e20fc2e4c3f248c60c39bd652f3c1347298bb977b8b4d5903b85055620603", "ab"},
		{"ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad", "abc"},
		{"88d4266fd4e6338d13b845fcf289579d209c897823b9217da3e161936f031589", "abcd"},
		{"36bbe50ed96841d10443bcb670d6554f0a34b761be67ec9c4a8ad2c0c44ca42c", "abcde"},
		{"bef57ec7f53a6d40beb640a780a639c83bc29ac8a9816f1fc6c5c6dcd93c4721", "abcdef"},
		{"7d1a54127b222502f5b79b5fb0803061152a44f92b37e23c6527baf665d4da9a", "abcdefg"},
		{"9c56cc51b374c3ba189210d5b6d4bf57790d351c96c47c02190ecf1e430635ab", "abcdefgh"},
		{"19cc02f26df43cc571bc9ed7b0c4d29224a3ec229529221725ef76d021c8326f", "abcdefghi"},
		{"72399361da6a7754fec986dca5b7cbaf1c810a28ded4abaf56b2106d06cb78b0", "abcdefghij"},
		{"a144061c271f152da4d151034508fed1c138b8c976339de229c3bb6d4bbb4fce", "Discard medicine more than two years old."},
	}
	for i := 0; i < len(da); i++ {
		want := sha256.New()
		want.Write([]byte(da[i].out))
		// fmt.Println("the want string is", want, len(want))
		// bs := want.Sum(nil)
		// fmt.Printf(" want is %x", bs)
		// w := bytes.NewBuffer(bs).String()\
		// w := fmt.Sprintf("%s", bs)
		// fmt.Println("string", string(w))
		sha1_hash := hex.EncodeToString(want.Sum(nil))
		fmt.Println(sha1_hash)
		// str := string(bs[:])
		// fmt.Println(str)
		// // fmt.Println("want is", w)
		got := UserSha256(da[i].out)
		fmt.Println("what we got", got)
		// fmt.Println("the output is ", got, len(got))
		if sha1_hash != got {
			t.Error("the output doesnot match ...want", sha1_hash == got)
		}
	}
}
