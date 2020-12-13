local points = require 'input'

max = 0

for _, p in ipairs(points) do
   local t = math.abs(p.p[1]) / math.abs(p.v[1])
   if t > max then
      max = t
   end
   t = math.abs(p.p[2]) / math.abs(p.v[2])
   if t > max then
      max = t
   end
end

print(max)

