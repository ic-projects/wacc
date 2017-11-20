-- Compiling...
-- Printing Assembly...
writeSnd.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 50
4		.ascii	"NullReferenceError: dereference a null reference\n\0"
5	msg_1:
6		.word 1
7		.ascii	"\0"
8	msg_2:
9		.word 5
10		.ascii	"%.*s\0"
11	
12	.text
13	
14	.global main
15	main:
16		PUSH {lr}
17		SUB sp, sp, #5
18		LDR r0, =8
19		BL malloc
20		MOV r4, r0
21		LDR r5, =10
22		LDR r0, =4
23		BL malloc
24		STR r5, [r0]
25		STR r0, [r4]
26		MOV r5, #'a'
27		LDR r0, =1
28		BL malloc
29		STRB r5, [r0]
30		STR r0, [r4, #4]
31		STR r4, [sp, #1]
32		LDR r4, [sp, #1]
33		MOV r0, r4
34		BL p_check_null_pointer
35		LDR r4, [r4, #4]
36		LDRSB r4, [r4]
37		STRB r4, [sp]
38		LDRSB r4, [sp]
39		MOV r0, r4
40		BL putchar
41		BL p_print_ln
42		MOV r4, #'Z'
43		LDR r5, [sp, #1]
44		MOV r0, r5
45		BL p_check_null_pointer
46		LDR r5, [r5, #4]
47		STRB r4, [r5]
48		LDR r4, [sp, #1]
49		MOV r0, r4
50		BL p_check_null_pointer
51		LDR r4, [r4, #4]
52		LDRSB r4, [r4]
53		STRB r4, [sp]
54		LDRSB r4, [sp]
55		MOV r0, r4
56		BL putchar
57		BL p_print_ln
58		ADD sp, sp, #5
59		LDR r0, =0
60		POP {pc}
61		.ltorg
62	p_check_null_pointer:
63		PUSH {lr}
64		CMP r0, #0
65		LDREQ r0, =msg_0
66		BLEQ p_throw_runtime_error
67		POP {pc}
68	p_print_ln:
69		PUSH {lr}
70		LDR r0, =msg_1
71		ADD r0, r0, #4
72		BL puts
73		MOV r0, #0
74		BL fflush
75		POP {pc}
76	p_throw_runtime_error:
77		BL p_print_string
78		MOV r0, #-1
79		BL exit
80	p_print_string:
81		PUSH {lr}
82		LDR r1, [r0]
83		ADD r2, r0, #4
84		LDR r0, =msg_2
85		ADD r0, r0, #4
86		BL printf
87		MOV r0, #0
88		BL fflush
89		POP {pc}
90	
===========================================================
-- Finished
