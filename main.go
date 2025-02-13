package main

import (
	"flag"
	"fmt"
	"math"
	"net/http"
	"strconv"
)

// 定义命令行参数
var (
	capital            = flag.Float64("capital", 5.0, "初始金额或每年固定存的金额")
	annualInterestRate = flag.Float64("annualInterestRate", 0.08, "年利率")
	year               = flag.Float64("year", 30, "存款年数")
	method             = flag.String("method", "dongCunXi", "存款方式")
)

func main() {
	// 解析命令行参数
	flag.Parse()

	http.HandleFunc("/calculate", calculateHandler)
	fmt.Println("服务器启动，监听端口8080")
	http.ListenAndServe(":8080", nil)

	fmt.Printf("本金或每年定存：%v \n", *capital)
	fmt.Printf("年利率：%v \n", *annualInterestRate) // 修改了这里
	fmt.Printf("存款年数：%v \n", *year)              // 修改了这里
	fmt.Printf("存款方式：%v \n", *method)

	fmt.Printf("\n本金 X 年限 %v * %v = %.2f\n", *capital, *year, *capital*(*year))
	fmt.Printf("动态存钱法-最终金额：%.2f\n", dongCunxi(*year, *capital, *annualInterestRate))
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	// 解析URL参数
	query := r.URL.Query()
	yearStr := query.Get("year")
	capitalStr := query.Get("capital")
	annualInterestRateStr := query.Get("annualInterestRate")
	//method := query.Get("method")

	// 将字符串参数转换为适当的类型
	year, _ := strconv.ParseFloat(yearStr, 64)
	capital, _ := strconv.ParseFloat(capitalStr, 64)
	annualInterestRate, _ := strconv.ParseFloat(annualInterestRateStr, 64)

	fmt.Fprintf(w, "本金或每年定存：%v \n", capital)
	fmt.Fprintf(w, "存款年数：%v \n", year)
	fmt.Fprintf(w, "年利率：%v \n", annualInterestRate)
	fmt.Fprintf(w, "存款方式：每年存款%v，存%v年，年利率%v  \n", capital, year, annualInterestRate)

	fmt.Fprintf(w, "\n本金*年限 %v * %v = %.2f\n", capital, year, capital*(year))
	fmt.Fprintf(w, "复利存钱法-最终金额：%.2f\n", dongCunxi(year, capital, annualInterestRate))
}

// 固定存钱法计算最终金额
func guCunXi(year, capital, annualInterestRate float64) float64 {
	// 正确计算每年的利息
	return capital * math.Pow(1+annualInterestRate, float64(year))
}

// 动态存钱法计算最终金额
func dongCunxi(year, capital, annualInterestRate float64) float64 {
	finalMoney := 0.0 // 从初始金额开始累加
	for i := year; i > 0; i-- {
		curCapitalMoney := guCunXi(i, capital, annualInterestRate)
		finalMoney += curCapitalMoney // 累加增长的部分
	}
	return finalMoney
}
