package main

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "github.com/spf13/cobra"
	"learn.go/chapter02/015.fatRate.refactor/calculator"
)

func main() {
	// 录入
	var (
		name   string
		sex    string
		tall   float64
		weight float64
		age    int
	)

	cmd := &cobra.Command{
		Use:   "healthcheck",
		Short: "体脂计算器，根据身高、体重、性别、年龄计算出体脂比，并给出健康建议",
		Long:  "该体脂计算器是基于BMI的体脂计算器......",
		Run: func(cmd *cobra.Command, args []string) {
			// 回调
			fmt.Println("name: ", name)
			fmt.Println("sex: ", sex)
			fmt.Println("tall: ", tall)
			fmt.Println("weight: ", weight)
			fmt.Println("age: ", age)

			// 计算
			bmi := calculator.CalcBMI(tall, weight)
			fatRate := calculator.CalcFatRate(bmi, age, sex)
			fmt.Println("fatRate: ", fatRate)
			// 评估结果
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "姓名")
	cmd.Flags().Float64Var(&tall, "tall", 0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0, "体重")
	cmd.Flags().IntVar(&age, "age", 0, "年龄")

	cmd.Execute()
}
