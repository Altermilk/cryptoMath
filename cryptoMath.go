package cryptoMath

import (
	"math"
	"strconv"
)

func Gcd(a, b int) (d, x, y int) {
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

func ModInv(c, m int) int{
	_, _, d :=  Gcd(m, c)
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