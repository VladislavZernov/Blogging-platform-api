# Blogging-platform-api
This is a test blogging platform. The main goal is to become familiar with  technologies and education.
получение списка пользователей

curl -X GET http://localhost:8080/users


###########################################
добавить юзера
 
 curl -X POST http://localhost:8080/users \
     -H "Content-Type: application/json" \
     -d '{
           "username": "till",
           "email": "till@example.com",
           "password": "goland"
         }'

####################################################

Опубликовать пост

curl -X POST http://localhost:8080/posts \
     -H "Content-Type: application/json" \
     -d '{
           "title": "My First Blog Post",
           "content": "This is the content of my first blog post.",
           "user_id": ***(необходимо проверить, после создания пользователя)
         }'
