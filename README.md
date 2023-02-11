# cities-informer

Последнее задание в блоке GO-разработчик от skillbox.

*Далее представлено задание*

<details>
<summary>Описание промежуточной аттестации</summary>
Цель работы

Проверить и закрепить знания, полученные на курсе «Go-разработчик»:

- основы синтаксиса языка;
- условные операторы и циклы;
- работа с файловой системой;
- структуры данных;
- сериализация;
- многопоточность;
- обмен данными по сети.

Что нужно сделать
Вам нужно разработать сервис, предоставляющий информацию о городах. Данные хранятся в файле. В момент старта сервиса данные из файла кешируются в память, в момент завершения работы сервиса данные перезаписываются обратно в файл.

В каждой строке файла через запятую перечислена информация о городе:
- id (уникальный номер);
- name (название города);
- region (регион);
- district (округ);
- population (численность населения);
- foundation (год основания).


Требуется реализовать сервис имеющий следующий функционал:
- [X] получение информации о городе по его id;
- [X] добавление новой записи в список городов;
- [X] удаление информации о городе по указанному id;
- [X] обновление информации о численности населения города по указанному id;
- [X] получение списка городов по указанному региону;
- [X] получение списка городов по указанному округу;
- [X] получения списка городов по указанному диапазону численности населения;
- [X] получения списка городов по указанному диапазону года основания;
- [ ] api для получения данных.

</details>