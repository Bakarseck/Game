<h1 align="center"> Tic Tac Toe</h1>

## Rules
* Premièrement, vous devez renseignez les pseudos des joueurs 1 et 2 respectivent.
* Après, chaque joueur doit donnez la position où il veut poser, par un chiffre allant de 1 à 9.
* Le premier joueur qui aligne son symbole trois fois de suite gagne la partie.

## Description

## Usage

```bash
    go run .
```

## Fonctions

Le Jeu fonctionne grâce à ces fonctions suivantes :

```go
    ClearConsole()
    Play()
    PrintBoard()
    Saisie()
    GetPosition()
    Size()
    Win()
```

## Explications des Fonctions

* ClearConsole:<hr>
C'est la fonction qui utilise le bibliothèque os/exec pour executer la commande `clear` et il redirige la sortie sur la sortie standard.

* Play():<hr>
C'est la fonctionne qui controle le jeu, il utilise une boucle sur la condition que le tableau n'est pas completement rempli.

* Size():<hr>
C'est la fonction qui calcule le nombre de `X` et de `O` qui existent dans le tableau.

* PrintBoard():<hr>
C'est la fonction qui affiche le tableau en tant que tableau de jeu 