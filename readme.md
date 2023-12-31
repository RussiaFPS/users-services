# users-services
# Краткое описание 
Основные данные из *docker-compose.yml*
 - Порт для API: 8080
 - Пользователь для БД: postgres
 - Пароль для БД: qwer1234
 - Порт для БД: 5430 

Основные данные для подключения к БД прописанны в env файле, необходима 
база данных __users-services__

# Endpoints

- Добавление пользователя
```
    [POST] http://localhost:8080/api/add
```
 Обязательные поля имя и фамилия
 ```
 {
    "name": "Dmitriy",
    "surname": "Ushakov",
    "patronymic": "Vasilevich" // необязательно
}
 ```

- Получение всех пользователей
```
    [GET] http://localhost:8080/api/user
```
- Так же можно получить данные по страницам
```
    [GET] http://localhost:8080/api/user?page=2
```
- Получение пользователя по id
```
    [GET] http://localhost:8080/api/user/2
```
- Обновить пользователя по id
```
    [PUT] http://localhost:8080/api/update/19
```
Отпрака новых данных, так же обязательные поля имя и фамилия
 ```
 {
    "name": "Dmitriy",
    "surname": "Ushakov",
    "patronymic": "Vasilevich" // необязательно
}
 ```
- Удаление пользователя по id
```
    [DELETE] http://localhost:8080/api/user/7
```
### Реализовано
✅ 4 rest метода, с пагинацией и фильтрами

✅ Входные сообщения обогащенны данными

✅ Поднята БД Postgres, для таблицы users настроена автомиграция

✅ Код покрыт debug и info логами

✅ Конфигурационные данные хранятся в env файле
