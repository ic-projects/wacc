-- Compiling...
-- Printing Assembly...
charAssignment.s contents are:
===========================================================
0	.text
1	
2	.global main
3	main:
4		PUSH {lr}
5		SUB sp, sp, #1
6		MOV r4, #'a'
7		STRB r4, [sp]
8		MOV r4, #'Z'
9		STRB r4, [sp]
10		ADD sp, sp, #1
11		LDR r0, =0
12		POP {pc}
13		.ltorg
14	
===========================================================
-- Finished
