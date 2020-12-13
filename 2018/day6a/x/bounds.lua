
points = require 'points'

maxX = -1000000
minX = 1000000

maxY = -1000000
minY = 1000000

for _, p in pairs(points) do
   if p[1] < minX then minX = p[1] end
   if p[1] > maxX then maxX = p[1] end
   if p[2] < minY then minY = p[2] end
   if p[2] > maxY then maxY = p[2] end
end

print(minX, maxX)
print(minY, maxY)

