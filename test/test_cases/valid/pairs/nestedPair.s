-- Compiling...
-- Printing Assembly...
nestedPair.s contents are:
===========================================================
0	.text
1	
2	.global main
3	main:
4		PUSH {lr}
5		SUB sp, sp, #8
6		LDR r0, =8
7		BL malloc
8		MOV r4, r0
9		LDR r5, =2
10		LDR r0, =4
11		BL malloc
12		STR r5, [r0]
13		STR r0, [r4]
14		LDR r5, =3
15		LDR r0, =4
16		BL malloc
17		STR r5, [r0]
18		STR r0, [r4, #4]
19		STR r4, [sp, #4]
20		LDR r0, =8
21		BL malloc
22		MOV r4, r0
23		LDR r5, =1
24		LDR r0, =4
25		BL malloc
26		STR r5, [r0]
27		STR r0, [r4]
28		LDR r5, [sp, #4]
29		LDR r0, =4
30		BL malloc
31		STR r5, [r0]
32		STR r0, [r4, #4]
33		STR r4, [sp]
34		ADD sp, sp, #8
35		LDR r0, =0
36		POP {pc}
37		.ltorg
38	
===========================================================
-- Finished
