## Компьютерные сети. Лабораторная работа III

Написать клиент-серверное приложение на основе веб сокетов и протокола TCP для передачи бинарного файла. Добавить возможность подключения нескольких клиентов одновременно и подсчета количества переданных каждым из них файлов.

<hr>

Запуск сервера:
```
cd server && go run main.go

Listening on 127.0.0.1:8081...

New connection ID: 1
---------------------------
Connection ID:      1
File name:          README.md
File size:          505 bytes
Sent files number:  1
---------------------------
```

Запуск клиента:
```
cd client && go run main.go

Set the path of file to be send: ../README.md
Send file to server:  README.md
```