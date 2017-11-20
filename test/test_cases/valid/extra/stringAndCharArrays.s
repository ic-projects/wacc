-- Compiling...
-- Printing Assembly...
stringAndCharArrays.s contents are:
===========================================================
0	.data
1	
2	msg_0:
3		.word 5
4		.ascii	"hello"
5	msg_1:
6		.word 4
7		.ascii	"test"
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
28		SUB sp, sp, #16
29		LDR r4, =msg_0
30		STR r4, [sp, #12]
31		LDR r0, =7
32		BL malloc
33		MOV r4, r0
34		MOV r5, #'a'
35		STRB r5, [r4, #4]
36		MOV r5, #'b'
37		STRB r5, [r4, #5]
38		MOV r5, #'c'
39		STRB r5, [r4, #6]
40		LDR r5, =3
41		STR r5, [r4]
42		STR r4, [sp, #8]
43		LDR r4, [sp, #12]
44		STR r4, [sp, #-4]!
45		BL f_f
46		ADD sp, sp, #4
47		MOV r4, r0
48		STR r4, [sp, #4]
49		LDR r4, [sp, #8]
50		STR r4, [sp, #-4]!
51		BL f_f
52		ADD sp, sp, #4
53		MOV r4, r0
54		STR r4, [sp, #12]
55		LDR r4, [sp, #12]
56		STR r4, [sp, #-4]!
57		BL f_f2
58		ADD sp, sp, #4
59		MOV r4, r0
60		STR r4, [sp, #8]
61		LDR r4, [sp, #8]
62		STR r4, [sp, #-4]!
63		BL f_f2
64		ADD sp, sp, #4
65		MOV r4, r0
66		STR r4, [sp, #4]
67		LDR r4, =msg_1
68		STR r4, [sp]
69		ADD sp, sp, #16
70		LDR r0, =0
71		POP {pc}
72		.ltorg
73	
===========================================================
-- Finished
