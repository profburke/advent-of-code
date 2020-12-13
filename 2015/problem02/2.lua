#!/usr/bin/env lua

--[[
Solution to Problem 2 from the Advent of Code
http://adventofcode.com/day/2
--]]


if #arg < 1 then
   print(string.format('Usage: %s <infile>', arg[0]))
   os.exit(1)
end

local min = math.min

local function sa(w, l, h)
   return (2 * (l*w + w*h + l*h))
end

local inp = io.open(arg[1], 'r')

local sum = 0
for g in inp:lines() do
   local _, _, w, l, h = g:find '(%d+)x(%d+)x(%d+)'
   sum = sum + sa(w, l, h) + min(l*w, l*h, w*h)
end

inp:close()
print(sum)
