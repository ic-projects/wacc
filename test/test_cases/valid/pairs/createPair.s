-- Compiling...
-- Printing Assembly...
createPair.s contents are:
===========================================================
0	.text
1	
2	.global main
3	main:
4		PUSH {lr}
5		SUB sp, sp, #4
6		LDR r0, =8
7		BL malloc
8		MOV r4, r0
9		LDR r5, =10
10		LDR r0, =4
11		BL malloc
12		STR r5, [r0]
13		STR r0, [r4]
14		LDR r5, =3
15		LDR r0, =4
16		BL malloc
17		STR r5, [r0]
18		STR r0, [r4, #4]
19		STR r4, [sp]
20		ADD sp, sp, #4
21		LDR r0, =0
22		POP {pc}
23		.ltorg
24	
===========================================================
-- Finished