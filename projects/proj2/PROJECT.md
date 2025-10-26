# Laboratorium 1 - Przygotowanie formatu pliku

Celem laboratorium jest przygotowanie własnego formatu pliku, który będzie służyć do składowania danych na dysku.

## Format pliku

Zaproponowany format powinien wspierać następujące funkcje:
* Przechowywanie dowolnej ilości kolumn
* Kolumny mają dwa dozwolone typy: 64-bitowa liczba całkowita ze znakiem oraz napis dowolnej długości (VARCHAR)
* Format pliku powinien wspierać kompresję.
* Dane w pliku powinny być tabelaryczne (tzn. każda kolumna w pliku ma taka samą długość)

## Serializator i deserializator
Zaproponowany plik powinien być wykorzystywany w tworzonym analitycznym systemie DBMS jako źródło danych do przetwarzania.
Jako początek systemu przygotuj komponent odpowiedzialny za zapis oraz odczyt przygotowanego pliku.

## Reprezentacja w pamięci
Aby móc operować na wczytanych danych (lub je zapisać do pliku) potrzebna jest reprezentacja danych w pamięci.
Przygotuj typ reprezentujący kolumnowe dane gotowe do przetwarzania.
**Przygotuj strukturę danych w taki sposób, aby procesor mógł pracować na nich najbardziej wydajnie.**

## Struktura projektu
Jest to baza do zaliczenia kolejnych projektów. Proponuję wybrać technologię, która ułatwi pracę z danymi na tak niskim poziomie oraz będzie miała do dyspozycji dobrze znaną bibliotekę do implementacji REST API.

Do zaliczenia należy przygotować program, który wczytuje plik w zadanym formacie, wykonuje deserializację oraz wyznacza dwie metryki z wczytanych danych:
* Dla każdej kolumny o typie całkowitym oblicza średnią z wartości kolumny.
* Dla każdej kolumny o typie VARCHAR wyznacza liczbę wszystkich występujących znaków ASCII.

Warto także napisać program, który wykonuje operację serializacji.

## Materiały do przeczytania
* Variable length int encoding
* Delta int encoding
* LZ4
* ZTSD
