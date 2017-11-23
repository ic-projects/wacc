-- Compiling...
-- Printing Assembly...
arrayLookup2.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 44
4		.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
5	msg_1:
6		.word 45
7		.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
8	msg_2:
9		.word 3
10		.ascii	"%d\0"
11	msg_3:
12		.word 1
13		.ascii	"\0"
14	msg_4:
15		.word 82
16		.ascii	"OverflowError: the result is too small/large to store in a 4-byte signed-integer.\n"
17	msg_5:
18		.word 5
19		.ascii	"%.*s\0"
20	
21	.text
22	
23	.global main
24	main:
25		PUSH {lr}
26		SUB sp, sp, #4
27		LDR r0, =20
28		BL malloc
29		MOV r4, r0
30		LDR r5, =43
31		STR r5, [r4, #4]
32		LDR r5, =2
33		STR r5, [r4, #8]
34		LDR r5, =18
35		STR r5, [r4, #12]
36		LDR r5, =1
37		STR r5, [r4, #16]
38		LDR r5, =4
39		STR r5, [r4]
40		STR r4, [sp]
41		ADD r4, sp, #0
42		LDR r5, =0
43		LDR r4, [r4]
44		MOV r0, r5
45		MOV r1, r4
46		BL p_check_array_bounds
47		ADD r4, r4, #4
48		ADD r4, r4, r5, LSL #2
49		LDR r4, [r4]
50		MOV r0, r4
51		BL p_print_int
52		BL p_print_ln
53		ADD r4, sp, #0
54		LDR r5, =1
55		LDR r4, [r4]
56		MOV r0, r5
57		MOV r1, r4
58		BL p_check_array_bounds
59		ADD r4, r4, #4
60		ADD r4, r4, r5, LSL #2
61		LDR r4, [r4]
62		MOV r0, r4
63		BL p_print_int
64		BL p_print_ln
65		ADD r4, sp, #0
66		LDR r5, =2
67		LDR r4, [r4]
68		MOV r0, r5
69		MOV r1, r4
70		BL p_check_array_bounds
71		ADD r4, r4, #4
72		ADD r4, r4, r5, LSL #2
73		LDR r4, [r4]
74		MOV r0, r4
75		BL p_print_int
76		BL p_print_ln
77		ADD r4, sp, #0
78		LDR r5, =1
79		LDR r6, =1
80		ADDS r5, r5, r6
81		BLVS p_throw_overflow_error
82		LDR r6, =2
83		ADDS r5, r5, r6
84		BLVS p_throw_overflow_error
85		LDR r6, =1
86		SUBS r5, r5, r6
87		BLVS p_throw_overflow_error
88		LDR r4, [r4]
89		MOV r0, r5
90		MOV r1, r4
91		BL p_check_array_bounds
92		ADD r4, r4, #4
93		ADD r4, r4, r5, LSL #2
94		LDR r4, [r4]
95		MOV r0, r4
96		BL p_print_int
97		BL p_print_ln
98		ADD sp, sp, #4
99		LDR r0, =0
100		POP {pc}
101		.ltorg
102	p_check_array_bounds:
103		PUSH {lr}
104		CMP r0, #0
105		LDRLT r0, =msg_0
106		BLLT p_throw_runtime_error
107		LDR r1, [r1]
108		CMP r0, r1
109		LDRCS r0, =msg_1
110		BLCS p_throw_runtime_error
111		POP {pc}
112	p_print_int:
113		PUSH {lr}
114		MOV r1, r0
115		LDR r0, =msg_2
116		ADD r0, r0, #4
117		BL printf
118		MOV r0, #0
119		BL fflush
120		POP {pc}
121	p_print_ln:
122		PUSH {lr}
123		LDR r0, =msg_3
124		ADD r0, r0, #4
125		BL puts
126		MOV r0, #0
127		BL fflush
128		POP {pc}
129	p_throw_overflow_error:
130		LDR r0, =msg_4
131		BL p_throw_runtime_error
132	p_throw_runtime_error:
133		BL p_print_string
134		MOV r0, #-1
135		BL exit
136	p_print_string:
137		PUSH {lr}
138		LDR r1, [r0]
139		ADD r2, r0, #4
140		LDR r0, =msg_5
141		ADD r0, r0, #4
142		BL printf
143		MOV r0, #0
144		BL fflush
145		POP {pc}
146	
===========================================================
-- Finished
