#include <stdio.h>

void cesar(char str[], int shift) {
  int i = 0;

  while (str[i] != '\0') {
    if (str[i] >= 'A' && str[i]<= 'Z') {
        // On enleve 97 dans le code ascii
        char c = str[i] - 'A';

        // On ajoute 97 à la nouveau code
        c += shift;

        // On fait un modulo 26
        c = c % 26;

        // Le nouveau caractere sera le nouveau code ascii + 'A'
        str[i] = c + 'A';
    }
    i++;
  }
  printf("%s\n", str);
}

int main()
{
    char str[] = "ABCDEFGHIJKMNOPQRSTUVWXYZ";
    cesar(str, 10);
    return 0;
}