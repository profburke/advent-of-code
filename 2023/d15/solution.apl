
h←{256|17×⍺+⍵}
hash←{⍺←0⋄0=≢⍵:⍺⋄(⍺h⎕ucs⊃⍵)∇1↓⍵} ⍝ This just computes the hash
split←{⍵⊆⍨⍺≠⍵}
splitComma←{⍺←','⋄⍺split ⍵}


sample←'rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7'


⍝ need to split input on , then apply rh to each and sum
