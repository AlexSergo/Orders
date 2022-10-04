# Orders
В БД: 
Развернуть локально postgresql
Создать свою бд
Настроить своего пользователя. 
Создать таблицы для хранения полученных данных.
В сервисе:
:white_check_mark: Подключение и подписка на канал в nats-streaming
:white_check_mark: Полученные данные писать в Postgres
:white_check_mark: Так же полученные данные сохранить in memory в сервисе (Кеш)
:white_check_mark: В случае падения сервиса восстанавливать Кеш из Postgres
:white_check_mark: Поднять http сервер и выдавать данные по id из кеша
:white_check_mark: Сделать простейший интерфейс отображения полученных данных, для
их запроса по id

![video](https://raw.githubusercontent.com/AlexSergo/Orders/main/video.gif)
