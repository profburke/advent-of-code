local points = require 'points'
local grid = require 'generated'
local colorFor = require 'color'

function love.draw()
   for x = 0, 359 do
      for y = 0, 359 do
         local color = colorFor(grid[x][y])
         love.graphics.setColor(color.r, color.g, color.b)
         love.graphics.rectangle('fill', x * 2, y * 2, 2, 2)
      end
   end

   love.graphics.setColor(1, 1, 1, 1)
   for _, p in ipairs(points) do
      love.graphics.rectangle('fill', p[1] * 2, p[2] * 2, 2, 2)
   end

end
