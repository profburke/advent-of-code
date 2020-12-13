points = require 'input'

local w, h  
local step = 10828

local max = 60
local timer = max

function love.load()
   w = love.graphics.getWidth()
   h = love.graphics.getHeight()
   for _, p in ipairs(points) do
      p.p[1] = p.p[1] + p.v[1] * 10828
      p.p[2] = p.p[2] + p.v[2] * 10828
   end
end

function love.update(dt)
   if timer == 0 then
      for _, p in ipairs(points) do
         p.p[1] = p.p[1] + p.v[1]
         p.p[2] = p.p[2] + p.v[2]
      end
      step = step + 1
      timer = max
   else
      timer = timer - 1
   end
end

function love.draw()
   for _, p in ipairs(points) do
      local x = p.p[1] + w/2
      local y = p.p[2] + h/2
      if x > 0 and y > 0 and x < w and y < h then
         love.graphics.points(x, y)
      end
   end
   love.graphics.print('step ' .. step, 20, 20)
end
