-- Compiling...
-- Printing Assembly...
readArray.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 31
4		.ascii	"input an integer to continue..."
5	msg_1:
6		.word 44
7		.ascii	"ArrayIndexOutOfBoundsError: negative index\n\0"
8	msg_2:
9		.word 45
10		.ascii	"ArrayIndexOutOfBoundsError: index too large\n\0"
11	msg_3:
12		.word 3
13		.ascii	"%d\0"
14	msg_4:
15		.word 1
16		.ascii	"\0"
17	msg_5:
18		.word 5
19		.ascii	"%.*s\0"
20	msg_6:
21		.word 3
22		.ascii	"%d\0"
23	
24	.text
25	
26	.global main
27	main:
28		PUSH {lr}
29		SUB sp, sp, #4
30		LDR r0, =16
31		BL malloc
32		MOV r4, r0
33		LDR r5, =0
34		STR r5, [r4, #4]
35		LDR r5, =1
36		STR r5, [r4, #8]
37		LDR r5, =2
38		STR r5, [r4, #12]
39		LDR r5, =3
40		STR r5, [r4]
41		STR r4, [sp]
42		ADD r4, sp, #0
43		LDR r5, =0
44		LDR r4, [r4]
45		MOV r0, r5
46		MOV r1, r4
47		BL p_check_array_bounds
48		ADD r4, r4, #4
49		ADD r4, r4, r5, LSL #2
50		LDR r4, [r4]
51		MOV r0, r4
52		BL p_print_int
53		BL p_print_ln
54		LDR r4, =msg_0
55		MOV r0, r4
56		BL p_print_string
57		BL p_print_ln
58		ADD r4, sp, #0
59		LDR r5, =0
60		LDR r4, [r4]
61		MOV r0, r5
62		MOV r1, r4
63		BL p_check_array_bounds
64		ADD r4, r4, #4
65		ADD r4, r4, r5, LSL #2
66		MOV r4, r4
67		MOV r0, r4
68		BL p_read_int
69		ADD r4, sp, #0
70		LDR r5, =0
71		LDR r4, [r4]
72		MOV r0, r5
73		MOV r1, r4
74		BL p_check_array_bounds
75		ADD r4, r4, #4
76		ADD r4, r4, r5, LSL #2
77		LDR r4, [r4]
78		MOV r0, r4
79		BL p_print_int
80		BL p_print_ln
81		ADD sp, sp, #4
82		LDR r0, =0
83		POP {pc}
84		.ltorg
85	p_check_array_bounds:
86		PUSH {lr}
87		CMP r0, #0
88		LDRLT r0, =msg_1
89		BLLT p_throw_runtime_error
90		LDR r1, [r1]
91		CMP r0, r1
92		LDRCS r0, =msg_2
93		BLCS p_throw_runtime_error
94		POP {pc}
95	p_print_int:
96		PUSH {lr}
97		MOV r1, r0
98		LDR r0, =msg_3
99		ADD r0, r0, #4
100		BL printf
101		MOV r0, #0
102		BL fflush
103		POP {pc}
104	p_print_ln:
105		PUSH {lr}
106		LDR r0, =msg_4
107		ADD r0, r0, #4
108		BL puts
109		MOV r0, #0
110		BL fflush
111		POP {pc}
112	p_print_string:
113		PUSH {lr}
114		LDR r1, [r0]
115		ADD r2, r0, #4
116		LDR r0, =msg_5
117		ADD r0, r0, #4
118		BL printf
119		MOV r0, #0
120		BL fflush
121		POP {pc}
122	p_read_int:
123		PUSH {lr}
124		MOV r1, r0
125		LDR r0, =msg_6
126		ADD r0, r0, #4
127		BL scanf
128		POP {pc}
129	p_throw_runtime_error:
130		BL p_print_string
131		MOV r0, #-1
132		BL exit
133	
===========================================================
-- Finished
