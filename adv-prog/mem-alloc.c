#include <stdio.h>
#include <stdlib.h>

// Header for each block of memory.
struct header_t {
  // Size of memory block.
  size_t size;
  // Indicates whether memory is free or not.
  unsigned is_free;
  // Pointer to next memory block in linked list of blocks.
  struct header_t *next;
}

void main() {
}

void *malloc(size_t size) {
  void *block;
  // If size is positive, increments brk by size bytes thus allocating memory.
  // If size is negative, decrements brk by size bytes thus freeing memory.
  block = sbrk(size);

  // Check for failure from calling sbrk(size).
  if (block == (void*) -1) {
    return NULL;
  }

  return block;
}
