% roll(1) | user manual 0.0.0 07bf70f
% R. S. Doiel  
% 2023-06-05

# NAME

roll

# SYSNOPSIS

roll DICE_NOTATIOJN [DICE_NOTATION ...]

# DESCRIPTION

roll parses the [dice notation](https://en.wikipedia.org/wiki/Dice_notation)and models the rolling of the dice described by the notation.

roll is part of the gameset collection of game programs.

# DICE NOTATION

Dice notation is formed by an integer (number of dice, A) the letter "d" followed by an integer (number of sides, X). The dice description can optionally be followed by a plus or minus and an integer (L) expressing the offset. The offset if present will be applied to the result of the dice rolled.  The roll program displays the dice notation used, the rolls and offsets if applied along with the total results. The dice notation forms are as follows.

AdX
: roll A number of dice with X sides, "3d4", "2d6", "1d8"

AdX+L
: roll A number of dice with X sides add L, "3d4+1", "2d6+2", "1d8+3"

AdX-L
: roll A number of dice with X side substract L, "3d4-1", "2d6-2", "1d8-3"


# OPTIONS

-help
: display help

-version
: display version

-license
: display license

-i FILENAME
: read from FILENAME

-o FILENAME
: Write to FILENAME

-t
: show only the roll total


# EXAMPLES

Rolling one twenty sided dice.

~~~
	roll 1d20
~~~

Rolling a three four sided dices plus two.

~~~
	roll 3d4+2
~~~

Rolling a four four sided dice minus two.

~~~
    roll 4d4-3
~~~


