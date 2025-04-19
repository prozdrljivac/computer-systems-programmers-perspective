#include <stdio.h>

int main()
{
    int x = 0x87654321;

    // A.
    printf("A: %x\n", x & 0xFF);

    // B.
    printf("C: %x\n", x ^ 0xFFFFFF00);

    // C.
    printf("C: %x\n", x | 0xFF);

    return 0;
}