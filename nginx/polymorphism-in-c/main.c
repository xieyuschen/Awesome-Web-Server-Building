#include "animal.h"
#include <stdio.h>

int main(void)
{
    struct animal kitty = { CAT, "Kitty" };
    struct animal lassie = { DOG, "Lassie" };

    printf("%s says %s\n", kitty.name, animal_sound(&kitty));
    printf("%s says %s\n", lassie.name, animal_sound(&lassie));

    return 0;
}