#!/usr/bin/env lua

--[[
Solution to Problem 1 from the Advent of Code
http://adventofcode.com/day/1
--]]


if #arg < 1 then
   print(string.format('Usage: %s <infile>', arg[0]))
   os.exit(1)
end


local inp = io.open(arg[1], 'r')
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
end

print(string.format('Floor %d', floor))


