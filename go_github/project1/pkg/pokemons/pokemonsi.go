package pokemons

func AboutMe() string {
	return "Ali Zhumatayev 21B030676 :)"
}

func Info() string {
	str := "This is the my favorite Pokemons list :D " + AboutMe()
	return str
}
