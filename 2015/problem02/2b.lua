#!/usr/bin/env lua

--[[
Solution to Problem 2b from the Advent of Code
http://adventofcode.com/day/2
--]]


if #arg < 1 then
   print(string.format('Usage: %s <infile>', arg[0]))
   os.exit(1)
end

local min = math.min

local function v(w, l, h)
   return l*w*h
end

local function perim(u, v)
   return 2*(u+v)
end

local inp = io.open(arg[1], 'r')

local sum = 0
for g in inp:lines() do
   local _, _, w, l, h = g:find '(%d+)x(%d+)x(%d+)'
   sum = sum + v(w, l, h) + min(perim(l, w), perim(l, h), perim(w, h))
end

inp:close()
print(sum)
