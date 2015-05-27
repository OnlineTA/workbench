#include <stdlib.h>
#include <stdio.h>

#define STACK_ERROR_UNKNOWN   (1)
#define STACK_ERROR_UNDERFLOW (2)
#define STACK_ERROR_BADINT    (3)

typedef struct node {
  struct node *tail;
  int value;
} stack_t;

int
stack_push(stack_t **stack, int value)
{
  stack_t *new_stack = (stack_t*)malloc(sizeof(stack_t));
  if (new_stack == NULL)
  {
    // Couldn't malloc, abandon the ship!
    exit(EXIT_FAILURE);
  }

  new_stack->value = value;
  new_stack->tail = *stack;

  *stack = new_stack;

  return 0;
}

int
stack_pop(stack_t **stack, int *value)
{
  if (*stack == NULL)
  {
    return STACK_ERROR_UNDERFLOW;
  }

  *value = (*stack)->value;
  *stack = (*stack)->tail;

  free(*stack);

  return 0;
}

int
stack_peek(stack_t **stack, int *value)
{
  if (*stack == NULL) {
    return STACK_ERROR_UNDERFLOW;
  } else {
    *value = (*stack)->value;
    return 0;
  }
}

int
plus(stack_t **stack)
{
  int x, y;

  if (stack_pop(stack, &x) != 0)
  {
    return STACK_ERROR_UNDERFLOW;
  }

  if (stack_pop(stack, &y) != 0)
  {
    return STACK_ERROR_UNDERFLOW;
  }

  x += y;

  if (stack_push(stack, x) != 0)
  {
    return STACK_ERROR_UNKNOWN;
  }

  return 0;
}

int
main()
{
  stack_t *stack;
  char c;
  int v;

  stack = NULL;

  while(1)
  {
    c = getc(stdin);

    switch (c) {
    case '+':
      if (plus(&stack) != 0)
      {
        exit(STACK_ERROR_UNKNOWN);
      }
      break;
    case 'p':
      if (stack_peek(&stack, &v) != 0)
      {
        exit(STACK_ERROR_UNDERFLOW);
      }
      if (fprintf(stdout, "%d\n", v) < 0)
      {
        exit(STACK_ERROR_UNKNOWN);
      }
      break;
    case '\n':
      continue;
    case EOF:
      exit(0);
      break;
    default:
      if (ungetc(c, stdin) != c)
      {
        exit(STACK_ERROR_UNKNOWN);
      }
      if (fscanf(stdin, "%d\n", &v) != 1)
      {
        exit(STACK_ERROR_BADINT);
      }

      if (stack_push(&stack, v) != 0)
      {
        exit(STACK_ERROR_UNKNOWN);
      }
      break;
    }
  }

  return 0;
}
