Try calculate `1+2+...+10(=55)` by minicpu!

```shell
$ go run main.go
 i       OP      PC      IR      Reg0    Reg1    Reg2    Reg3
 0      LDH         1    4800       0       0   0       0
 1      LDL         2    4000       0       0   0       0
 2      LDH         3    4900       0       0   0       0
 3      LDL         4    4101       0       1   0       0
 4      LDH         5    4a00       0       1   0       0
 5      LDL         6    4200       0       1   0       0
 6      LDH         7    4b00       0       1   0       0
 7      LDL         8    430a       0       1   0       10
 8      ADD         9     a20       0       1   1       10
 9      ADD        10     840       1       1   1       10
10      ST         11    7040       1       1   1       10
11      CMP        12    5260       1       1   1       10
12      JE         13    580e       1       1   1       10
13      JMP         8    6008       1       1   1       10
14      ADD         9     a20       1       1   2       10
15      ADD        10     840       3       1   2       10
16      ST         11    7040       3       1   2       10
17      CMP        12    5260       3       1   2       10
18      JE         13    580e       3       1   2       10
19      JMP         8    6008       3       1   2       10
20      ADD         9     a20       3       1   3       10
21      ADD        10     840       6       1   3       10
22      ST         11    7040       6       1   3       10
23      CMP        12    5260       6       1   3       10
24      JE         13    580e       6       1   3       10
25      JMP         8    6008       6       1   3       10
26      ADD         9     a20       6       1   4       10
27      ADD        10     840      10       1   4       10
28      ST         11    7040      10       1   4       10
29      CMP        12    5260      10       1   4       10
30      JE         13    580e      10       1   4       10
31      JMP         8    6008      10       1   4       10
32      ADD         9     a20      10       1   5       10
33      ADD        10     840      15       1   5       10
34      ST         11    7040      15       1   5       10
35      CMP        12    5260      15       1   5       10
36      JE         13    580e      15       1   5       10
37      JMP         8    6008      15       1   5       10
38      ADD         9     a20      15       1   6       10
39      ADD        10     840      21       1   6       10
40      ST         11    7040      21       1   6       10
41      CMP        12    5260      21       1   6       10
42      JE         13    580e      21       1   6       10
43      JMP         8    6008      21       1   6       10
44      ADD         9     a20      21       1   7       10
45      ADD        10     840      28       1   7       10
46      ST         11    7040      28       1   7       10
47      CMP        12    5260      28       1   7       10
48      JE         13    580e      28       1   7       10
49      JMP         8    6008      28       1   7       10
50      ADD         9     a20      28       1   8       10
51      ADD        10     840      36       1   8       10
52      ST         11    7040      36       1   8       10
53      CMP        12    5260      36       1   8       10
54      JE         13    580e      36       1   8       10
55      JMP         8    6008      36       1   8       10
56      ADD         9     a20      36       1   9       10
57      ADD        10     840      45       1   9       10
58      ST         11    7040      45       1   9       10
59      CMP        12    5260      45       1   9       10
60      JE         13    580e      45       1   9       10
61      JMP         8    6008      45       1   9       10
62      ADD         9     a20      45       1   10      10
63      ADD        10     840      55       1   10      10
64      ST         11    7040      55       1   10      10
65      CMP        12    5260      55       1   10      10
66      JE         14    580e      55       1   10      10
67      HLT        15    7800      55       1   10      10
ram[64] = 55
```
