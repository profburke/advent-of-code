local color = function(p)
   if p == -1 then
      return { r = 0, g = 0, b = 0 }
   elseif p % 3 == 0 then
      return { r = ((30 + p) * 5)/255, g = 0, b = 0 }
   elseif p % 3 == 1 then
      return { g = ((30 + p) * 5)/255, r = 0, b = 0 }
   else
      return { b = ((30 + p) * 5)/255, r = 0, g = 0 }
   end
end


return color
