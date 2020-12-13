#! /usr/bin/env gforth

\ Solution to Problem 1 from the Advent of Code
\ http://adventofcode.com/day/1
\
\ @author Matthew M. Burke <matthew@bluedino.net>
\ @date 2016-03-05
\
\ NOTE: I've re-formatted the input, rather than re-implented
\ how gforth parses source files.
\

: (  1 + ;
: )  1 - ;

0

s" 1.input.fth" included
." Floor " . cr bye

