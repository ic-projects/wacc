-- Compiling...
-- Printing Assembly...
functionDeclaration.s contents are:
===========================================================
0	.text
1	
2	.global main
3	f_f:
4		PUSH {lr}
5		LDR r4, =0
6		MOV r0, r4
7		POP {pc}
8		POP {pc}
9		.ltorg
10	main:
11		PUSH {lr}
12		LDR r0, =0
13		POP {pc}
14		.ltorg
15	
===========================================================
-- Finished
