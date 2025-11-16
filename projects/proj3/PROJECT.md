# Laboratorium 3 - Wczytywanie danych z pliku CSV

- [Laboratorium 3 - Wczytywanie danych z pliku CSV](#laboratorium-3---wczytywanie-danych-z-pliku-csv)
  - [Metastore -- logiczna struktura bazy danych](#metastore----logiczna-struktura-bazy-danych)
  - [Interfejs użytkownika](#interfejs-użytkownika)
    - [Operacje na tabelach](#operacje-na-tabelach)
    - [Wykonywanie zapytań](#wykonywanie-zapytań)
      - [Zapytanie COPY](#zapytanie-copy)
      - [Zapytanie SELECT](#zapytanie-select)
  - [Architektura aplikacji](#architektura-aplikacji)
  - [Wymagania techniczne](#wymagania-techniczne)
  - [Sugerowane sposoby testowania](#sugerowane-sposoby-testowania)
    - [Public interface testing (PIT)](#public-interface-testing-pit)
    - [Testy poszczególnych modułów](#testy-poszczególnych-modułów)
  - [Materiały do przeczytania](#materiały-do-przeczytania)


Celem laboratorium jest przygotowanie działającego systemu bazodanowego, który potrafi wczytać plik CSV i zapisać go do tabeli.
Zapis danych przedstawiających tabelę jest już zaimplementowany w projekcie nr 2.

## Metastore -- logiczna struktura bazy danych

W poprzednim projekcie przygotowana została warstwa składowania danych (potocznie zwana fizyczną).
Składa się na nią format pliku, reprezentacja danych w pamięci oraz operacja serializacji/deserializacji pomiędzy nimi.
Na tak sformatowanych danych wykonywane będą operacje analityczne.

Interfejs użytkownika w bazach danych składa się z tabel i kolumn (potocznie zwane warstwą logiczną).
**Nie ma tutaj mowy o plikach**.
Potrzeba jest więc struktura danych, która przetłumaczy abstrakty bazodanowe (tabele, widoki, indexy itp.) z warstwy logicznej na warstwę składowania danych.
Takim miejscem jest **metastore**.

Jest to struktura danych, która dostarcza wymaganych informacji do zaplanowania wykonania zapytania.
Jednym z głównych zadań planowania jest wyznaczenie listy plików do odczytania, aby uzyskać pełen zestaw danych oczekiwanych przez użytkownika.
Robione jest to na podstawie translacji logicznych nazw tabel i kolumn na odpowiednie pliki przechowujące dane.

**Wszelkie struktury danych budujące metastore powinny mieć trwałość większą niż system bazodanowy, tzn. powinny zostać zachowane pomiędzy uruchomieniami bazy danych.**
Najczęściej oznacza to, że metastore także musi zostać zapisany na dysk.

## Interfejs użytkownika

W [pliku](../../resources/dbmsInterface.yaml) można znaleźć interfejs bazy danych do wykonania w tym projekcie.
Jest to minimalny zestaw poleceń w twoim projecie, który umożliwi sprawdzanie projektów.

Operacje podzielone są na dwa typy: `schema` oraz `execution`.
Poniżej można znaleźć opis operacji znajdujących się w poszczególnych grupach.

### Operacje na tabelach

Głównym endpointem tej grupy jest `/tables`, który zwraca listę wszystkich tabel w systemie.
Na podstawie odczytanych danych można odczytać szczegóły tabeli odpytując endpoint `/table/{tableId}`.
Tam można znaleźć strukturę całej tabeli.

Dodatkowo istnieją dwa endpointy służące do tworzenia oraz usuwania tabel z systemu.

Tworzenie tabeli odbywa się poprzez wysłanie zapytania `PUT` do endpointu `/table`, w którym zostanie opisana cała struktura tabeli.
Jeżeli taką tabelę można utworzyć (ma ona unikatową nazwę oraz posiada unikatowe nazwy kolumn), zapytanie powinno zakończyć się sukcesem. Wpw. użytkownik powinien otrzymać błąd opisujący występujący problem.

Usuwanie tabeli odbywa się poprzez wysłanie zapytania `DELETE` do endpointu `/table/{tableId}`.
Jeśli tabela istnieje, powinna zostać usunięta z rejestru tabel.
Dodatkowo jej pliki będą usunięte w pewnym momencie po wykonaniu akcji.

**Obydwie operacje powinny mieć obserwowalne efekty uboczne dla zapytań następujących po wykonaniu akcji!.
Oznacza to, że aktualnie trwające zapytania nie powinny utracić dostępu do plików właśnie usuwanej tabeli.**

### Wykonywanie zapytań

Przez to, że interfejs został zaprojektowany w duchu REST, nie jest on stanowy.
Jednak nasz system ewidentnie posiada stan, ponieważ zapytania mogą trwać *pewien* czas od ich zlecenia do zakończenia.

Z tego powodu głównym endpointem jest `/queries`, który dostarcza podstawowe informacje o wykonanych zapytaniach do systemu.
Na podstawie odczytanych danych można odczytać szczegóły zapytań odpytując endpoint `/query/{queryId}`.

System przewiduje dwa rodzaje zapytań: `COPY` oraz `SELECT`.

#### Zapytanie COPY

Operacja `COPY` służy do wstawiania danych z pliku w formacie CSV.
Posiada ona dwa obowiązkowe parametry: ścieżkę do wczytywanego pliku CSV oraz nazwę docelowej tabeli.

Poza tym są tam dwa dodatkowe parametry: flaga informująca, czy CSV posiada nagłówek oraz mapowanie kolumn CSV do kolumn w tabeli.
Mapowanie jest konieczne w sytuacji wstawiania danych, gdzie tabela posiada mniej kolumn niż plik CSV.
Umożliwia ono odpowiednie przekierowanie nadwyżki danych do wcześniej przygotowanej tabeli.

**Tak jak w przypadku operacji usuwania tabeli, efekt uboczny wstawiania danych ma być widoczny dopiero po całkowitym zakończeniu wstawiania danych. Niedopuszczalne są odczyty stanu pośredniego tabeli.**
Najprościej można wykonać taką operację poprzez tworzenie nowego pliku przy każdej operacji COPY i dodanie tych plików do metastore dopiero po zakończeniu całej operacji.

#### Zapytanie SELECT

Operacja `SELECT` przyjmuje nazwę tabeli i ma za zadanie pobrać wszystkie wiersze z tabeli (odpowiednik operacji `SELECT * FROM tableName;` w bazie SQL).
Jej celem jest obserwacja wstawionych danych poprzez interfejs użytkownika.
**Kolejność wierszy w zwróconych danych nie jest ustalona.**

W przypadku tego zapytania istnieje specjalny endpoint `/results/{queryId}`, który służy do pobierania wyniku po zakończeniu wykonania zapytania.
Istnieje możliwość ograniczenia zwróconych wierszy poprzez opcjonalny parametr w ciele zapytania.

## Architektura aplikacji



## Wymagania techniczne

Baza danych powinna nasłuchiwać na sockecie TCP wiadomości z protokołu HTTP.
Zamknięcie aplikacji nie powinno powodować utarty lub korupcji danych.

## Sugerowane sposoby testowania

### Public interface testing (PIT)

### Testy poszczególnych modułów

## Materiały do przeczytania