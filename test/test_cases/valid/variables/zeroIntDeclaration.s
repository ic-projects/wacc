-- Compiling...
-- Printing Assembly...
zeroIntDeclaration.s contents are:
===========================================================
0	.text
1	
2	.global main
3	main:
4		PUSH {lr}
5		SUB sp, sp, #4
6		LDR r4, =0
7		STR r4, [sp]
8		ADD sp, sp, #4
9		LDR r0, =0
10		POP {pc}
11		.ltorg
12	
===========================================================
-- Finished
