
1. Зайдіть у cmd/weather/config.yaml та вставте свій api key з сайту - https://openweathermap.org/
Щоб запустити проект виконайте команду `make dev`

2. Перш за все, вам потрібно зареєструватись

http://localhost:8888/signUp POST
`
{
 "first_name":"John",
 "last_name":"Doe",
 "email":"john@example.com",
 "password":"secret"
}
`

Потім залогіньтесь, отримайте jwt і робіть з ним запит для отримання погоди
http://localhost:8888/login POST
`
{
"email":"john@example.com",
"password":"secret"
}
`

3. Приклади запитів:
`http://localhost:8888/api/forecast/now?latitude=52.5200&longitude=13.4050`
`http://localhost:8888/api/forecast/now?city=London&date=1637059200`
`http://localhost:8888/api/forecast/now?city=New York`
`http://localhost:8888/api/forecast/history` - отрмання з бази даних записаних раз на годину
