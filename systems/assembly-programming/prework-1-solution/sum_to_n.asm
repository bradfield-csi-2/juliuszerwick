section .text
global sum_to_n
sum_to_n:					; edi stores first argument (n) and eax stores return value
	xor eax, eax		; Sets value of eax to 0 with one instruction
.loop:
	add eax, edi		; Add values in eax and edi together, store result in eax -> eax = eax + edi
	sub edi, 1			; Subtract 1 from the value in edi (argument n) -> counting down from n to 0
	jg .loop				; jump when greater than -> jump to .loop label if edi value is greater than 1
	ret							; return from program
