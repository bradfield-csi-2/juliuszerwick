#include <stdio.h>
#include <stdlib.h>

// Stub variable of size 16 bytes.
typedef char ALIGN[16];

// Union ensures that header is aligned to 16 bytes due to stub.
union header {
  // Header for each block of memory.
  struct header_t {
    // Size of memory block.
    size_t size;
    // Indicates whether memory is free or not.
    unsigned is_free;
    // Pointer to next memory block in linked list of blocks.
    struct header_t *next;
  } s;
  ALIGN stub;
};

typedef union header header_t;

header_t *head, *tail;

pthread_mutex_t global_malloc_lock;

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
