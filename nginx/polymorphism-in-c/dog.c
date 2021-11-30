#include "animal.h"

static const char *sound(void)
{
    return "arf!";
}

const struct animal_vtable_ DOG[] = { { sound } };