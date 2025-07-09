import{
	"fmt"
}
func printscore(score []int, name string) {
	sum := 0
	for i:=0; i<21; i++{
		sum += score[i]
	}

	fmt.Printf("%s の現在のスコア：%d\n", name, sum)
// dif={
// 	{+,|,-,-,-,-,-,-,-,-,-,-,-,-,-,+},
// 	{|,1,"",2,"",3,"",,},
// 	{+,-,-,-,-,-,-,-,-,-,-,-,-,-,+},
// 	{|,name,|},
}