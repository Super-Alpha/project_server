package data

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

func calculateShareFee(minute int, totalMinute int, totalFee float64) (float64, error) {
	//分摊比例
	var r float64
	if totalMinute != 0 {
		r = float64(minute) / float64(totalMinute)
	} else {
		r = 0
	}
	//分摊费用
	shareFee, err := strconv.ParseFloat(fmt.Sprintf("%0.2f", totalFee*r), 64)
	if err != nil {
		return 0, err
	}
	return shareFee, nil
}

func calculateShare(minute int, totalMinute int, totalFee float64) (float64, error) {

	if totalMinute == 0 {
		return 0, nil
	}

	shareFee, _ := decimal.NewFromInt(int64(minute)).Div(decimal.NewFromInt(int64(totalMinute))).Mul(decimal.NewFromFloat(totalFee)).Float64()

	//s, err := strconv.ParseFloat(fmt.Sprintf("%0.2f", shareFee), 64)
	//if err != nil {
	//	return 0, err
	//}

	return shareFee, nil
}

func Main1() {

	var totalFee = 0.003
	var m = []int{15193, 181, 1568}
	var totalMinute = 16942

	//for _,v:= range m{
	//	fee,_:=calculateShareFee(v,totalMinute,totalFee)
	//	fmt.Println("fee:", fee)
	//}

	for _, v := range m {
		fee, _ := calculateShare(v, totalMinute, totalFee)
		fmt.Println("fees:", fee)
	}
}

func add(v1, v2 float64) float64 {

	res := decimal.Sum(decimal.NewFromFloat(v1), decimal.NewFromFloat(v2)).InexactFloat64()

	return res
}

func main() {
	f1 := 1129.6
	fmt.Println((f1 * 100)) //输出：112959.99999999999

	var f2 float64 = 1129.6
	fmt.Println((f2 * 100)) //输出：112959.99999999999

	m1 := 8.2
	m2 := 3.8
	fmt.Println(m1 - m2) // 期望是4.4，结果打印出了4.399999999999999

	v1 := 824.43
	v2 := 12.40
	fmt.Println(add(v1, v2))
}
