#include "animal.h"

static const char *sound(void)
{
    return "meow!";
}

const struct animal_vtable_ CAT[] = { { sound } };