-- Compiling...
-- Printing Assembly...
doubleFree.s contents are:
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
14		SUB sp, sp, #4
15		LDR r0, =8
16		BL malloc
17		MOV r4, r0
18		LDR r5, =10
19		LDR r0, =4
20		BL malloc
21		STR r5, [r0]
22		STR r0, [r4]
23		MOV r5, #'a'
24		LDR r0, =1
25		BL malloc
26		STRB r5, [r0]
27		STR r0, [r4, #4]
28		STR r4, [sp]
29		LDR r4, [sp]
30		MOV r0, r4
31		BL p_free_pair
32		LDR r4, [sp]
33		MOV r0, r4
34		BL p_free_pair
35		ADD sp, sp, #4
36		LDR r0, =0
37		POP {pc}
38		.ltorg
39	p_free_pair:
40		PUSH {lr}
41		CMP r0, #0
42		LDREQ r0, =msg_0
43		BEQ p_throw_runtime_error
44		PUSH {r0}
45		LDR r0, [r0]
46		BL free
47		LDR r0, [sp]
48		LDR r0, [r0, #4]
49		BL free
50		POP {r0}
51		BL free
52		POP {pc}
53	p_throw_runtime_error:
54		BL p_print_string
55		MOV r0, #-1
56		BL exit
57	p_print_string:
58		PUSH {lr}
59		LDR r1, [r0]
60		ADD r2, r0, #4
61		LDR r0, =msg_1
62		ADD r0, r0, #4
63		BL printf
64		MOV r0, #0
65		BL fflush
66		POP {pc}
67	
===========================================================
-- Finished
