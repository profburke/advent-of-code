points = require 'points'
print 'local grid = {}'

function distance(x0, y0, x1, y1)
   return math.abs(x0 - x1) + math.abs(y0 - y1)
end

function closestTo(x, y)
   local ties = {}
   local min = math.huge
   local index = -1

   for i, p in ipairs(points) do
      local d = distance(x, y, p[1], p[2])
      if d < min then
         min = d
         index = i
      end

      ties[d] = (ties[d] or 0) + 1
   end

   if ties[min] > 1 then
      return -1
   else
      return index
   end
end

for x = 0, 359 do
   print('grid[' .. x .. '] = {}')
end
for x = 0, 359 do
   for y =  0, 359 do
      local c = closestTo(x, y)
      print('grid[' .. x .. '][' .. y ..'] = ' .. c)
   end
end

print 'return grid'
