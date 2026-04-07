
**Zadanie 1** Docker

:white_check_mark: 3.0 obraz ubuntu z Pythonem w wersji 3.10 </br> [Link do commita ](https://github.com/Szeliah/E-Biznes/commit/d6472b8ce5eead4465d9c7d2d884041f382066e3) </br>
[link do obrazu](https://hub.docker.com/layers/szeligaszymon/e-biznes/task_3.0/images/sha256-7d764b279270236dd828d0a26492f199fe716b6b67ee4d7d8b5ba3e052272751)

:white_check_mark: 3.5 obraz ubuntu:24.02 z Javą w wersji 8 oraz Kotlinema </br>  [Link do commita ](https://github.com/Szeliah/E-Biznes/commit/8646d75d88cf6a5c970198de647a1600dc9b2a04) </br>
[link do obrazu](https://hub.docker.com/layers/szeligaszymon/e-biznes/task_3.5/images/sha256-35065d04cfe874201d282979b9a410b671715f9227dd54ef771d243e972bc0a2)

:white_check_mark: 4.0 do powyższego należy dodać najnowszego Gradle’a oraz paczkę JDBC
SQLite w ramach projektu na Gradle (build.gradle) </br>  [Link do commita ](https://github.com/Szeliah/E-Biznes/commit/cceffe07fe295022139b01765a01b0b2021ff19f) </br>
[link do obrazu](https://hub.docker.com/layers/szeligaszymon/e-biznes/task_4.0/images/sha256-8fbf66541adb6c71f093c2513c3043d0596a82919a1b806318831e0a9c21d804)

:white_check_mark: 4.5 stworzyć przykład typu HelloWorld oraz uruchomienie aplikacji
przez CMD oraz gradle </br>
[Link do commita ](https://github.com/Szeliah/E-Biznes/commit/3bcec9bfb0bb55846c2b664db8a74f45c130f6d6) </br>
[link do obrazu](https://hub.docker.com/layers/szeligaszymon/e-biznes/task_4.5/images/sha256-3fa9c8bc11dcf92e0b3bb46ea58e08f8b15566d7ddd703c66ec5b1ac6d0597c3)

:x: 5.0 dodać konfigurację docker-compose </br> [Link do commita ]() </br>
[link do obrazu]()


Kod: [Link do zadania 1](https://github.com/Szeliah/E-Biznes/tree/main/Zadanie01)

---

**Zadanie 2** Scala

Należy stworzyć aplikację na frameworku Play lub Scalatra.

:white_check_mark: 3.0 Należy stworzyć kontroler do Produktów </br> [Link do commita](https://github.com/Szeliah/E-Biznes/commit/2097f9442864727c7c03af718613e60a07f59599) </br>

:white_check_mark: 3.5 Do kontrolera należy stworzyć endpointy zgodnie z CRUD - dane
pobierane z listy </br>  [Link do commita ](https://github.com/Szeliah/E-Biznes/commit/2ae2e80baedc1397ba4b064dd784fb71dfaba576) </br>

:white_check_mark: 4.0 do Należy stworzyć kontrolery do Kategorii oraz Koszyka + endpointy
zgodnie z CRUD </br> 
[Link do commita ](https://github.com/Szeliah/E-Biznes/commit/a824704da044189de4d0d64027f9ef0f960b7648) </br>

:x: 4.5 Należy aplikację uruchomić na dockerze (stworzyć obraz) oraz dodać
skrypt uruchamiający aplikację via ngrok </br> 

:x: 5.0 Należy dodać konfigurację CORS dla dwóch hostów dla metod CRUD </br> 

Kod: [Link do zadania 2](https://github.com/Szeliah/E-Biznes/tree/main/Zadanie02/App/zadanie02app) </br> 
Demo: [link do nagrania](https://github.com/Szeliah/E-Biznes/blob/main/Assets/Zadanie02-demo.mp4)

---

**Zadanie 3** Kotlin

:white_check_mark: 3.0 Należy stworzyć aplikację kliencką w Kotlinie we frameworku Ktor,
która pozwala na przesyłanie wiadomości na platformę Discord </br> [Link do commita](https://github.com/Szeliah/E-Biznes/commit/9d743a0a964029fcd8d501d3e2753a52e10db422)

:white_check_mark: 3.5 Aplikacja jest w stanie odbierać wiadomości użytkowników z
platformy Discord skierowane do aplikacji (bota) </br>  [Link do commita ](https://github.com/Szeliah/E-Biznes/commit/a65c18f47cb3bd15adeb5521bdfdb8ed514fe61d)

:x: 4.0 Zwróci listę kategorii na określone żądanie użytkownika</br> 


:x: 4.5 Zwróci listę produktów wg żądanej kategoriik </br> 


:x: 5.0 Aplikacja obsłuży dodatkowo jedną z platform: Slack lub Messenger </br> 

</br></br>
Aplikację należy uruchomić na dockerze

Kod: [Link do zadania 3](https://github.com/Szeliah/E-Biznes/tree/main/Zadanie03/App)
Demo: [link do nagrania](https://github.com/Szeliah/E-Biznes/blob/main/Assets/Zadanie03-demo.mp4)