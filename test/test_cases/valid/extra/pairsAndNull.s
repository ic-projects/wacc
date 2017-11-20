-- Compiling...
-- Printing Assembly...
pairsAndNull.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 50
4		.ascii	"NullReferenceError: dereference a null reference\n\0"
5	msg_1:
6		.word 5
7		.ascii	"%.*s\0"
8	
9	.text
10	
11	.global main
12	f_f:
13		PUSH {lr}
14		LDR r4, [sp, #4]
15		MOV r0, r4
16		POP {pc}
17		POP {pc}
18		.ltorg
19	f_f2:
20		PUSH {lr}
21		LDR r4, [sp, #4]
22		MOV r0, r4
23		POP {pc}
24		POP {pc}
25		.ltorg
26	main:
27		PUSH {lr}
28		SUB sp, sp, #48
29		LDR r4, =0
30		STR r4, [sp, #44]
31		LDR r4, [sp, #44]
32		MOV r0, r4
33		BL p_check_null_pointer
34		LDR r4, [r4]
35		LDR r4, [r4]
36		STR r4, [sp, #40]
37		LDR r4, [sp, #40]
38		STR r4, [sp, #-4]!
39		BL f_f
40		ADD sp, sp, #4
41		MOV r4, r0
42		STR r4, [sp, #36]
43		LDR r0, =8
44		BL malloc
45		MOV r4, r0
46		LDR r5, [sp, #44]
47		LDR r0, =4
48		BL malloc
49		STR r5, [r0]
50		STR r0, [r4]
51		LDR r5, [sp, #40]
52		LDR r0, =4
53		BL malloc
54		STR r5, [r0]
55		STR r0, [r4, #4]
56		STR r4, [sp, #32]
57		LDR r0, =8
58		BL malloc
59		MOV r4, r0
60		LDR r5, [sp, #32]
61		LDR r0, =4
62		BL malloc
63		STR r5, [r0]
64		STR r0, [r4]
65		LDR r5, =0
66		LDR r0, =4
67		BL malloc
68		STR r5, [r0]
69		STR r0, [r4, #4]
70		STR r4, [sp, #28]
71		LDR r0, =8
72		BL malloc
73		MOV r4, r0
74		LDR r5, =0
75		LDR r0, =4
76		BL malloc
77		STR r5, [r0]
78		STR r0, [r4]
79		LDR r5, =0
80		LDR r0, =4
81		BL malloc
82		STR r5, [r0]
83		STR r0, [r4, #4]
84		STR r4, [sp, #24]
85		LDR r4, =0
86		STR r4, [sp, #-4]!
87		BL f_f
88		ADD sp, sp, #4
89		MOV r4, r0
90		STR r4, [sp, #20]
91		LDR r0, =8
92		BL malloc
93		MOV r4, r0
94		LDR r5, =1
95		LDR r0, =4
96		BL malloc
97		STR r5, [r0]
98		STR r0, [r4]
99		LDR r5, =0
100		LDR r0, =4
101		BL malloc
102		STR r5, [r0]
103		STR r0, [r4, #4]
104		STR r4, [sp, #16]
105		LDR r0, =8
106		BL malloc
107		MOV r4, r0
108		LDR r5, =1
109		LDR r0, =4
110		BL malloc
111		STR r5, [r0]
112		STR r0, [r4]
113		LDR r5, [sp, #16]
114		LDR r0, =4
115		BL malloc
116		STR r5, [r0]
117		STR r0, [r4, #4]
118		STR r4, [sp, #12]
119		LDR r4, [sp, #12]
120		MOV r0, r4
121		BL p_check_null_pointer
122		LDR r4, [r4, #4]
123		LDR r4, [r4]
124		STR r4, [sp, #8]
125		LDR r4, [sp, #16]
126		STR r4, [sp, #-4]!
127		BL f_f2
128		ADD sp, sp, #4
129		MOV r4, r0
130		STR r4, [sp, #4]
131		LDR r4, [sp, #44]
132		MOV r0, r4
133		BL p_check_null_pointer
134		LDR r4, [r4]
135		LDR r4, [r4]
136		STR r4, [sp]
137		ADD sp, sp, #48
138		LDR r0, =0
139		POP {pc}
140		.ltorg
141	p_check_null_pointer:
142		PUSH {lr}
143		CMP r0, #0
144		LDREQ r0, =msg_0
145		BLEQ p_throw_runtime_error
146		POP {pc}
147	p_throw_runtime_error:
148		BL p_print_string
149		MOV r0, #-1
150		BL exit
151	p_print_string:
152		PUSH {lr}
153		LDR r1, [r0]
154		ADD r2, r0, #4
155		LDR r0, =msg_1
156		ADD r0, r0, #4
157		BL printf
158		MOV r0, #0
159		BL fflush
160		POP {pc}
161	
===========================================================
-- Finished
