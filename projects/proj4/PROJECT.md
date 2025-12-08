# Laboratorium 4 - Tworzenie zapytań o dane

- [Laboratorium 4 - Tworzenie zapytań o dane](#laboratorium-4---tworzenie-zapytań-o-dane)
  - [Klauzula SELECT](#klauzula-select)
    - [Wyrażenia kolumnowe](#wyrażenia-kolumnowe)
      - [Operatory](#operatory)
      - [Funkcje](#funkcje)
      - [Przykłady wyrażeń kolumnowych](#przykłady-wyrażeń-kolumnowych)
    - [Filtrowanie danych - klauzula `WHERE`](#filtrowanie-danych---klauzula-where)
    - [Sortowanie danych](#sortowanie-danych)
    - [Ograniczenie wierszy wynikowych](#ograniczenie-wierszy-wynikowych)
  - [Planowanie i wykonanie zapytania](#planowanie-i-wykonanie-zapytania)
  - [Materiały do przeczytania](#materiały-do-przeczytania)


Celem laboratorium nr 4 jest przygotowanie systemu zapytań do wcześniej wstawionych danych.
System udostępni elastyczny język zapytań, który definiuje potrzebne dane bez informacji jak te dane uzyskać.

## Klauzula SELECT

W tym projekcie należy zaimplementować rozszerzoną klauzulę `SELECT`, która od teraz będzie umożliwiać wybór poszczególnych kolumn oraz wykonywanie na nich wiele operacji.
Od teraz na operację `SELECT` składają się następujące elementy:
1. Lista *wyrażeń kolumnowych*, które reprezentują wynikowy kształt zapytania.
2. Wyrażenie filtrujące. Służy do ograniczenia liczby wierszy potrzebnych w wyniku zapytania.
3. Lista indeksów kolumn wyjściowych (odniesienie do kolumn z punktu pierwszego), po których należy posortować wynik zapytania.
4. Ograniczenie liczby wierszy w wyniku zapytania.

### Wyrażenia kolumnowe

Wyrażenie kolumnowe jest najogólniejszym sposobem opisania wyniku zapytania przez użytkownika.
Składa się ono z podtypów:
1. Odniesienie do kolumny - jest do para `(nazwa tabeli, nazwa kolumny)`, która jednoznacznie wyznacza kolumnę z danymi. 
2. Literał - dane znanego typu wpisane przez użytkownika wprost (np. `"Ala ma kota"` albo `1234`).
3. Funkcja - dobrze zdefiniowana operacja, która przyjmuje pewną ilość argumentów, z których wyznacza wartości dobrze określonego typu (np. `concat("Hello", concat("World", "!"))` - funkcja może jako argument przyjąć dowolne inne wyrażenie kolumnowe).
4. Operator binarny - przyjmuje dwa wyrażenia kolumnowe i wykonuje na nich pewną fundamentalną operację. Argumenty oraz wynik operacją posiadają typy (np. "col1 + 2" - ma sens tylko, jeśli kolumna `col1` jest typu `INT64`; wtedy wynikiem jest wartość typu `INT64`).
5. Operator unarny - podobnie jak w przypadku operatora binarnego, tylko liczba argumentów jest równa 1.

#### Operatory

Do zaimplementowania są następujące operatory:
1. Dodawanie - przyjmuje dwie liczby całkowity i zwraca liczbę całkowitą
2. Odejmowanie - przyjmuje dwie liczby całkowity i zwraca liczbę całkowitą
3. Mnożenie - przyjmuje dwie liczby całkowity i zwraca liczbę całkowitą
4. Dzielenie - przyjmuje dwie liczby całkowity i zwraca liczbę całkowitą
5. Logiczna koniunkcja (AND) - przyjmuje dwie wartości logiczne i zwraca wartość logiczną
6. Logiczna alternatywa (OR) - przyjmuje dwie wartości logiczne i zwraca wartość logiczną
7. Operator równości - przyjmuje dwie wartości dowolnego typu (dwie muszą być tego samego typu) i zwraca wartość logiczną. W przypadku liczb całkowitych i wartości logicznych porównuje ich wartości, a w przypadku napisów porównania następuje znak po znaku.
8.  Operator nierówności - przyjmuje dwie wartości dowolnego typu (dwie muszą być tego samego typu) i zwraca wartość logiczną. W przypadku liczb całkowitych i wartości logicznych porównuje ich wartości, a w przypadku napisów porównania następuje znak po znaku.
9.  Operator mniejszości - przyjmuje dwie wartości dowolnego typu (dwie muszą być tego samego typu) i zwraca wartość logiczną. W przypadku liczb całkowitych i wartości logicznych porównuje ich wartości, a w przypadku napisów porównania następuje w porządku leksykograficznym.
10. Operator większości - przyjmuje dwie wartości dowolnego typu (dwie muszą być tego samego typu) i zwraca wartość logiczną. W przypadku liczb całkowitych i wartości logicznych porównuje ich wartości, a w przypadku napisów porównania następuje w porządku leksykograficznym.
11. Operator mniejsze lub równe - przyjmuje dwie wartości dowolnego typu (dwie muszą być tego samego typu) i zwraca wartość logiczną. W przypadku liczb całkowitych i wartości logicznych porównuje ich wartości, a w przypadku napisów porównania następuje w porządku leksykograficznym.
12. Operator większe lub równe - przyjmuje dwie wartości dowolnego typu (dwie muszą być tego samego typu) i zwraca wartość logiczną. W przypadku liczb całkowitych i wartości logicznych porównuje ich wartości, a w przypadku napisów porównania następuje w porządku leksykograficznym.
13. Operator negacji logicznej - przyjmuje jedną wartość logiczną i zwraca wartość logiczna o odwrotnej wartości niż wejście.
14. Operator negacji - przyjmuje liczbę całkowitą i wyznacza jej wartość ujemną.

#### Funkcje

Do zaimplementowania są następujące funkcje:
1. `STRLEN` - przyjmuje jeden argument typu `VARCHAR` i zwraca liczbę całkowitą oznaczającą ilość znaków w napisie (zakładamy tylko znaki ASCII).
2. `CONCAT` - przyjmuje dwa argumenty typu `VARCHAR` i zwraca wartość typu `VARCHAR`, która jest połączeniem dwóch wejściowych napisów.
3. `UPPER` - przyjmuje jeden argument typu `VARCHAR` i zwraca napis z przekształconymi literami małymi w wielkie (zakładamy tylko znaki ASCII).
4. `LOWER` - przyjmuje jeden argument typu `VARCHAR` i zwraca napis z przekształconymi literami wielkimi w małe (zakładamy tylko znaki ASCII).

#### Przykłady wyrażeń kolumnowych

Poniżej można znaleźć przykłady wyrażeń kolumnowych wyrażonych w języku zrozumiałym dla człowieka.
Interfejs bazy danych definiuje struktury opisujące te wyrażenia bardzo precyzyjnie.

Nie wszystkie wyrażenia kolumnowe można wykonać, jednak można je wprowadzić poprzez interfejs.
Odpowiedzialnością planisty jest ocena czy wyrażenie jest wykonywalne.
Przykłady wyrażeń kolumnowych:
* `1 + 2 + 3`
* `"Ala ma kota" / 2` - to wyrażenie kolumnowe nie jest wykonywalne
* `strlen(concat(col1, col2))` - obliczanie długości z konkatenacji wartości z kolumn `col1` i `col2`. Będzie wykonywalne tylko w przypadku gdy kolumny są typu `VARCHAR`.

### Filtrowanie danych - klauzula `WHERE`

Wyrażenie filtrujące ma za zadanie wybrać wiersze do dalszego przetwarzania.
Podane wyrażenie kolumnowe powinno ewaluować się do typu bool.
Jeśli tak nie jest, należy zwrócić błąd planowania.

Jeśli dla danego wiersza wyrażenie ewaluuje się do wartości `TRUE`, to dane powinny być częścią wyniku.
W przeciwnym wypadku dane można odrzucić.

Przykłady:
* `col1 > col2` 
* `strlen(col1) < 16`
* `-col1 >= 0`

### Sortowanie danych

Sortowanie danych może odbywać się tylko na danych wynikowych.
Jest to bardzo przydatny mechanizm do pobierania danych z klauzulą `LIMIT`.
Bez użycia sortowania użytkownik nie może nic założyć o kolejności zwracanych danych.

Sortowanie w przypadku liczba całkowitych i wartości logicznych odbywa się względem ich wartości.
Dla wartości typu `VARCHAR` należy użyć porządku leksykograficznego.

### Ograniczenie wierszy wynikowych

Jeśli użytkownik poda wartość tego elementu, to przetwarzanie można zakończyć po uzyskaniu N wierszy, gdzie N to uzyskany parametr.
Po posortowaniu wartości oczekiwaną wartością jest N najmniejszych lub największych wartości względem podanej miary.

## Planowanie i wykonanie zapytania

TBD
* external merge sort
* kolejnosc przetwarzania danych
* sprawdzanie poprawnosci zapytania

## Materiały do przeczytania
* [Algebra relacji](https://en.wikipedia.org/wiki/Relational_algebra)
