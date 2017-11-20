-- Compiling...
-- Printing Assembly...
skip.s contents are:
===========================================================
0	.text
1	
2	.global main
3	main:
4		PUSH {lr}
5		LDR r0, =0
6		POP {pc}
7		.ltorg
8	
===========================================================
-- Finished
