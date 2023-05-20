% cards(1) user manual 0.0.0 971caf9
% R. S. Doiel
% 2023-05-20

# NAME

cards

# SYSNOPSIS

cards ACTION DECK [MODIFIERS]

# DESCRIPTION

cards uses the command line an a data file presenting the
deck to play card games. The form of the command uses a verb
and modifier structure similar to programs a like the Go compiler
and Git source versioning system.  cards doesn't apply or
maintain game ruleset it only supports managing the state of the
data file holding the representation of the deck.

cards is part of the gameset collection of game programs.

# ACTIONS

## Deck operations

new
: creates a new card deck data file

reset
: sets the card deck back to the "new" state

game
: sets the name of the card game and list of players

shuffle
: shuffle the whole deck of playing cards setting

deal
: deal the cards to the players. Players list must be set previously.

deal-faceup
: deals cards to the players in a visible state

show
: shows the contents of the deck

### player's hand operations

show-hand
: shows the all contents of a player's hand (both held and visible)

reveal-card
: moves a card in player's hand from held to visible

hide-card
: moves a card in a player's hand from visible to held

# OPTIONS

-help
: display help

-license
: display license

-version
: display version

-i FILENAME
: read from filename

-o FILENAME
: write to filename


# EXAMPLE

Creating a new deck called "my_deck.json".

~~~
cards new my_deck.json
~~~

Reset deck, resets the deck to a new state.

~~~
cards reset my_deck.json
~~~

Setting the name for the game using "guess-cards.json".
Sets the name and adds two players.

~~~
	cards game my_deck.json "guess a my card" "jane" "joe"
~~~

Shuffle cards

~~~
	cards suffle my_deck.json
~~~

Deal cards

~~~
cards deal my_deck.json
~~~


