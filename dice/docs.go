package dice

var (
	HelpText = `% {app_name}(1) user manual
% R. S. Doiel
% 2022-10-01

# NAME

{app_name}

# SYSNOPSIS

{app_name} DICE_ROLL_EXPR [DICE_ROLL_EXPR ...]

# DESCRIPTION

{app_name} parses the dice roll expressions and
models the rolling of the dice described by the expressions.
Dice roll expressions are forms by an integer (number of dice)
the letter "d" followed by an integer (number of sides) and optionally
follow by a plus or minus and integer which is treated as an offset.

{app_name} is part of the gameset collection of game programs.

# EXAMPLE

Rolling one twenty sided dice.

` + "```" + `
	{app_name} 1d20
` + "```" + `

Rolling a three four sided dices plus two.

` + "```" + `
	{app_name} 3d4+2
` + "```" + `

`
)
