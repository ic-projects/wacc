-- Compiling...
-- Printing Assembly...
stringArrays2.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 3
4		.ascii	"snd"
5	msg_1:
6		.word 44
7		.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
8	msg_2:
9		.word 45
10		.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
11	msg_3:
12		.word 5
13		.ascii	"true\0"
14	msg_4:
15		.word 6
16		.ascii	"false\0"
17	msg_5:
18		.word 1
19		.ascii	"\0"
20	msg_6:
21		.word 5
22		.ascii	"%.*s\0"
23	
24	.text
25	
26	.global main
27	main:
28		PUSH {lr}
29		SUB sp, sp, #12
30		LDR r0, =7
31		BL malloc
32		MOV r4, r0
33		MOV r5, #'f'
34		STRB r5, [r4, #4]
35		MOV r5, #'s'
36		STRB r5, [r4, #5]
37		MOV r5, #'t'
38		STRB r5, [r4, #6]
39		LDR r5, =3
40		STR r5, [r4]
41		STR r4, [sp, #8]
42		LDR r4, =msg_0
43		STR r4, [sp, #4]
44		LDR r0, =12
45		BL malloc
46		MOV r4, r0
47		LDR r5, [sp, #8]
48		STR r5, [r4, #4]
49		LDR r5, [sp, #4]
50		STR r5, [r4, #8]
51		LDR r5, =2
52		STR r5, [r4]
53		STR r4, [sp]
54		ADD r4, sp, #0
55		LDR r5, =0
56		LDR r4, [r4]
57		MOV r0, r5
58		MOV r1, r4
59		BL p_check_array_bounds
60		ADD r4, r4, #4
61		ADD r4, r4, r5, LSL #2
62		LDR r4, [r4]
63		LDR r5, [sp, #8]
64		CMP r4, r5
65		MOVEQ r4, #1
66		MOVNE r4, #0
67		MOV r0, r4
68		BL p_print_bool
69		BL p_print_ln
70		ADD r4, sp, #0
71		LDR r5, =0
72		LDR r4, [r4]
73		MOV r0, r5
74		MOV r1, r4
75		BL p_check_array_bounds
76		ADD r4, r4, #4
77		ADD r4, r4, r5, LSL #2
78		LDR r5, =1
79		LDR r4, [r4]
80		MOV r0, r5
81		MOV r1, r4
82		BL p_check_array_bounds
83		ADD r4, r4, #4
84		ADD r4, r4, r5
85		LDRSB r4, [r4]
86		MOV r5, #'f'
87		CMP r4, r5
88		MOVEQ r4, #1
89		MOVNE r4, #0
90		MOV r0, r4
91		BL p_print_bool
92		BL p_print_ln
93		ADD r4, sp, #0
94		LDR r5, =1
95		LDR r4, [r4]
96		MOV r0, r5
97		MOV r1, r4
98		BL p_check_array_bounds
99		ADD r4, r4, #4
100		ADD r4, r4, r5, LSL #2
101		LDR r5, =2
102		LDR r4, [r4]
103		MOV r0, r5
104		MOV r1, r4
105		BL p_check_array_bounds
106		ADD r4, r4, #4
107		ADD r4, r4, r5
108		LDRSB r4, [r4]
109		MOV r5, #'d'
110		CMP r4, r5
111		MOVEQ r4, #1
112		MOVNE r4, #0
113		MOV r0, r4
114		BL p_print_bool
115		BL p_print_ln
116		ADD sp, sp, #12
117		LDR r0, =0
118		POP {pc}
119		.ltorg
120	p_check_array_bounds:
121		PUSH {lr}
122		CMP r0, #0
123		LDRLT r0, =msg_1
124		BLLT p_throw_runtime_error
125		LDR r1, [r1]
126		CMP r0, r1
127		LDRCS r0, =msg_2
128		BLCS p_throw_runtime_error
129		POP {pc}
130	p_print_bool:
131		PUSH {lr}
132		CMP r0, #0
133		LDRNE r0, =msg_3
134		LDREQ r0, =msg_4
135		ADD r0, r0, #4
136		BL printf
137		MOV r0, #0
138		BL fflush
139		POP {pc}
140	p_print_ln:
141		PUSH {lr}
142		LDR r0, =msg_5
143		ADD r0, r0, #4
144		BL puts
145		MOV r0, #0
146		BL fflush
147		POP {pc}
148	p_throw_runtime_error:
149		BL p_print_string
150		MOV r0, #-1
151		BL exit
152	p_print_string:
153		PUSH {lr}
154		LDR r1, [r0]
155		ADD r2, r0, #4
156		LDR r0, =msg_6
157		ADD r0, r0, #4
158		BL printf
159		MOV r0, #0
160		BL fflush
161		POP {pc}
162	
===========================================================
-- Finished
