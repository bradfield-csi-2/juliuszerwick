	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 11, 0	sdk_version 11, 1
	.globl	_pagecount              ## -- Begin function pagecount
	.p2align	4, 0x90
_pagecount:                             ## @pagecount
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	movq	%rdi, %rax
	xorl	%edx, %edx
	divq	%rsi
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function
	.section	__TEXT,__literal16,16byte_literals
	.p2align	4               ## -- Begin function main
LCPI1_0:
	.long	1127219200              ## 0x43300000
	.long	1160773632              ## 0x45300000
	.long	0                       ## 0x0
	.long	0                       ## 0x0
LCPI1_1:
	.quad	4841369599423283200     ## double 4503599627370496
	.quad	4985484787499139072     ## double 1.9342813113834067E+25
	.section	__TEXT,__literal8,8byte_literals
	.p2align	3
LCPI1_2:
	.quad	4696837146684686336     ## double 1.0E+6
LCPI1_3:
	.quad	4741671816366391296     ## double 1.0E+9
LCPI1_4:
	.quad	4711630319722168320     ## double 1.0E+7
	.section	__TEXT,__text,regular,pure_instructions
	.globl	_main
	.p2align	4, 0x90
_main:                                  ## @main
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	pushq	%r15
	pushq	%r14
	pushq	%r13
	pushq	%r12
	pushq	%rbx
	subq	$40, %rsp
	.cfi_offset %rbx, -56
	.cfi_offset %r12, -48
	.cfi_offset %r13, -40
	.cfi_offset %r14, -32
	.cfi_offset %r15, -24
	xorl	%ebx, %ebx
	movl	$10000000, %r14d        ## imm = 0x989680
	callq	_clock
	movq	%rax, -72(%rbp)         ## 8-byte Spill
	movabsq	$-6148914691236517205, %r15 ## imm = 0xAAAAAAAAAAAAAAAB
	leaq	l___const.main.msizes(%rip), %r12
	leaq	l___const.main.psizes(%rip), %rsi
	xorl	%ecx, %ecx
	xorl	%r13d, %r13d
	.p2align	4, 0x90
LBB1_1:                                 ## =>This Inner Loop Header: Depth=1
	movq	%rbx, %rax
	mulq	%r15
	shlq	$2, %rdx
	andq	$-8, %rdx
	leaq	(%rdx,%rdx,2), %rax
	movq	%rcx, %rdx
	subq	%rax, %rdx
	movl	(%r12,%rdx), %eax
	addl	(%rsi,%rdx), %eax
	leal	1(%r13,%rax), %r13d
	addq	$8, %rcx
	incq	%rbx
	decl	%r14d
	jne	LBB1_1
## %bb.2:
	callq	_clock
	movq	%rax, -64(%rbp)         ## 8-byte Spill
	movl	$10000000, -44(%rbp)    ## 4-byte Folded Spill
                                        ## imm = 0x989680
	xorl	%ebx, %ebx
	callq	_clock
	movq	%rax, -56(%rbp)         ## 8-byte Spill
	xorl	%r14d, %r14d
	.p2align	4, 0x90
LBB1_3:                                 ## =>This Inner Loop Header: Depth=1
	movq	%rbx, %rax
	mulq	%r15
	shlq	$2, %rdx
	andq	$-8, %rdx
	leaq	(%rdx,%rdx,2), %rax
	movq	%r14, %rcx
	subq	%rax, %rcx
	movq	(%r12,%rcx), %r15
	leaq	l___const.main.psizes(%rip), %rax
	movq	(%rax,%rcx), %r12
	movq	%r15, %rdi
	movq	%r12, %rsi
	callq	_pagecount
	addl	%r15d, %r12d
	movabsq	$-6148914691236517205, %r15 ## imm = 0xAAAAAAAAAAAAAAAB
	addl	%r12d, %eax
	leaq	l___const.main.msizes(%rip), %r12
	addl	%eax, %r13d
	addq	$8, %r14
	incq	%rbx
	decl	-44(%rbp)               ## 4-byte Folded Spill
	jne	LBB1_3
## %bb.4:
	callq	_clock
	movq	-72(%rbp), %rcx         ## 8-byte Reload
	subq	-64(%rbp), %rcx         ## 8-byte Folded Reload
	subq	-56(%rbp), %rcx         ## 8-byte Folded Reload
	addq	%rax, %rcx
	movq	%rcx, %xmm1
	punpckldq	LCPI1_0(%rip), %xmm1 ## xmm1 = xmm1[0],mem[0],xmm1[1],mem[1]
	subpd	LCPI1_1(%rip), %xmm1
	movapd	%xmm1, %xmm0
	unpckhpd	%xmm1, %xmm0    ## xmm0 = xmm0[1],xmm1[1]
	addsd	%xmm1, %xmm0
	divsd	LCPI1_2(%rip), %xmm0
	movsd	LCPI1_3(%rip), %xmm1    ## xmm1 = mem[0],zero
	mulsd	%xmm0, %xmm1
	divsd	LCPI1_4(%rip), %xmm1
	leaq	L_.str(%rip), %rdi
	movl	$10000000, %esi         ## imm = 0x989680
	movb	$2, %al
	callq	_printf
	movl	%r13d, %eax
	addq	$40, %rsp
	popq	%rbx
	popq	%r12
	popq	%r13
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function
	.section	__TEXT,__const
	.p2align	4               ## @__const.main.msizes
l___const.main.msizes:
	.quad	4294967296              ## 0x100000000
	.quad	1099511627776           ## 0x10000000000
	.quad	4503599627370496        ## 0x10000000000000

	.p2align	4               ## @__const.main.psizes
l___const.main.psizes:
	.quad	4096                    ## 0x1000
	.quad	65536                   ## 0x10000
	.quad	4294967296              ## 0x100000000

	.section	__TEXT,__cstring,cstring_literals
L_.str:                                 ## @.str
	.asciz	"%.2fs to run %d tests (%.2fns per test)\n"

.subsections_via_symbols
