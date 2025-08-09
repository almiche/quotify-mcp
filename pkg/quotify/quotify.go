package quotify

import (
	"math/rand"
	"time"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

type Quotify struct {
	Authors []string
	Quotes  []string
	Spacer  string
}

func New() *Quotify {
	return &Quotify{
		Authors: []string{
			"Ivanka Trump",
			"Dog The Bounty Hunter",
			"Master Yoda",
			"Abe Lincoln",
			"Stev-o",
			"Wee-man",
			"Albus Dumbledore",
			"Satan",
			"Momo Taleb",
			"Undertaker",
			"John Cena",
			"Triple H",
			"Kane",
			"Big Show",
			"The Rock",
			"Bing Han",
			"Sarah Palin",
			"Dj Khaled",
			"21 Savage",
			"Sugar Ray Robinson",
			"Soulja Boy",
			"Betsy DeVos",
			"The Red Power Ranger",
			"Oprah Winfrey",
			"Snoop Dogg",
			"Olga",
			"Logan Paul",
			"Judge Tyco",
			"Fred",
		},
		Quotes: []string{
			"The ode lives upon the ideal, the epic upon the grandiose, the drama upon the real.",
			"If we don't study the mistakes of the future we're doomed to repeat them for the first time.",
			"C++ supports OOP",
			"You can do anything, but not everything.",
			"Perfection is achieved, not when there is nothing more to add, but when there is nothing left to take away.",
			"The richest man is not he who has the most, but he who needs the least.",
			"You miss 100 percent of the shots you never take.",
			"Courage is not the absence of fear, but rather the judgement that something else is more important than fear.",
			"You must be the change you wish to see in the world.",
			"When hungry, eat your rice; when tired, close your eyes. Fools may laugh at me, but wise men will know what I mean.",
			"To the man who only has a hammer, everything he encounters begins to look like a nail.",
			"We are what we repeatedly do; excellence, then, is not an act but a habit.",
			"The weak can never forgive. Forgiveness is the attribute of the strong.",
			"Happiness is when what you think, what you say, and what you do are in harmony.",
			"An eye for eye only ends up making the whole world blind.",
			"Live as if you were to die tomorrow; learn as if you were to live forever.",
			"First they ignore you, then they laugh at you, then they fight you, then you win.",
			"You must not lose faith in humanity. Humanity is an ocean; if a few drops of the ocean are dirty, the ocean does not become dirty.",
			"The best way to find yourself is to lose yourself in the service of others.",
			"Strength does not come from physical capacity. It comes from an indomitable will.",
			"A man is but the product of his thoughts; what he thinks, he becomes.",
			"YES we can",
			"No we can't",
			"Those were alternative facts",
			"You can't see me",
			"I have a dream",
			"In god we trust",
			"Thou shall not pass",
			"What is a private email server?",
			"It's local on the the remote server",
			"Don't ever play yourself",
			"Just play. Have fun. Enjoy the game.",
			"Being independent, being confident and having fun is what matters.",
			"It's kind of fun to do the impossible.",
			"We gucci Fam - Ghandi",
			"I want my world to be fun.",
			"Everyone gets a car",
			"One more thing",
			"Wrong",
			"Bazinga!",
			"You shall not pass!",
			"Wingardium leviosa",
			"Go bing or go home !",
			"Chronological awareness.",
			"This is a problem right here.",
			"Oh my goodness..",
			"www.loganpaul.com/shop",
			"Get your merch on",
			"Best merch in the game",
			"Link in Bio",
			"Be a maverick",
			"DANG DAWG.",
			"On the mandem level",
			"you're born and then you die that's all there is to it",
			"i think dreams are a socialist construct",
			"The force is strong with this one",
			"Winter is coming",
			"May the odds be ever in your favor",
			"With great power comes great responsibility",
			"Life is like a box of chocolates",
			"Elementary my dear Watson",
			"Houston we have a problem",
			"I'll be back",
			"May the force be with you",
			"Keep your friends close but your enemies closer",
		},
		Spacer: " - ",
	}
}

func (q *Quotify) Generate() Quote {
	rand.Seed(time.Now().UnixNano())
	
	randomQuote := q.Quotes[rand.Intn(len(q.Quotes))]
	randomAuthor := q.Authors[rand.Intn(len(q.Authors))]
	
	return Quote{
		Text:   randomQuote,
		Author: randomAuthor,
	}
}

func (q *Quotify) GenerateString() string {
	quote := q.Generate()
	return quote.Text + q.Spacer + quote.Author
}