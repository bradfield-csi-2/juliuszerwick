section .text
global sum_to_n
sum_to_n:
	mov		rax, 0		; set return value named total to 0
	mov		rsi, 0		; set variable named i to 0 in %rsi
	;xor		rax, rax
.L1:
	add		rax, rsi	; add value in %rsi to return value in %rax
	inc		rsi				; increment variable named i
	cmp		rsi, rdi	; compare return value with first argument
	je		.L2				; if equal, jump to code to return (break out of loop)
	jmp		.L1				; jump back up to .L1 label to simulate a loop
.L2:
	ret


;section .text
;global sum_to_n
;sum_to_n:
;	xor eax, eax
;.loop:
;	add eax, edi
;	sub edi, 1
;	jg .loop	
;	ret
