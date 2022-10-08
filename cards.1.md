% cards(1) user manual
% R. S. Doiel
% 2022-10-07

# NAME

cards

# SYSNOPSIS

cards ACTION DECK_NAME [MODIFIERS]

# DESCRIPTION

cards uses the command line an a data file presenting the
deck to play card games. The form of the command uses a verb
and modifier structure similar to programs a like the Go compiler
and Git source versioning system.  cards doesn't apply or
maintain game ruleset it only supports managing the state of the
data file holding the representation of the deck.

cards is part of the gameset collection of game programs.

## Supported actions

### Deck operations

new [OPTION=]DECK\_NAME
: creates a new card deck data file. You specify a specialized deck
by prefixing the DECK\_NAME with the deck type, e.g. standard, rook,
pinochle.

reset DECK\_NAME
: sets the card deck back to the "new" state

game DECK\_NAME GAME\_NAME PLAYER\_NAME [PLAYER\_NAME ...]
: sets the name of the card game and list of players

shuffle DECK\_NAME
: shuffle the whole deck of playing cards setting

deal DECK\_NAME
: deal the cards to the players. Players list must be set previously.

deal-faceup DECK\_NAME COUNT
: deals cards to the players in a visible state

show DECK\_NAME
: shows the contents of the deck

### player's hand operations

show-hand DECK\_NAME PLAYER\_NAME
: shows the all contents of a player's hand (both held and visible)

reveal-card DECK\_NAME PLAYER\_NAME CARD
: moves a card in player's hand from held to visible

hide-card DECK\_NAME PLAYER\_NAME CARD
: moves a card in a player's hand from visible to held


# EXAMPLE

Creating a new deck called "my_deck.json".

```
	cards new my_deck.json
```

Reset deck, resets the deck to a new state.

```
	cards reset my_deck.json
```

Setting the name for the game using "guess-cards.json".
Sets the name and adds two players.

```
	cards game my_deck.json "guess a my card" "jane" "joe"
```

Shuffle cards

```
	cards suffle my_deck.json
```

Deal cards

```
	cards deal my_deck.json
```



