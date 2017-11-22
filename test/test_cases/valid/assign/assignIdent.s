-- Compiling...
-- Printing Assembly...
assignIdent.s contents are:
===========================================================
0	.text
1	
2	.global main
3	main:
4		PUSH {lr}
5		SUB sp, sp, #4
6		LDR r4, =0
7		STR r4, [sp]
8		LDR r4, =1
9		STR r4, [sp]
10		LDR r4, [sp]
11		MOV r0, r4
12		BL exit
13		ADD sp, sp, #4
14		LDR r0, =0
15		POP {pc}
16		.ltorg
17	
===========================================================
-- Finished
