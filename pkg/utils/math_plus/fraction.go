package math_plus

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
	"strings"
)

//goland:noinspection GoSnakeCaseUsage
const DIV_SYMBOL = "/"

//goland:noinspection GoSnakeCaseUsage
var (
	ZERO_DENOMINATOR = errors.New("Denominator is 0 ! ")
)

var One = Fraction{
	numerator:   1,
	denominator: 1,
}
var Zero = Fraction{
	numerator:   0,
	denominator: 1,
}

// Fraction 分数
type Fraction struct {
	// 分子
	numerator int64
	// 分母
	denominator int64
}

// String 格式化输出
func (f Fraction) String() string {
	return fmt.Sprintf("%v%s%v", f.numerator, DIV_SYMBOL, f.denominator)
}

// Float64 格式化输出
func (f Fraction) Float64() float64 {
	return float64(f.numerator) / float64(f.denominator)
}

// ToRealFraction 约分
func (f Fraction) ToRealFraction() Fraction {
	gcd := GreatestCommonDivisor(f.numerator, f.denominator)
	f.numerator /= gcd
	f.denominator /= gcd
	return f
}

// ToFakeFraction 真分数扩大
func (f Fraction) ToFakeFraction(multiple int64) Fraction {
	return newFraction(f.numerator*multiple, f.denominator*multiple)
}

//Add 加
func (f Fraction) Add(fra Fraction) Fraction {
	if fra.denominator == 0 {
		return fra
	}
	lcm := LeastCommonMultiple(f.denominator, fra.denominator)
	return newFraction(f.numerator*lcm/f.denominator+fra.numerator*lcm/fra.denominator, lcm).ToRealFraction()
}

//Sub 减
func (f Fraction) Sub(fra Fraction) Fraction {
	fra.numerator = -fra.numerator
	return f.Add(fra)
}

//Mul 乘
func (f Fraction) Mul(fra Fraction) Fraction {
	return newFraction(f.numerator*fra.numerator, f.denominator*fra.denominator).ToRealFraction()
}

func (f Fraction) MulInt64(n int64) Fraction {
	return newFraction(f.numerator*n, f.denominator)
}

// Div 除
func (f Fraction) Div(fra Fraction) Fraction {
	return f.Mul(fra.Reverse())
}

// New 创建假分数
func New(n, d int64) (Fraction, error) {
	if d == 0 {
		return Fraction{}, ZERO_DENOMINATOR
	}
	return newFraction(n, d), nil
}

//Reverse 倒数
func (f Fraction) Reverse() Fraction {
	return newFraction(f.denominator, f.numerator)
}

func newFraction(n, d int64) Fraction {
	return Fraction{
		numerator:   n,
		denominator: d,
	}
}

// NewFromFloat 从小数创建
func NewFromFloat(n, d float64) (Fraction, error) {
	if d == 0.0 {
		return Fraction{}, ZERO_DENOMINATOR
	}

	nn, err := FloatToFraction(n)
	if err != nil {
		return Fraction{}, err
	}
	dd, err := FloatToFraction(d)
	if err != nil {
		return Fraction{}, err
	}
	return nn.Div(dd), nil
}

// NewFromFloatByDecimal 使用decimal 计算，结果可能不完全正确
func NewFromFloatByDecimal(n, d float64) Fraction {
	nn, dd := decimal.NewFromFloat(n), decimal.NewFromFloat(d)
	result, _ := nn.Div(dd).Float64()
	f, _ := NewFromFloat(result, 1)
	return f
}

// NewFromString 从字符串创建,字符串中分子分母需为整数
func NewFromString(str string) (Fraction, error) {
	fraction := Fraction{}
	arr := strings.Split(str, "/")
	if len(arr) == 1 {
		n, err := strconv.Atoi(str)
		if err != nil {
			return Fraction{0, 1}, err
		}
		return Fraction{int64(n), 1}, nil
	}

	if len(arr) != 2 {
		return fraction, errors.New(" Fraction string is invalid! ")
	}
	n, err := strconv.ParseInt(arr[0], 10, 64)
	if err != nil {
		return Fraction{}, errors.New("Numerator is invalid! ")
	}
	d, err := strconv.ParseInt(arr[1], 10, 64)
	if err != nil {
		return Fraction{}, errors.New("Denominator is invalid! ")
	}
	if d == 0 {
		return Fraction{}, ZERO_DENOMINATOR
	}
	return newFraction(n, d).ToRealFraction(), nil
}
func FloatToFraction(f float64) (Fraction, error) {
	if f == 0 {
		return newFraction(0, 1), nil
	} else {
		str := fmt.Sprint(f)
		if strings.Contains(str, ".") {
			arr := strings.Split(str, ".")
			n, err := strconv.ParseInt(arr[0]+arr[1], 10, 64)
			if err != nil {
				return Fraction{}, err
			}
			return newFraction(n, int64(math.Pow10(len(arr[1])))), nil

		} else {
			return newFraction(int64(int(f)), 1), nil
		}

	}
}
