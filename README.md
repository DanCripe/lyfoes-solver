# Lyfoes Solver

This repository contains a puzzle solver for the Lyfoes game (by IBNPlay, with which I have
no connection).

The implementation is in *GOLANG* but makes no use of any of that languages characteristic
capabilities (goroutines, channels). In addition, it is a brute-force method, simply trying
random (legal) moves until the problem is solved or a stalemate condition is detected.

# Motivation

My initial motivation for writing this program was to get past a difficult puzzle (4-1144)
that had no published solution (that I could find). Unwilling to give up, I wrote the
original version of this program in an afternoon, although the initial working
implementation took about 3 minutes to find a solution (I've improved that to at most a
handful of seconds - some puzzles are solved in sub-second time).

# Improvements

I've made some effort to eliminate stupid moves (such as moves that simply undo previous
moves or that move from a stack of three onto a stack of one, when the reverse would have
the same result, but two fewer moves). There are still cases where the program generates
less-than-ideal solutions, but I'm currently satisfied with the results as being useful
for getting past any particularly tough puzzles.

The available color pallet is likely not enough to cover the most difficult level of
puzzles. New colors can be simply added to the ```colors.go``` file and the maps that
are used to translate to and from text extended.

Note that I used integer-based color definitions in the solver, rather than string-based
enum values, simply out of concern that comparing and processing strings would require
more CPU time.

# Input

The input file is expected to contain a description of the puzzle board using two-character
color definitions, with each column separated by a space:

 - DG (Dark Green)
 - MG (Medium Green)
 - LG (Light Green)
 - DB (Dark Blue)
 - MB (Medium Blue)
 - LB (Light Blue)
 - DR (Dark Red)
 - MR (Medium Red aka bright pink)
 - LR (Light Red aka light pink)
 - Or (Orange)
 - Wh (White)
 - Gr (Grey)
 - Ye (Yellow)
 - Pu (Purple)

An example board:
```
Wh Ye LG Or LR MR LG DR DB LR Gr LB MG
DB Wh Pu LR Ye Gr MB DB MR DR MB Pu MB
MG MR LG Gr LB Ye Or Wh Or LB DB MR MG
LB Pu DR LG MB Ye DR Wh Gr Pu MG LR Or
```

# Output

The output of the program includes a representation of the solved board and
all of the moves (source column, destination column, and color) that resulted
in the solution, and the number of iterations it took to solve it. The program
also prints the time to each power-of-ten number of iteration attempts as they
are encountered, just to give an idea of how fast the program is running through
iterations.

# Future Enhancements

The board input could use some attention in the areas of validating the input. There
is currently no handling for color representations that are invalid or mistyped colors
that result in an uneven number of each color.

There are a few cases of "stupid moves" that remain undetected that would make the results
shorter (and less frustrating to enter in the game).

Finally, being able to represent the moves in color and even have an animated solution
would be fantastic. Unfortunately, I'm not skilled in animation of any sort, so it's
unlikely that I'll ever attempt it.
