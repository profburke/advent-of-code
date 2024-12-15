#! /usr/bin/env gforth

\ Solution to Problem 1b from the Advent of Code
\ http://adventofcode.com/day/1

variable pos
: pos++  1 pos +! ;
pos++

: (  1 +  pos++ ;
: underground?  dup 0< ;
: )  1 -  underground? if pos ? cr bye else pos++ then ;

0

s" 1.input.fth" included

." Never went underground"  cr bye

