package main

import (
	"fmt"
	"strconv"
)

var h []string = []string{
	"6a09e667", "bb67ae85", "3c6ef372", "a54ff53a",
	"510e527f", "9b05688c", "1f83d9ab", "5be0cd19",
}

var k []string = []string{
	"428a2f98", "71374491", "b5c0fbcf", "e9b5dba5",
	"3956c25b", "59f111f1", "923f82a4", "ab1c5ed5",
	"d807aa98", "12835b01", "243185be", "550c7dc3",
	"72be5d74", "80deb1fe", "9bdc06a7", "c19bf174",
	"e49b69c1", "efbe4786", "0fc19dc6", "240ca1cc",
	"2de92c6f", "4a7484aa", "5cb0a9dc", "76f988da",
	"983e5152", "a831c66d", "b00327c8", "bf597fc7",
	"c6e00bf3", "d5a79147", "06ca6351", "14292967",
	"27b70a85", "2e1b2138", "4d2c6dfc", "53380d13",
	"650a7354", "766a0abb", "81c2c92e", "92722c85",
	"a2bfe8a1", "a81a664b", "c24b8b70", "c76c51a3",
	"d192e819", "d6990624", "f40e3585", "106aa070",
	"19a4c116", "1e376c08", "2748774c", "34b0bcb5",
	"391c0cb3", "4ed8aa4a", "5b9cca4f", "682e6ff3",
	"748f82ee", "78a5636f", "84c87814", "8cc70208",
	"90befffa", "a4506ceb", "bef9a3f7", "c67178f2",
}

