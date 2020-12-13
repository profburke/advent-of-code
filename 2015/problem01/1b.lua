#!/usr/bin/env lua

--[[
Solution to Problem 1b from the Advent of Code
http://adventofcode.com/day/1
--]]


local inp = io.open('1.input', 'r')
local instructions = inp:read '*a'
inp:close()

local floor = 0

for p = 1,instructions:len() do
   local c = instructions:sub(p, p)
   if c == '(' then
      floor = floor + 1
   elseif c == ')' then
      floor = floor - 1
   end

   if floor < 0 then
      print(string.format('Position %d', p))
      break
   end
end


