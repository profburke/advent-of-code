#!/usr/bin/env lua

file = io.open(arg[1])

function tablify(p)
   if type(p) == "table" then
      return p
   else
      local r = {}
      r[1] = p
      return r
   end
end

function deepcopy(orig)
    local orig_type = type(orig)
    local copy
    if orig_type == 'table' then
        copy = {}
        for orig_key, orig_value in next, orig, nil do
            copy[deepcopy(orig_key)] = deepcopy(orig_value)
        end
        setmetatable(copy, deepcopy(getmetatable(orig)))
    else -- number, string, boolean, etc
        copy = orig
    end
    return copy
end

function lcomp(p, q)
   -- deepcopy is needed because table.remove is destructive
   -- making this an invalid comp function for table.sort
   p = deepcopy(p)
   q = deepcopy(q)
   if #p == 0 and #q == 0 then
      return nil
   elseif #p == 0 and #q > 0 then
      return true
   elseif #p > 0 and #q == 0 then
      return false
   else
      local pf = p[1]
      local qf = q[1]

      local pt = type(pf)
      local qt = type(qf)
      
      table.remove(p, 1)
      table.remove(q, 1)

      if pt == "number" and qt == "number" then
         if pf < qf then
            return true
         elseif qf < pf then
            return false
         else
            return lcomp(p, q)
         end
         
      elseif pf == "table" and qf == "table" then
         local result = lcomp(pf, qf)
         if result ~= nil then
            return result
         else
            return lcomp(p, q)
         end
         
      else
         pf = tablify(pf)
         qf = tablify(qf)
         local result = lcomp(pf, qf)
         if result ~= nil then
            return result
         else
            return lcomp(p,q)
         end
      end
   end
end

-- Part 1
packets = {}

-- sum = 0
-- count = 0
-- index = 1
-- for line in file:lines() do
--    if line == "" then
--       local result = lcomp(packets[0], packets[1]) -- packets are always lists
--       local r
--       if result then
--          r = "true"
--       else
--          r = "false"
--       end
--       print("got result ... " .. r)
--       if result then
--          sum = sum + index
--       end
      
--       index = index + 1
--       count = 0
--    else
--       packets[count] = load('return ' .. line)()
--       count = (count + 1) % 2
--    end
-- end

-- print(sum)

for line in file:lines() do
   if line ~= "" then
      line = line:gsub('%[', '{')
      line = line:gsub('%]', '}')
      local p = load('return ' .. line)()
      table.insert(packets, p)
   end
end

table.insert(packets, {{2}})
table.insert(packets, {{6}})


table.sort(packets, lcomp)

decoderKey = 1

function isDivider(packet, val)
   if #packet ~= 1 then return false end
   if type(packet[1]) ~= "table" then return false end
   if #packet[1] ~= 1 then return false end
   if packet[1][1] ~= val then return false end
   return true
end

for k,v in pairs(packets) do
   if isDivider(v, 2) or isDivider(v, 6) then
      decoderKey = decoderKey * k
   end
end

print('decoder key ' .. decoderKey)
