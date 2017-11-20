-- Compiling...
-- Printing Assembly...
ifBasic.s contents are:
===========================================================
0	.text
1	
2	.global main
3	main:
4		PUSH {lr}
5		MOV r4, #1
6		CMP r4, #0
7		BEQ L0
8		B L1
9	L0:
10	L1:
11		LDR r0, =0
12		POP {pc}
13		.ltorg
14	
===========================================================
-- Finished
