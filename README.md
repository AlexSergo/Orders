# Orders
В БД: 
Развернуть локально postgresql
Создать свою бд
Настроить своего пользователя. 
Создать таблицы для хранения полученных данных.
В сервисе:
1. Подключение и подписка на канал в nats-streaming
2. Полученные данные писать в Postgres
3. Так же полученные данные сохранить in memory в сервисе (Кеш)
4. В случае падения сервиса восстанавливать Кеш из Postgres
5. Поднять http сервер и выдавать данные по id из кеша
6. Сделать простейший интерфейс отображения полученных данных, для
их запроса по id
