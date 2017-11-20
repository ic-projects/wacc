-- Compiling...
-- Printing Assembly...
arrayBasic.s contents are:
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
9		LDR r5, =0
10		STR r5, [r4, #4]
11		LDR r5, =1
12		STR r5, [r4]
13		STR r4, [sp]
14		ADD sp, sp, #4
15		LDR r0, =0
16		POP {pc}
17		.ltorg
18	
===========================================================
-- Finished
