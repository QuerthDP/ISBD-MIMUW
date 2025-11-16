# Laboratorium z przedmiotu Implementacja systemów baz danych
Prowadzący: Krzysztof Smogór (krzysztof.smogor@gmail.com)

- [Laboratorium z przedmiotu Implementacja systemów baz danych](#laboratorium-z-przedmiotu-implementacja-systemów-baz-danych)
  - [Opis przedmiotu](#opis-przedmiotu)
  - [Zasady zaliczenia](#zasady-zaliczenia)
    - [Oddawanie projektów](#oddawanie-projektów)
  - [Tabela punktacji za projekty](#tabela-punktacji-za-projekty)
  - [Wymagania techniczne](#wymagania-techniczne)



## Opis przedmiotu

W ramach laboratorium będziemy implementować prosty system bazodanowy przeznaczony do zapytań analitycznych (brak transakcji). W ramach jego funkcjonalności będzie:
* import danych (z pliku CSV)
* wybieranie danych do zapytania (projekcja)
* transformacje na danych (operatory i funkcje)

## Zasady zaliczenia

Za laboratorium student może otrzymać maksymalnie 100 pkt.
Przedmiot podzielony jest na 4 projekty punktowane tak jak w [tabelce poniżej](#tabela-punktacji-za-projekty).
Przedmiot zdaje się po uzyskaniu >50 punktów.

| Liczba punktów | Ocena |
| -------------- | ----- |
| >50            | 3.0   |
| >60            | 3.5   |
| >70            | 4.0   |
| >80            | 4.5   |
| >90            | 5.0   |


Projekty należy wykonać samodzielnie (studenci oceniani są indywidualnie).
Mimo samodzielności wykonywania projektów zachęcam do wzajemnej dyskusji w gronie studentów oraz ze mną (droga mailową lub na zajęciach).

### Oddawanie projektów
Oddanie projektu oznacza dla mnie dostarczenie repozytorium git do mnie przed terminem oddania projektu.
Stan projektu będzie sprawdzany w wersji uzyskanej z następującego polecenia 
```
git rev-list -n 1 --first-parent --before="data-oddania-2025 23:59" master
```

## Tabela punktacji za projekty

| Numer Projektu | Opis                           | Punktacja | Termin oddania (nr laboratorium) | Odnośnik                          |
| -------------- | ------------------------------ | --------- | -------------------------------- | --------------------------------- |
| 1              | Badanie sposobów odczytu pliku | 20        | 20.10.2025r. (3)                 | [opis](projects/proj1/PROJECT.md) |
| 2              | Przygotowanie formatu pliku    | 20        | 13.11.2025r. (6)                 | [opis](projects/proj2/PROJECT.md) |
| 3              | Wczytywanie danych z pliku CSV | 30        | 08.12.2025r. (10)                 | [opis](projects/proj3/PROJECT.md) |
| 4              | Tworzenie zapytań o dane       | 30        | 19.01.2026r. (15)                | [opis](projects/proj4/PROJECT.md) |

## Wymagania techniczne

Przedmiot będzie wymagać pisania kodu na komputery z systemem Linux.
W ramach laboratorium będziemy używać funkcji ze standardu POSIX.
Język wykonania projektów jest dowolny (preferowane języki kompilowane do kodu maszynowego).
W ramach projektów 2-4 wasze programy będą implementować dostarczony interfejs REST do wykonywania poleceń użytkownika (polecam wybrać język, w którym wykonanie takiego interfejsu jest proste)

Aby ułatwić sprawdzanie, proszę w skończonym projekcie dostarczyć plik `Makefile`, który posiada target `all`.
Taki target powinien zbudować program, który będzie wykonywalny na systemie operacyjnym Linux.
W przypadku dodatkowych kroków (np. konfiguracje środowisk wykonania) proszę zawrzeć w pliku `README.md` znajdującym się w roocie projektu.

