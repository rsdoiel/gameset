package dice

var (
	HelpText = `% {app_name}(1) user manual
% R. S. Doiel
% 2022-10-07

# NAME

{app_name}

# SYSNOPSIS

{app_name} DICE_NOTATIOJN [DICE_NOTATION ...]

# DESCRIPTION

{app_name} parses the [dice notation](https://en.wikipedia.org/wiki/Dice_notation)and models the rolling of the dice described by the notation.

{app_name} is part of the gameset collection of game programs.

## DICE NOTATION

Dice notation is formed by an integer (number of dice, A) the letter "d" followed by an integer (number of sides, X). The dice description can optionally be followed by a plus or minus and an integer (L) expressing the offset. The offset if present will be applied to the result of the dice {app_name}ed.  The {app_name} program displays the dice notation used, the {app_name}s and offsets if applied along with the total results. The dice notation forms are as follows.

AdX
: {app_name} A number of dice with X sides, "3d4", "2d6", "1d8"

AdX+L
: {app_name} A number of dice with X sides add L, "3d4+1", "2d6+2", "1d8+3"

AdX-L
: {app_name} A number of dice with X side substract L, "3d4-1", "2d6-2", "1d8-3"


# OPTIONS

-t
: show only the {app_name} total

-help
: display help

-version
: display gameset version

# EXAMPLES

Rolling one twenty sided dice.

` + "```" + `
	{app_name} 1d20
` + "```" + `

Rolling a three four sided dices plus two.

` + "```" + `
	{app_name} 3d4+2
` + "```" + `

Rolling a four four sided dice minus two.

` + "```" + `
    {app_name} 4d4-3
` + "```" + `

`
)
