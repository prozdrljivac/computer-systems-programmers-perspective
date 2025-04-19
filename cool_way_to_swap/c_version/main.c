#include <stdio.h>

int main()
{
    int x = 1;
    int y = 5;

    printf("X is: %d\n", x);
    printf("Y is: %d\n", y);

    int *p = &x;
    int *q = &y;

    *q = *p ^ *q;
    *p = *p ^ *q;
    *q = *p ^ *q;

    printf("X is: %d\n", x);
    printf("Y is: %d\n", y);

    return 0;
}