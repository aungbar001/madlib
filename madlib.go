package main 

import "fmt"

func main() {
  var name1 string
  var name2 string 
  var adjec1 string 
  var adjec2 string 
  var adverb string
  var verbs1 string 
  var verbs2 string
  var verPast string
  var bodyPart string

  fmt.Println("Please Enter in 2 names")
  fmt.Scanln(&name1)
  fmt.Scanln(&name2)
  fmt.Println("You Entered:",name1, name2)

  fmt.Println("Please Enter in 2 Adjectives")
  fmt.Scanln(&adjec1)
  fmt.Scanln(&adjec2)
  fmt.Println("You Entered:", adjec1, adjec2)

  fmt.Println("Please Enter in 1 Adverb")
  fmt.Scanln(&adverb)
  fmt.Println("You Entered:", adverb)

  fmt.Println("Please Enter in 2 Verbs")
  fmt.Scanln(&verbs1)
  fmt.Scanln(&verbs2)
  fmt.Println("You Entered:", verbs1, verbs2)

  fmt.Println("Please Enter in a Verb Past Tence")
  fmt.Scanln(&verPast)
  fmt.Println("You Entered:", verPast)

  fmt.Println("Please Enter in a Body Part")
  fmt.Scanln(&bodyPart)
  fmt.Println("You Entered:", bodyPart)
  
  fmt.Println("Once apon a time there was a princess. Her name was", name1,",and she was quite", adjec1,", so much so she had peopel", adverb," for her. One day her father, King", name2, "and he was in a ", adjec2, "there was no one who wanted ot mairre his daughter and with no son's, he needed her to mary. so he went to her and",verbs1, ",her to pick a sutor. she",adverb, "at his pequest and pinted at his",bodyPart, "and said 'are you going to get ride of that', the king in a",verbs2, "sent ehr to a far away place never to be seen agine...FIN")



}