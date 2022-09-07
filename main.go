package ColdWeatherActivitiesNow

import "fmt"

package ageGroup

func child(playwith,friend,snowEquipment){
	games = ["Build snowman","Make snow angels"," Engage in a snowball fight","Make a snow castle"]
	activities = ["Watch TV","Play board games","Play videogames","Make a hot cup of coffee or tea"]
	
	if playingwith == "oneself":
		if friend == 0:
			if snowEquipment == true:
			fmt.PrintIn("Do either of these" + game[0] + "or if staying inside" + activity[0] + "or" activity[1] + ".")
	else: fmt.PrintIn()
	if playingwith == "friend1":
		if snowEquipment == false:
			fmt.PrintIn("Do either of these activities" + activity[1] + "or try" + activity[2] + ".")
	else: fmt.PrintIn("If none of these activties interest you or you cannot do these activties, I would reccomend pursuing your current interests.")
}

func adult(playingwith){
	adultActivities = ["Play a board game or video game","Go out for drinks","Cuddle together","Make popcorn and watch a movie","Work if your job requires a laptop"]
	if playingwith == "oneself":
		fmt.PrintIn("If you are by yourself you can do the following activties: "+ adultActivities[0] + "," + adultActivities[3] + ", and" + adultActivities[4] + ".")
	if playingwith == "friend":
		fmt.PrintIn("If you are with a friend, you can do the following activities: " + adultActivities[0] + "," + adultActivities[1] + ", and" + adultActivities[3] + ".")
	if playingwith == "significant other":
		fmt.PrintIn("If you are with your significant other, the both of you can do the following: " + adultActivities[1] + "," + adultActivities[2] + ", and" + adultActivities[3] + ".")
	else: fmt.PrintIn("If none of these activties interest you or you cannot do these activties, I would reccomend pursuing your current interests.")

}