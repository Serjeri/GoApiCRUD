Убедитесь, что у вас установлены:
-Docker
-Docker Compose

скачайте проект спомощью 
git clone https://github.com/Serjeri/GoApiCRUD.git

Запустите проект:
 docker-compose up --build 
 для остановки проекта нажмите ctrl + c
 
 передите по ссылке 0.0.0.0:5001/api/ если у вас есть надпись  "message": "Tasks retrieved successfully"
 то все ок и работает 

Вариант 2: Вручную (без Docker)
скачайте проект спомощью 
git clone https://github.com/Serjeri/GoApiCRUD.git

Настройте .env файл

Запустите сервер:
go run main.go

При первом запуске Docker Compose автоматически:
 -Создаст том для PostgreSQL
 -Инициализирует БД через init.sql
 -Запустит миграции

примеры для отправки запросов:
curl --location --request GET '0.0.0.0:5001/api/' \
--header 'Content-Type: application/json'

ответ от сервера:
![workingServer](https://github.com/user-attachments/assets/19fc244b-7147-418d-a5b4-8782da8ea901)

Создание Такси:
curl --location '0.0.0.0:5001/api/create' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Создание Такси",
    "description": "Утренняя зарядка 15 минут"
}'
ответ от сервера 
![create](https://github.com/user-attachments/assets/b6cf6eb6-8269-4803-a55b-7e0a657a2e86)

Получение всех Тасок
curl --location '0.0.0.0:5001/api/getalltasks' \
--data ''
ответ от сервера:
![getalltask](https://github.com/user-attachments/assets/9128575c-e9f9-4296-ac24-050e095097a1)

Обновление Таски 
curl --location --request PUT '0.0.0.0:5001/api/update/1' \
--header 'Content-Type: application/json' \
--data '{
    "title": "test",
    "description": " 15 минут"
}'

ответ от сервера 
![updateTask](https://github.com/user-attachments/assets/00b8064d-5da3-4f0c-92c0-12638892dbfc)

удаление таски 
curl --location --request DELETE '0.0.0.0:5001/api/delete/1' \
--header 'Content-Type: application/json' \
--data '{
    "title": "test",
    "description": " 15 минут"
}'
ответ от сервера
![DeleteTask](https://github.com/user-attachments/assets/74dc1f1e-387c-4159-ad55-973a4e4fa944)


