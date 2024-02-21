package main

import (
	"github.com/naripsaejan/cinema/movie"
	"github.com/naripsaejan/cinema/ticket"
)

func main() {
	movieName := movie.FindMovieName("tt1825683")
	movie.ReviewMovie(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
