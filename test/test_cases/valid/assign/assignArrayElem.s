-- Compiling...
-- Printing Assembly...
assignArrayElem.s contents are:
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
9		.word 5
10		.ascii	"%.*s\0"
11	
12	.text
13	
14	.global main
15	main:
16		PUSH {lr}
17		SUB sp, sp, #4
18		LDR r0, =8
19		BL malloc
20		MOV r4, r0
21		LDR r5, =0
22		STR r5, [r4, #4]
23		LDR r5, =1
24		STR r5, [r4]
25		STR r4, [sp]
26		LDR r4, =1
27		ADD r5, sp, #0
28		LDR r6, =0
29		LDR r5, [r5]
30		MOV r0, r6
31		MOV r1, r5
32		BL p_check_array_bounds
33		ADD r5, r5, #4
34		ADD r5, r5, r6, LSL #2
35		STR r4, [r5]
36		ADD r4, sp, #0
37		LDR r6, =0
38		LDR r4, [r4]
39		MOV r0, r6
40		MOV r1, r4
41		BL p_check_array_bounds
42		ADD r4, r4, #4
43		ADD r4, r4, r6, LSL #2
44		LDR r4, [r4]
45		MOV r0, r4
46		BL exit
47		ADD sp, sp, #4
48		LDR r0, =0
49		POP {pc}
50		.ltorg
51	p_check_array_bounds:
52		PUSH {lr}
53		CMP r0, #0
54		LDRLT r0, =msg_0
55		BLLT p_throw_runtime_error
56		LDR r1, [r1]
57		CMP r0, r1
58		LDRCS r0, =msg_1
59		BLCS p_throw_runtime_error
60		POP {pc}
61	p_throw_runtime_error:
62		BL p_print_string
63		MOV r0, #-1
64		BL exit
65	p_print_string:
66		PUSH {lr}
67		LDR r1, [r0]
68		ADD r2, r0, #4
69		LDR r0, =msg_2
70		ADD r0, r0, #4
71		BL printf
72		MOV r0, #0
73		BL fflush
74		POP {pc}
75	
===========================================================
-- Finished
