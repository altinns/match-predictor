package main

import(
	"fmt"
)

func fivelastG(teamOneW,teamOneD,teamOneL  float32 )float32{

wins:=teamOneW*3
Draw:=teamOneD*1
Losses:=teamOneL*0
so:=wins+Draw+Losses
 resulto:=so/15
return resulto*100
}

func goalavg(scored,conced,mostScore,mostConced  float32 ) {

	scorePlus:=scored/mostScore*100
	scoreMinus:=conced/mostConced *100
	
	fmt.Printf("score rate: %v   conced :%v  \n",scorePlus,scoreMinus)

}

func tableranking(teamsnr,teamPosition float32){

	result:=teamPosition/teamsnr
	result=result*100

	fmt.Printf("chances of winning the leuage are:%v \n",result)
	 



}




func main(){
	var teamOneW,teamOneD,teamOneL float32 
	var teamOne string
fmt.Print("enter team 1:")
//5 last games
fmt.Scan(&teamOne)
fmt.Printf("out of 5 games how many wins:\n")
fmt.Scan(&teamOneW)
fmt.Printf("out of 5 games how many Draws:\n")
fmt.Scan(&teamOneD)
fmt.Printf("out of 5 games how many Losses:\n")
fmt.Scan(&teamOneL)
fivelastG(teamOneW,teamOneD,teamOneL)
 result:=fivelastG(teamOneW,teamOneD,teamOneL)


//goalavg
var mostScore,mostConced,scored,conced float32
 fmt.Println("Most leuage goals scored:")
 fmt.Scan(&mostScore)
 fmt.Println("Most leuage goals conceded:")
  fmt.Scan(&mostConced)
  fmt.Println("Team goals scored:")
   fmt.Scan(&scored)
   fmt.Println("Team goals conced:")
   fmt.Scan(&conced)
// national table ranking
var teamsnr,teamPosition float32
fmt.Println("How many points does the table leader have:")
 fmt.Scan(&teamsnr)
 fmt.Println("How many points does your team have:")
fmt.Scan(&teamPosition)


 fmt.Printf("%s:\n",teamOne)
 fmt.Printf("win percentage is :%v\n",result)
 goalavg(scored,conced,mostScore,mostConced)
tableranking(teamsnr,teamPosition)
//Team two 2

var teamTwoW,teamTwoD,teamTwoL float32 
	var teamTwo string
fmt.Println("enter team 2:")
//5 last games second team
fmt.Scan(&teamTwo)
fmt.Printf("out of 5 games how many wins:\n")
fmt.Scan(&teamTwoW)
fmt.Printf("out of 5 games how many Draws:\n")
fmt.Scan(&teamTwoD)
fmt.Printf("out of 5 games how many Losses:\n")
fmt.Scan(&teamTwoL)
fivelastG(teamTwoW,teamTwoD,teamTwoL)
 resulttwo:=fivelastG(teamTwoW,teamTwoD,teamTwoL)

//goalavg second team
 var mostScore2,mostConced2,scored2,conced2 float32
 fmt.Println("Most leuage goals scored:")
 fmt.Scan(&mostScore2)
 fmt.Println("Most leuage goals conceded:")
  fmt.Scan(&mostConced2)
  fmt.Println("Team goals scored:")
   fmt.Scan(&scored2)
   fmt.Println("Team goals conced:")
   fmt.Scan(&conced2)

//national table ranking of team 2
var teamsnr2,teamPosition2 float32
fmt.Println("How many points does the table leader have:")
 fmt.Scan(&teamsnr2)
 fmt.Println("How many points does your team have:")
fmt.Scan(&teamPosition2)



fmt.Printf("%s:\n",teamOne)
 fmt.Printf("win percentage is :%v\n",result)
 goalavg(scored,conced,mostScore,mostConced)
tableranking(teamsnr,teamPosition)




fmt.Printf("%s:\n",teamTwo)
fmt.Printf("win percentage is :%v\n",resulttwo)
goalavg(scored2,conced2,mostScore2,mostConced2)
tableranking(teamsnr2,teamPosition2)

//team1winning



}