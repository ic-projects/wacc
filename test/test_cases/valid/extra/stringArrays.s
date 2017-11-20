-- Compiling...
-- Printing Assembly...
stringArrays.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 3
4		.ascii	"fst"
5	msg_1:
6		.word 3
7		.ascii	"snd"
8	msg_2:
9		.word 44
10		.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
11	msg_3:
12		.word 45
13		.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
14	msg_4:
15		.word 5
16		.ascii	"true\0"
17	msg_5:
18		.word 6
19		.ascii	"false\0"
20	msg_6:
21		.word 1
22		.ascii	"\0"
23	msg_7:
24		.word 5
25		.ascii	"%.*s\0"
26	
27	.text
28	
29	.global main
30	main:
31		PUSH {lr}
32		SUB sp, sp, #12
33		LDR r4, =msg_0
34		STR r4, [sp, #8]
35		LDR r4, =msg_1
36		STR r4, [sp, #4]
37		LDR r0, =12
38		BL malloc
39		MOV r4, r0
40		LDR r5, [sp, #8]
41		STR r5, [r4, #4]
42		LDR r5, [sp, #4]
43		STR r5, [r4, #8]
44		LDR r5, =2
45		STR r5, [r4]
46		STR r4, [sp]
47		ADD r4, sp, #0
48		LDR r5, =0
49		LDR r4, [r4]
50		MOV r0, r5
51		MOV r1, r4
52		BL p_check_array_bounds
53		ADD r4, r4, #4
54		ADD r4, r4, r5, LSL #2
55		LDR r4, [r4]
56		LDR r5, [sp, #8]
57		CMP r4, r5
58		MOVEQ r4, #1
59		MOVNE r4, #0
60		MOV r0, r4
61		BL p_print_bool
62		BL p_print_ln
63		ADD r4, sp, #0
64		LDR r5, =0
65		LDR r4, [r4]
66		MOV r0, r5
67		MOV r1, r4
68		BL p_check_array_bounds
69		ADD r4, r4, #4
70		ADD r4, r4, r5, LSL #2
71		LDR r5, =1
72		LDR r4, [r4]
73		MOV r0, r5
74		MOV r1, r4
75		BL p_check_array_bounds
76		ADD r4, r4, #4
77		ADD r4, r4, r5
78		LDRSB r4, [r4]
79		MOV r5, #'f'
80		CMP r4, r5
81		MOVEQ r4, #1
82		MOVNE r4, #0
83		MOV r0, r4
84		BL p_print_bool
85		BL p_print_ln
86		ADD r4, sp, #0
87		LDR r5, =1
88		LDR r4, [r4]
89		MOV r0, r5
90		MOV r1, r4
91		BL p_check_array_bounds
92		ADD r4, r4, #4
93		ADD r4, r4, r5, LSL #2
94		LDR r5, =2
95		LDR r4, [r4]
96		MOV r0, r5
97		MOV r1, r4
98		BL p_check_array_bounds
99		ADD r4, r4, #4
100		ADD r4, r4, r5
101		LDRSB r4, [r4]
102		MOV r5, #'d'
103		CMP r4, r5
104		MOVEQ r4, #1
105		MOVNE r4, #0
106		MOV r0, r4
107		BL p_print_bool
108		BL p_print_ln
109		ADD sp, sp, #12
110		LDR r0, =0
111		POP {pc}
112		.ltorg
113	p_check_array_bounds:
114		PUSH {lr}
115		CMP r0, #0
116		LDRLT r0, =msg_2
117		BLLT p_throw_runtime_error
118		LDR r1, [r1]
119		CMP r0, r1
120		LDRCS r0, =msg_3
121		BLCS p_throw_runtime_error
122		POP {pc}
123	p_print_bool:
124		PUSH {lr}
125		CMP r0, #0
126		LDRNE r0, =msg_4
127		LDREQ r0, =msg_5
128		ADD r0, r0, #4
129		BL printf
130		MOV r0, #0
131		BL fflush
132		POP {pc}
133	p_print_ln:
134		PUSH {lr}
135		LDR r0, =msg_6
136		ADD r0, r0, #4
137		BL puts
138		MOV r0, #0
139		BL fflush
140		POP {pc}
141	p_throw_runtime_error:
142		BL p_print_string
143		MOV r0, #-1
144		BL exit
145	p_print_string:
146		PUSH {lr}
147		LDR r1, [r0]
148		ADD r2, r0, #4
149		LDR r0, =msg_7
150		ADD r0, r0, #4
151		BL printf
152		MOV r0, #0
153		BL fflush
154		POP {pc}
155	
===========================================================
-- Finished
