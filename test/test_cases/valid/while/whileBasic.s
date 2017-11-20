-- Compiling...
-- Printing Assembly...
whileBasic.s contents are:
===========================================================
0	.text
1	
2	.global main
3	main:
4		PUSH {lr}
5		B L0
6	L1:
7	L0:
8		MOV r4, #0
9		CMP r4, #1
10		BEQ L1
11		LDR r0, =0
12		POP {pc}
13		.ltorg
14	
===========================================================
-- Finished
