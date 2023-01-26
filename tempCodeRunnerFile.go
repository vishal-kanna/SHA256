package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func binary(s string) string {
// 	res := ""
// 	for _, c := range s {
// 		res = fmt.Sprintf("%s%.8b", res, c)
// 	}
// 	return res
// }

// func conversion_64bit(dec int64) string {
// 	dec_bin := strconv.FormatInt(dec, 2)
// 	length_of_bin := len(dec_bin)
// 	for i := 0; i < 64-length_of_bin; i++ {
// 		dec_bin = "0" + dec_bin
// 	}
// 	return dec_bin
// }
// func PaddingMessage(msg string) string {
// 	original_msg_len := len(msg)
// 	pad_one := msg + "1"
// 	pad_len := len(pad_one)
// 	k := 0
// 	for ((pad_len+k)-448)%512 != 0 {
// 		k++
// 	}
// 	for i := 0; i < k; i++ {
// 		pad_one += "0"
// 	}
// 	msg_bit := conversion_64bit(int64(original_msg_len))
// 	// fmt.Println("the length of the message into binary of 64 bits", msg_bit)
// 	pad_one = pad_one + msg_bit
// 	return pad_one
// }
// func string_array_64(padded_msg string) [64]string {
// 	var bit_array [64]string
// 	for i := 0; i < 16; i++ {
// 		bit_array[i] = padded_msg[32*i : 32*(i+1)]
// 	}
// 	for i := 16; i < 64; i++ {
// 		bit_array[i] = "00000000000000000000000000000000"
// 	}
// 	return bit_array
// }