DLX-Sudoku
==========
Solves puzzles using [Knuth's DLX](http://en.wikipedia.org/wiki/Dancing_Links).  DLX solves all Sudokus really fast, with modest memory requirements and without too much code.

This code follows his original paper fairly closely. Input to the program is a file containing puzzles in 81 character format. (This seems to be a conventional computer representation for Sudoku puzzles.)

The DLX code is separated from the Sudoku specific code and so should be usable for other constraint problems, but no other problems are shown here.  Also the code sets up the Sudoku specific constraints is in a separate function, but is pretty much hard coded for Sudoku and not generalized to allow constraint input in any kind of friendly format.

