-- Compiling...
-- Printing Assembly...
createPair02.s contents are:
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
9		MOV r5, #'a'
10		LDR r0, =1
11		BL malloc
12		STRB r5, [r0]
13		STR r0, [r4]
14		MOV r5, #'b'
15		LDR r0, =1
16		BL malloc
17		STRB r5, [r0]
18		STR r0, [r4, #4]
19		STR r4, [sp]
20		ADD sp, sp, #4
21		LDR r0, =0
22		POP {pc}
23		.ltorg
24	
===========================================================
-- Finished
