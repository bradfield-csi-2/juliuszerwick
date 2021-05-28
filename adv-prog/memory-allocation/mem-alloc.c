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
  size_t total_size;
  void *block;
  header_t *header;

  if (!size) {
    return NULL;
  }

  pthread_mutex_lock(&global_malloc_lock);
  header = get_free_block(size);

  if (header) {
    header->s.is_free = 0;
    pthread_mutex_unlock(&global_malloc_lock);
    return (void*)(header + 1)
  }

  total_size = sizeof(header_t) + size;
  // If size is positive, increments brk by size bytes thus allocating memory.
  // If size is negative, decrements brk by size bytes thus freeing memory.
  block = sbrk(total_size);

  // Check for failure from calling sbrk(size).
  if (block == (void*) -1) {
    pthread_mutex_unlock(&global_malloc_lock);
    return NULL;
  }

  header = block;
  header->s.size = size;
  header->s.is_free = 0;
  header->s.next = NULL;

  if (!head) {
    head = header;
  }

  if (tail) {
    tail->s.next = header;
  }

  tail = header;
  pthread_mutex_unlock(&global_malloc_lock);
  return (void*)(header + 1);
}

// Traverse linked list of memory blocks and find the first block
// that is both free and greater or equal to the requested size.
header_t *get_free_block(size_t size) {
  header_t *curr = head;

  while(curr) {
    if (curr->s.is_free && curr->s.size >= size) {
      return curr;
    }
    curr = curr->s.next;
  }

  return NULL;
}
