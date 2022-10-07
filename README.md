
[![Project Status: Concept – Minimal or no implementation has been done yet, or the repository is only intended to be a limited example, demo, or proof-of-concept.](https://www.repostatus.org/badges/latest/concept.svg)](https://www.repostatus.org/#concept)

Gameset
=======

A set of go packages and command line programs for modeling simple games.

The current gameset models dice and cards.

dice
----

The dice packages models multisided dice. It is demonstrated by the program `roll`. The `roll` program accepts parameters describing the dice to be modeled in [Dice notation](https://en.wikipedia.org/wiki/Dice_notation).

deck
----

The deck package models [playing_cards](https://en.wikipedia.org/wiki/Playing_card). It is demonstrated by the program `cards`.  The `cards` program supports multiple types of playing cards and the back actions you might perform with a "deck" such as shuffle and deal. The `cards` program stores the state of the deck (including player's hands) in a file using JSON encoding. The `cards` program does not support rules or game flow. But instead it can be scripted such that you can simulate a card game like solitare or go fish.