func binary(s string) string {
	res := ""
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

func conversion_64bit(dec int64) string {
	dec_bin := strconv.FormatInt(dec, 2)
	length_of_bin := len(dec_bin)
	for i := 0; i < 64-length_of_bin; i++ {
		dec_bin = "0" + dec_bin
	}
	return dec_bin
}
func conversion(d int64, b int) string {
	dec_bin := strconv.FormatInt(d, 2)
	length_of_bin := len(dec_bin)
	for i := 0; i < b-length_of_bin; i++ {
		dec_bin = "0" + dec_bin
	}
	return dec_bin
}
func binary_32bit(num int64) string {
	return conversion(num, 32)
}
func HextoDecimal(h string) int64 {
	s, _ := strconv.ParseInt(h, 16, 64)
	return s
}

func StringToHexa(a [32]string) string {
	var b uint64
	var d string
	for i := 0; i < 32; i++ {
		b, _ = strconv.ParseUint(a[i], 2, 64)
		c := fmt.Sprintf("%x", b)
		d = d + c
	}
	return d
}
func splitt(b string) [4]string {
	var a [4]string
	for j := 0; j < 4; j++ {
		a[j] = b[(j * 8):(j*8 + 8)]
	}
	return a
}
func finalConver(b [8]string) string {
	var a [32]string
	var d [4]string
	j := 0
	for i := 0; i < 8; i++ {
		d = splitt(b[i])
		for k := 0; k < 4; k++ {
			a[j] = d[k]
			j++
		}
	}
	e := StringToHexa(a)
	return e
}
func PaddingMessage(msg string) string {
	original_msg_len := len(msg)
	pad_one := msg + "1"
	pad_len := len(pad_one)
	k := 0
	for ((pad_len+k)-448)%512 != 0 {
		k++
	}
	for i := 0; i < k; i++ {
		pad_one += "0"
	}
	msg_bit := conversion_64bit(int64(original_msg_len))
	// fmt.Println("the length of the message into binary of 64 bits", msg_bit)
	pad_one = pad_one + msg_bit
	return pad_one
}

//	func string_array_64(padded_msg string) [64]string {
//		var bit_array [64]string
//		for i := 0; i < 16; i++ {
//			bit_array[i] = padded_msg[32*i : 32*(i+1)]
//		}
//		for i := 16; i < 64; i++ {
//			bit_array[i] = "00000000000000000000000000000000"
//		}
//		return bit_array
//	}
func RightRotate(a string, i int) string {
	return a[32-i:32] + a[:32-i]
}
func RightShift(a string, i int) string {
	var res string
	for j := 0; j < i; j++ {
		res = "0" + res
	}
	res = res + a[:32-i]
	return res
}
func xorr(s1, s2 string) string {

	var res string
	for i := 0; i < 32; i++ {
		if s1[i] == s2[i] {
			res = res + "0"
		} else {
			res = res + "1"
		}
	}
	return res
}
func binaryStringAnd(s1, s2 string) string {
	var res string
	for i := 0; i < 32; i++ {
		if s1[i] == '1' && s2[i] == '1' {
			res = res + "1"
		} else {
			res = res + "0"
		}
	}
	return res
}
func binaryStringNot(s string) string {
	var res string
	for i := 0; i < 32; i++ {
		if s[i] == '1' {
			res = res + "0"
		} else {
			res = res + "1"
		}
	}
	return res
}
func ConsToBin() [64]string {
	var g [64]string
	for i := 0; i < 64; i++ {
		n := binary_32bit(HextoDecimal(k[i]))
		g[i] = n
	}
	return g
}
func binaryAdditionOfStrings(s string, a string) string {
	var b string
	k := len(a) - 1
	j := "0"
	for i := len(s) - 1; i > -1; i-- {
		if k < 0 || i < 0 {
			break
		}
		if s[i] == '1' && a[k] == '1' {
			if j == "1" {
				b = "1" + b
				j = "1"
			} else {
				b = "0" + b
				j = "1"
			}
		} else if s[i] == '0' && a[k] == '0' {
			if j == "1" {
				b = "1" + b
				j = "0"
			} else {
				b = "0" + b
				j = "0"
			}
		} else {
			if j == "1" {
				b = "0" + b
				j = "1"
			} else if j == "0" {
				b = "1" + b
				j = "0"
			}
		}
		if (i == 0 || k == 0) && j == "1" {
			b = "1" + b
			j = "1"

		}
		k--
	}
	c := b[len(b)-32 : len(b)]
	return c
}

// func binaryAdditionOfStrings(string1, string2 string) string {
// 	// checking if the length of the first string is greater then
// 	// second then calling the function by swapping the parameters
// 	if len(string1) > len(string2) {
// 		return binaryAdditionOfStrings(string2, string1)
// 	}
// 	// finding the difference between the length of the strings
// 	difference := len(string2) - len(string1)

// 	// making both strings equal by adding 0 in front of a smaller string
// 	for i := 0; i < difference; i++ {
// 		string1 = "0" + string1
// 	}
// 	// initializing a variable carry to keep the track of carry after

// 	// each addition
// 	carry := "0"

// 	// In this variable we will store our final string
// 	answer := ""

// 	// traversing the strings and adding them by picking the index from the end /*
// 	/* For example, we are adding “100” and ”110”.
// 	   So, for the last characters in the string i.e “0” and “0” the first else
// 	   if condition will run.
// 	      Then for the middle characters i.e “0” and “1” the last else if
// 	   condition will run and
// 	      for the first characters i.e “1” and “1” the first if condition will run.
// 	*/
// 	for i := len(string1) - 1; i >= 0; i-- {
// 		if string1[i] == '1' && string2[i] == '1' {
// 			if carry == "1" {
// 				answer = "1" + answer
// 			} else {
// 				answer = "0" + answer
// 				carry = "1"
// 			}
// 		} else if string1[i] == '0' && string2[i] == '0' {
// 			if carry == "1" {
// 				answer = "1" + answer
// 				carry = "0"
// 			} else {
// 				answer = "0" + answer
// 			}
// 		} else if string1[i] != string2[i] {
// 			if carry == "1" {
// 				answer = "0" + answer
// 			} else {
// 				answer = "1" + answer
// 			}
// 		}
// 	}
// 	if carry == "1" {
// 		answer = "1" + answer
// 	}
// 	return answer
// }

// func addition_of_strings(s1, s2 string) string {

// }
func compression(a [8]string, ar [64]string, c [64]string) string {
	b := a
	//  S1 := (e rightrotate 6) xor (e rightrotate 11) xor (e rightrotate 25)
	//     ch := (e and f) xor ((not e) and g)
	//     temp1 := h + S1 + ch + k[i] + w[i]
	//     S0 := (a rightrotate 2) xor (a rightrotate 13) xor (a rightrotate 22)
	//     maj := (a and b) xor (a and c) xor (b and c)
	//     temp2 := S0 + maj

	//     h := g
	//     g := f
	//     f := e
	//     e := d + temp1
	//     d := c
	//     c := b
	//     b := a
	//     a := temp1 + temp2
	for i := 0; i < 64; i++ {
		s1 := xorr(xorr(RightRotate(a[4], 6), RightRotate(a[4], 11)), RightRotate(a[4], 25))
		ch := xorr(binaryStringAnd(a[4], a[5]), binaryStringAnd(binaryStringNot(a[4]), a[6]))
		temp1 := binaryAdditionOfStrings(binaryAdditionOfStrings(binaryAdditionOfStrings(binaryAdditionOfStrings(a[7], s1), ch), ar[i]), c[i])
		s0 := xorr(xorr(RightRotate(a[0], 2), RightRotate(a[0], 13)), RightRotate(a[0], 22))
		maj := xorr(xorr(binaryStringAnd(a[0], a[1]), binaryStringAnd(a[0], a[2])), binaryStringAnd(a[1], a[2]))
		temp2 := binaryAdditionOfStrings(s0, maj)

		a[7] = a[6]
		a[6] = a[5]
		a[5] = a[4]
		a[4] = binaryAdditionOfStrings(a[3], temp1)
		a[3] = a[2]
		a[2] = a[1]
		a[1] = a[0]
		a[0] = binaryAdditionOfStrings(temp1, temp2)

		// Add the compressed chunk to the current hash value:
		// h0 := h0 + a
		// h1 := h1 + b
		// h2 := h2 + c
		// h3 := h3 + d
		// h4 := h4 + e
		// h5 := h5 + f
		// h6 := h6 + g
		// h7 := h7 + h
	}
	for i := 0; i < 8; i++ {
		b[i] = binaryAdditionOfStrings(b[i], a[i])
	}
	ans := finalConver(b)
	return ans

}

// func binary_32bit()
func UserSha256(input string) string {
	// message := "hello w
	data := binary(input)
	Padded_data := PaddingMessage(data)
	// fmt.Println("the length of the padded string is", len(Padded_data))
	// fmt.Println("the padded string is ", Padded_data)
	//converting the string bit into array of each entry 32 bit
	var bit_array [64]string
	for i := 0; i < 16; i++ {
		bit_array[i] = Padded_data[32*i : 32*(i+1)]
	}
	for i := 16; i < 64; i++ {
		bit_array[i] = "00000000000000000000000000000000"
	}
	// fmt.Println("the string to array is", bit_array)
	//applying the changes to the array of bits

	// 	for i from 16 to 63
	//         s0 := (w[i-15] rightrotate  7) xor (w[i-15] rightrotate 18) xor (w[i-15] rightshift  3)
	//         s1 := (w[i-2] rightrotate 17) xor (w[i-2] rightrotate 19) xor (w[i-2] rightshift 10)
	//         w[i] := w[i-16] + s0 + w[i-7] + s1
	// //

	for i := 16; i < 64; i++ {
		s01 := RightRotate(bit_array[i-15], 7)
		// fmt.Println("the s01 string is", s01)
		s02 := RightRotate(bit_array[i-15], 18)
		// fmt.Println("the s02 result is", s02)
		s03 := RightShift(bit_array[i-15], 3)
		// s0 := RightRotate(bit_array[i-15], 7)
		// fmt.Println("the s03 result is", s03)

		xor1 := xorr(s01, s02)
		// fmt.Println("the xorr is", xor1)
		s0 := xorr(xor1, s03)
		// fmt.Println("the s0 result is ", s0)
		s11 := RightRotate(bit_array[i-2], 17)
		s12 := RightRotate(bit_array[i-2], 19)
		s13 := RightShift(bit_array[i-2], 10)
		xor2 := xorr(s11, s12)
		s1 := xorr(xor2, s13)
		bit_array[i] = binaryAdditionOfStrings(binaryAdditionOfStrings(binaryAdditionOfStrings(bit_array[i-16], s0), bit_array[i-7]), s1)
	}
	// fmt.Println("the array is after modification", bit_array)

	// Initialize working variables to current hash value:
	// a := h[0]
	// b := h[1]
	// c := h[2]
	// d := h[3]
	// e := h[4]
	// f := h[5]
	// g := h[6]
	// h := h[7]
	var hh [8]string
	for k := 0; k < 8; k++ {
		hh[k] = binary_32bit(HextoDecimal(h[k]))
	}
	// fmt.Println("the hash or buffer array is", hh)
	k1 := ConsToBin()
	// fmt.Println("the kq value is", k1)
	e := compression(hh, k1, bit_array)
	// return ""
	return e
}
func main() {

	// fmt.Println(UserSha256("a"))
	// s := "vishal"
	// h := sha256.New()
	// h.Write([]byte(s))
	// bs := h.Sum(nil)
	// fmt.Printf("%x", bs)
}
