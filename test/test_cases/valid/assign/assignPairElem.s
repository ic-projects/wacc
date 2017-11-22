-- Compiling...
-- Printing Assembly...
assignPairElem.s contents are:
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
12	main:
13		PUSH {lr}
14		SUB sp, sp, #8
15		LDR r0, =8
16		BL malloc
17		MOV r4, r0
18		LDR r5, =0
19		LDR r0, =4
20		BL malloc
21		STR r5, [r0]
22		STR r0, [r4]
23		LDR r5, =0
24		LDR r0, =4
25		BL malloc
26		STR r5, [r0]
27		STR r0, [r4, #4]
28		STR r4, [sp, #4]
29		LDR r4, =1
30		LDR r5, [sp, #4]
31		MOV r0, r5
32		BL p_check_null_pointer
33		LDR r5, [r5]
34		STR r4, [r5]
35		LDR r4, [sp, #4]
36		MOV r0, r4
37		BL p_check_null_pointer
38		LDR r4, [r4]
39		LDR r4, [r4]
40		STR r4, [sp]
41		LDR r4, [sp]
42		MOV r0, r4
43		BL exit
44		ADD sp, sp, #8
45		LDR r0, =0
46		POP {pc}
47		.ltorg
48	p_check_null_pointer:
49		PUSH {lr}
50		CMP r0, #0
51		LDREQ r0, =msg_0
52		BLEQ p_throw_runtime_error
53		POP {pc}
54	p_throw_runtime_error:
55		BL p_print_string
56		MOV r0, #-1
57		BL exit
58	p_print_string:
59		PUSH {lr}
60		LDR r1, [r0]
61		ADD r2, r0, #4
62		LDR r0, =msg_1
63		ADD r0, r0, #4
64		BL printf
65		MOV r0, #0
66		BL fflush
67		POP {pc}
68	
===========================================================
-- Finished
