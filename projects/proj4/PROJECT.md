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
    - [Ograniczenia zapytań SELECT](#ograniczenia-zapytań-select)
    - [Potok przetwarzania danych (Pipeline)](#potok-przetwarzania-danych-pipeline)
    - [External Merge Sort](#external-merge-sort)
    - [Common Subexpression Elimination (CSE)](#common-subexpression-elimination-cse)
    - [Zrównoleglenie potoku przetwarzania](#zrównoleglenie-potoku-przetwarzania)
    - [Sprawdzanie poprawności zapytania](#sprawdzanie-poprawności-zapytania)
  - [Wymagania projektowe](#wymagania-projektowe)
    - [Podział punktów](#podział-punktów)
    - [Ocena projektu](#ocena-projektu)
    - [Wymagania techniczne](#wymagania-techniczne)
    - [Sugerowane sposoby testowania](#sugerowane-sposoby-testowania)
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

Planowanie zapytania polega na przygotowaniu planu wykonania na podstawie zapytania użytkownika.
Plan wykonania jest ciągiem operacji, które należy wykonać, aby uzyskać wynik zapytania.
Każda operacja w ma zdefiniowany kształ wejścia oraz kolumny na których ma działać.
Wynik pojedynczej operacji najczęściej będzie zapisany w pewnej kolumnie (istniejącej lub dodanej w trakcie działania operacji).

### Potok przetwarzania danych (Pipeline)

Wykonanie zapytania `SELECT` można przedstawić jako potok przetwarzania danych składający się z następujących etapów:

```
Reader → Transformation → Filter → Transformation → Sort → Limit → Output
```

![pipeline](res/pipeline.svg)

Opis poszczególnych etapów:

1. **Reader** - operator odczytujący dane z plików tabeli. Produkuje strumień wierszy w formacie kolumnowym (batch).
2. **Transformation (projekcja wejściowa)** - opcjonalna transformacja przygotowująca dane do filtrowania. Oblicza wyrażenia kolumnowe potrzebne w klauzuli `WHERE`.
3. **Filter** - operator filtrujący wiersze na podstawie wyrażenia z klauzuli `WHERE`. Przepuszcza tylko wiersze, dla których wyrażenie ewaluuje się do `TRUE`.
4. **Transformation (projekcja wyjściowa)** - oblicza wyrażenia kolumnowe zdefiniowane w liście `SELECT`. Wyznacza finalny kształt danych wyjściowych.
5. **Sort** - operator sortujący dane według wskazanych kolumn. Wymaga zebrania wszystkich danych przed rozpoczęciem sortowania (operator blokujący).
6. **Limit** - operator ograniczający liczbę wierszy w wyniku. Po osiągnięciu limitu może zakończyć przetwarzanie wcześniej.


### External Merge Sort

W przypadku dużych zbiorów danych, które nie mieszczą się w pamięci operacyjnej, należy zaimplementować algorytm **External Merge Sort**:

1. **Faza podziału (Run Generation)**:
   - Wczytaj porcję danych, która mieści się w dostępnej pamięci.
   - Posortuj porcję w pamięci (np. algorytmem quicksort).
   - Zapisz posortowaną porcję do pliku tymczasowego (tzw. *run*).
   - Powtórz dla kolejnych porcji danych.

2. **Faza scalania (Merge Phase)**:
   - Otwórz wszystkie pliki tymczasowe (*runy*) jednocześnie.
   - Użyj algorytmu k-way merge z użyciem kolejki priorytetowej (heap).
   - Produkuj posortowany wynik strumieniowo.

W przypadku zbyt dużej liczby *runów* do jednoczesnego scalenia, należy wykonać wielopoziomowe scalanie (multi-level merge).

### Common Subexpression Elimination (CSE)

Optymalizacja **CSE** (eliminacja wspólnych podwyrażeń) polega na wykrywaniu powtarzających się wyrażeń kolumnowych i obliczaniu ich tylko raz.
W przyapdku naszego laboratorium chcemy wyznaczy wspólne obliczenie pomiędzy transformacjami przed filtrowaniem oraz po filtrowaniu.

Przykład:
```sql
SELECT strlen(concat(col1, col2)), strlen(concat(col1, col2)) + 10
WHERE strlen(concat(col1, col2)) > 5
```

Wyrażenie `strlen(concat(col1, col2))` występuje trzykrotnie. Po zastosowaniu CSE oblczenie `strlen(concat(col1, col2))` zapisujemy w kolumnie tymczasowej (wyznaczonej w fazie projekcji wejściowej) i uzywamy w projekcji wyjściowej.

Jako element projektu, chciałbym aby ta funkcja została zaimpelemtnowana poprzez wyznaczenie funkcji skrótu wszystkich poddrzew operacji uzyskanych z wejścia uzytkownika.
Hash operacji powinien składać sie z unikalnego hasha wynikajacego z typu operacji oraz hasha jego dzieci.
Dla operacji przemiennych (np. dodawanie) nalezy uzyc kolejnosci hashowanie dla dzieci od najmniejszej do największej.
Poddrzewa o wspolnych hashach nalezy wyznaczyc przed filtrowaniem.

### Zrównoleglenie potoku przetwarzania

Potok przetwarzania można zrównoleglić poprzez powielenie początkowej części pipeline'u na N równoległych workerów.

![pipeline](res/parallel_pipeline.svg)

#### Zasada działania

1. **Faza równoległa (N workerów)**:
   - Dane wejściowe tabeli są podzielone na N partycji (np. po plikach lub zakresach wierszy).
   - Każdy worker niezależnie wykonuje pełny potok: Reader → Transformation → Filter → Transformation.
   - Workerzy działają równolegle, przetwarzając swoje partycje danych.
   - Każdy worker produkuje strumień przefiltrowanych i przekształconych wierszy.

2. **Faza zbierania (Sort)**:
   - Operator `Sort` jest pojedynczy i zbiera dane ze wszystkich N workerów.
   - W miarę napływu danych operator Sort "puchnie" - gromadzi wszystkie wiersze przed rozpoczęciem sortowania.
   - Jest to punkt synchronizacji - Sort musi poczekać na zakończenie pracy wszystkich workerów.

3. **Faza końcowa (Limit)**:
   - Po posortowaniu danych operator `Limit` jest również pojedynczy.
   - Ogranicza liczbę wierszy w końcowym wyniku.

### Sprawdzanie poprawności zapytania

Planista jest odpowiedzialny za walidację zapytania przed wykonaniem:

1. **Walidacja referencji do kolumn** - wszystkie odwołania do kolumn muszą istnieć w tabeli źródłowej.
2. **Walidacja typów** - operatory i funkcje muszą otrzymywać argumenty zgodnych typów.
3. **Walidacja wyrażenia filtrującego** - musi ewaluować się do typu `BOOL`.
4. **Walidacja indeksów sortowania** - muszą odnosić się do istniejących kolumn wyjściowych.
5. **Pojedyncza tabela źródłowa** - zapytanie może pobierać dane tylko z jednej tabeli. Operacje `JOIN` nie są wspierane.

W przypadku błędu walidacji należy zwrócić błąd planowania.

## Wymagania projektowe

### Podział punktów

Projekt jest oceniany w skali **30 punktów**, podzielonych na następujące części:

| Komponent | Punkty | Opis |
|-----------|--------|------|
| **External Merge Sort** | 10 pkt | Implementacja algorytmu sortowania zewnętrznego dla danych przekraczających rozmiar pamięci |
| **Common Subexpression Elimination** | 10 pkt | Optymalizacja eliminacji wspólnych podwyrażeń w planie zapytania |
| **Planowanie i wykonanie zapytań** | 10 pkt | Przygotowanie planu wykonania, walidacja zapytania oraz implementacja potoku przetwarzania |

### Ocena projektu

Ocena projektu zostanie wyznaczona na podstawie **testów PIT (Public Interface Testing)**.
Testy te weryfikują poprawność implementacji poprzez publiczny interfejs REST API bazy danych.

Testy PIT sprawdzają:
- Poprawność wyników zapytań `SELECT` z różnymi wyrażeniami kolumnowymi.
- Działanie filtrowania, sortowania i limitu.
- Poprawność sortowania dla dużych zbiorów danych (External Merge Sort).

**Ważne**: Upewnij się, że Twoja implementacja jest w pełni zgodna z interfejsem zdefiniowanym w pliku `dbmsInterface.yaml`.

### Sugerowane sposoby testowania

#### Testy External Merge Sort
- Przygotuj dane testowe większe niż dostępna pamięć (polecam parametr uruchomienia bazy danych definiujący dostępną pamięć M).
- Zweryfikuj poprawność sortowania dla różnych typów danych (`INT64`, `VARCHAR`, `BOOL`).
- Sprawdź sortowanie po wielu kolumnach.

#### Testy CSE
- Wygeneruj format wizualizacji planu. Prosta operacja wypisywania na standardowe wyjście będzie wystarczająca.
- Napisz testy jednostkowe potwierdzające, ze wspolna operacja zostala wykryta i plan zawiera tylko jedna instancje obliczen.

#### Testy integracyjne
- Wykonaj kompletne zapytania `SELECT` z filtrowaniem, transformacjami, sortowaniem i limitem.
- Porównaj wyniki z oczekiwanymi danymi.

## Materiały do przeczytania
* [Algebra relacji](https://en.wikipedia.org/wiki/Relational_algebra)
* [K-way merge sort]()
* 
