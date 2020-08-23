package main
import (
	"fmt"
	"time"
	"math/rand"
)


func findMax(arr [8]int) (int,int){
	//fmt.Println(arr)
	var temp int = 0
	var index int 
	for i := 0;i < len(arr);i++{
		
		if arr[i] > temp {
			temp = arr[i] 
		}
	}
	//fmt.Println(temp)
	//fmt.Println(arr)
	for i := 0;i < len(arr);i++{
		if arr[i] == temp {
			index = i
			break
		}
	}
	return temp,index
}


func findMin(arr [8]int) (int,int) {
	//fmt.Println(arr)
	var temp int = arr[0]
	var index int 
	for i := 0;i < len(arr);i++{
		if temp > arr[i]{
			temp = arr[i]
		}
	}
	//fmt.Println(temp)
	//fmt.Println(arr)
	for i := 0;i < len(arr);i++{
		if arr[i] == temp {
			index = i
			break
		}
	}
	return temp,index
}


func findWorstJudge(arr [8]int) (int,int,int,int,int,int,int,int,int,int,int,int) {
	var (
	sum int 
	//finalSum int
	avg int
	worstjudge int
	bestjudge int 
	worstjudgeIndex int
	bestjudgeIndex int
	worstjudgeDiff int
	bestjudgeDiff int
	maxjudge int
	minjudge int
	maxjudgeIndex int
	minjudgeIndex int
	numDiff [8]int
	)
	maxjudge,maxjudgeIndex = findMax(arr)
	minjudge,minjudgeIndex = findMin(arr)
	//fmt.Println(maxjudge,maxjudgeIndex,minjudge,minjudgeIndex)

	for i:=0;i<len(arr);i++ {
		if i == maxjudgeIndex || i == minjudgeIndex {
			continue
		}else{
			//fmt.Printf("%v\t%v\n",i,arr[i])
			sum += arr[i]
		}
		
	} 
	avg = sum / (len(arr) - 2)

	//fmt.Println(avg)
	for i:=0;i<len(arr);i++{
		if avg > arr[i]{
			numDiff[i]=avg - arr[i]
		}else if avg < arr[i] {
			numDiff[i] = arr[i] - avg
		}else  {
			bestjudgeIndex = i
			bestjudge = avg
		}
	}
	//fmt.Println(numDiff)

	worstjudgeDiff,worstjudgeIndex = findMax(numDiff)
	bestjudgeDiff,bestjudgeIndex = findMin(numDiff)
	worstjudge = arr[worstjudgeIndex]
	bestjudge = arr[bestjudgeIndex]
	//for i := 0;i < len(arr);i++{
	//	if i != worstjudgeIndex || i != bestjudgeIndex{
	//		finalSum += arr[i]
	//	}
	//}

	return worstjudge,worstjudgeIndex,worstjudgeDiff,bestjudge,bestjudgeIndex,
	bestjudgeDiff,avg,sum,maxjudge,maxjudgeIndex,minjudge,minjudgeIndex

}


func main() {
	var arr [8]int
	for i := 0;i < len(arr);i++{
		time.Sleep(time.Duration(100)*time.Millisecond)
		rand.Seed(time.Now().UnixNano())
		arr[i] = rand.Intn(100)
	}
	fmt.Println("8为评委打分如下")
	for i := 0;i<len(arr);i++{
		
		fmt.Printf("第%v位评委打分为:%v\n", i+1,arr[i])
	}
	
	valueMax,indexMax := findMax(arr)
	fmt.Printf("其中最高分为%v,由第%v号评委打出\n",valueMax,indexMax+1)
	valueMin,indexMin := findMin(arr)
	fmt.Printf("其中最低分为%v,由第%v号评委打出\n",valueMin,indexMin+1)
	worstjudge,worstjudgeIndex,worstjudgeDiff,bestjudge,bestjudgeIndex,bestjudgeDiff,
	avg,finalSum,valueMax,indexMax,valueMin,indexMin:= findWorstJudge(arr)
	fmt.Printf("去掉一个最高分%v,去掉一个最低分%v,该选手最终得分为%v\n",valueMax,valueMin,
	finalSum)
	fmt.Printf("平均分为%v\t最差评委的评分为%v\t与平均分差距为%v\t评委序号为%v\n",avg,
	worstjudge,worstjudgeDiff,worstjudgeIndex+1)
	fmt.Printf("平均分为%v\t最佳评委的评分为%v\t与平均分差距为%v \t评委序号为%v\n",avg,
	bestjudge,bestjudgeDiff,bestjudgeIndex+1)
	//fmt.Printf("平均分为%v\t最佳评委的评分为%v\t与平均分差距为%v\t评委序号为%v\n",avg,
	//bestjudge,bestjudgeDiff,bestjudgeIndex+1)
}

