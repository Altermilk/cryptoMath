package cryptoMath

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strconv"
)

func Gcd(a, b int) (d, x, y int) {
	if(b>a){
		a, b = b, a
	}
	u := [3]int{a, 1, 0}
	v := [3]int{b, 0, 1}
	for {
		if v[0] == 0 {
			return u[0], u[1], u[2]
		}
		q := int(u[0] / v[0])
		t := [3]int{u[0] % v[0], u[1] - q*v[1], u[2] - q*v[2]}
		u = v
		v = t
	}
}

func GcdRunes(a, b int32) (d, x, y int32) {
	if(b>a){
		a, b = b, a
	}
	u := [3]rune{a, 1, 0}
	v := [3]rune{b, 0, 1}
	for {
		if v[0] == 0 {
			return u[0], u[1], u[2]
		}
		q := rune(u[0] / v[0])
		t := [3]int32{u[0] % v[0], u[1] - q*v[1], u[2] - q*v[2]}
		u = v
		v = t
	}
}

func Gcd64(a, b uint64) (d, x, y uint64) {
	if(b>a){
		a, b = b, a
	}
	u := [3]uint64{a, 1, 0}
	v := [3]uint64{b, 0, 1}
	for {
		if v[0] == 0 {
			return u[0], u[1], u[2]
		}
		q := uint64(u[0] / v[0])
		t := [3]uint64{u[0] % v[0], u[1] - q*v[1], u[2] - q*v[2]}
		u = v
		v = t
	}
}

func ModInv(c, m int) int{
	_, _, d :=  Gcd(m, c)
	if d<0 {
		d += m
	}
	return d
}

func ModInvRunes(c, m int32) int32{
	_, _, d :=  GcdRunes(m, c)
	if d<0 {
		d += m
	}
	return d
}

func Modularizate(a, x, p int) int{
	t := int(math.Floor(math.Log2(float64(x))))
	xStr := strconv.FormatInt(int64(x), 2)
	x0b := make([]int, t + 1)
	for i := 0; i<t + 1; i++{
		x0b[i], _ = strconv.Atoi(string(xStr[i]))
	}
	
	y := 1
	for i:=0; i<t + 1; i++{
		y = (y*y)%p
		if(x0b[i] == 1){
			y = (y*a)%p
		}
	}
	return y
}
func ModularizateInfo(a, x, p int){
	t := int(math.Floor(math.Log2(float64(x))))
	fmt.Println("t = ", t,  " | math.Floor(math.Log2(float64(x)))")
	xStr := strconv.FormatInt(int64(x), 2)
	fmt.Println("X (", x, ")-> 0b = ",xStr )
	x0b := make([]int, t + 1)
	for i := 0; i<t + 1; i++{
		x0b[i], _ = strconv.Atoi(string(xStr[i]))
	}
	fmt.Println(x0b)
	y := 1
	for i:=0; i<t + 1; i++{
		y = (y*y)%p
		fmt.Println("y = (y*y)mod p = ", y)
		if(x0b[i] == 1){
			y = (y*a)%p
			fmt.Println("1 y = (y*a)mod p = ", y)
		}
	}
	fmt.Println("---- Y = ", y)
}
func ModularizateRune(a int32, x int32, p int32) int32 {
	t := int(math.Floor(math.Log2(float64(x))))
	xStr := strconv.FormatInt(int64(x), 2)
	x0b := make([]int, t + 1)
	for i := 0; i < t+1; i++ {
	  x0b[i], _ = strconv.Atoi(string(xStr[i]))
	}
  
	var y int32 = 1
	for i := 0; i < t+1; i++ {
	  y = (y * y) % p
	  if x0b[i] == 1 {
		y = (y * a) % p
	  }
	}
	return y
  }

func Modularizate64(a, x, p uint64) uint64{
	t := int(math.Floor(math.Log2(float64(x))))
	xStr := strconv.FormatInt(int64(x), 2)
	x0b := make([]int, t + 1)
	for i := 0; i<t + 1; i++{
		x0b[i], _ = strconv.Atoi(string(xStr[i]))
	}
	
	var y uint64 = 1
	for i:=0; i<t + 1; i++{
		y = (y*y)%p
		if(x0b[i] == 1){
			y = (y*a)%p
		}
	}
	return y
}

func Gamma(info, key [] byte) [] byte { // функция наложения гаммы
	k := make([]byte, len(info))
	j := 0
	for i := 0; i < len(info); i++ { // создаем гамму той же длины и записываем в нее ключ
	  k[i] = key[j]
	  if(j == len(key) - 1){  // если текст длиннее ключа, то дублируем гамму до его конца
		j = 0
	  }else{
		j++
	  }
	}
  
	return XOR(info, k) 
  }

  func XOR(a, b []byte) []byte{ 

	result := make([]byte, len(a))
  
	for i := range a{
	  result[i] = a[i]^b[i]
	}
  
	return result
  }

func GetRandomSimpleNum(rnd rand.Rand) int{
	num := rnd.Intn(len(SimpleNums))
	return SimpleNums[num]
}

func GetRandomSimpleNum64(a, b int, rnd rand.Rand) int64{
	if b > SimpleNums[len(SimpleNums)-1]{
		return int64(SimpleNums[len(SimpleNums)-1])
	}
	for{
		num := rnd.Intn(len(SimpleNums))
		if SimpleNums[num]<b && SimpleNums[num] > a{
			return int64(SimpleNums[num])
		}
	}
	
}

func GetRandomSimpleNumU64(rnd rand.Rand) uint64{
	num := rnd.Intn(len(SimpleNums))
	return uint64(SimpleNums[num])
}

func GetRandomBigSimpleNumU64(rnd rand.Rand) uint64{
	num := rnd.Intn(len(BigSimpleNumsUint))
	return BigSimpleNumsUint[num]
}

func GetRandomSimpleNumBIG(a, b *big.Int, rnd rand.Rand) *big.Int{
	tmp := new(big.Int)
	tmp.SetString(BigSimpleNums[len(BigSimpleNums)-1], 10)
	if b.Cmp(tmp) == 1{
		return tmp
	}
	for{
		num := rnd.Intn(len(BigSimpleNums))
		tmp.SetString(BigSimpleNums[num], 10)
		if tmp.Cmp(b) == -1 && tmp.Cmp(a) == 1{
			return tmp
		}
	}
}