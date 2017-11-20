-- Compiling...
-- Printing Assembly...
hiddenDoubleFree.s contents are:
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
28		STR r4, [sp, #4]
29		LDR r4, [sp, #4]
30		STR r4, [sp]
31		LDR r4, [sp, #4]
32		MOV r0, r4
33		BL p_free_pair
34		LDR r4, [sp]
35		MOV r0, r4
36		BL p_free_pair
37		ADD sp, sp, #8
38		LDR r0, =0
39		POP {pc}
40		.ltorg
41	p_free_pair:
42		PUSH {lr}
43		CMP r0, #0
44		LDREQ r0, =msg_0
45		BEQ p_throw_runtime_error
46		PUSH {r0}
47		LDR r0, [r0]
48		BL free
49		LDR r0, [sp]
50		LDR r0, [r0, #4]
51		BL free
52		POP {r0}
53		BL free
54		POP {pc}
55	p_throw_runtime_error:
56		BL p_print_string
57		MOV r0, #-1
58		BL exit
59	p_print_string:
60		PUSH {lr}
61		LDR r1, [r0]
62		ADD r2, r0, #4
63		LDR r0, =msg_1
64		ADD r0, r0, #4
65		BL printf
66		MOV r0, #0
67		BL fflush
68		POP {pc}
69	
===========================================================
-- Finished
