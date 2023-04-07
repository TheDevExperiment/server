package funnynamer

import (
	"math/rand"
	"time"
)

// List of funny adjectives
var adjectives = []string{
	"Awesome", "Silly", "Sassy", "Crazy", "Wacky", "Zany", "Epic",
	"Amazing", "Funky", "Quirky", "Goofy", "Fantastic", "Radical",
	"Legendary", "Marvelous", "Bodacious", "Terrific", "Groovy",
	"Outrageous", "Phenomenal", "Hilarious", "Superb", "Bizarre",
	"Mega", "Ultra", "Hyper", "Gigantic", "Colossal", "Monstrous",
}

// List of funny nouns
var nouns = []string{
	"Ninja", "Pirate", "Robot", "Banana", "Squid", "Chicken", "Monkey",
	"Kangaroo", "Llama", "Penguin", "Hippo", "Giraffe", "Cactus",
	"Bumblebee", "Buffalo", "Panda", "Dinosaur", "Lobster", "Dragon",
	"Yeti", "Unicorn", "Zombie", "Alien", "Goblin", "Wizard",
}

// GetFunnyName returns a random funny name
func GetFunnyName() string {
	rand.Seed(time.Now().UnixNano())
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	return adjective + "_" + noun
}
