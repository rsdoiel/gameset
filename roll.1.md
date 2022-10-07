% roll(1) user manual
% R. S. Doiel
% 2022-10-01

# NAME

roll

# SYSNOPSIS

roll DICE_ROLL_EXPR [DICE_ROLL_EXPR ...]

# DESCRIPTION

roll parses the dice roll expressions and
models the rolling of dice described by the expressions.
Dice roll expressions are formed by an integer (number of dice)
the letter "d" followed by an integer (number of sides) and optionally
follow by a plus or minus and an integer which is treated as an offset.

roll is part of the gameset collection of game programs.

# EXAMPLE

Rolling one twenty sided dice.

```
	roll 1d20
```

Rolling a three four sided dices plus two.

```
	roll 3d4+2
```


