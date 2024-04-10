INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Usuário1", "usuario1", "usuario1@gmail.com", "$2a$10$NH9kta/XWXvyKjyUdPpA5.xJxiy76.g436BJ2Eg9J.kA.lQFoB7cq"), -- usuário 1
("Usuário2", "usuario2", "usuario2@gmail.com", "$2a$10$NH9kta/XWXvyKjyUdPpA5.xJxiy76.g436BJ2Eg9J.kA.lQFoB7cq"), -- usuário 2
("Usuário3", "usuario3", "usuario3@gmail.com", "$2a$10$NH9kta/XWXvyKjyUdPpA5.xJxiy76.g436BJ2Eg9J.kA.lQFoB7cq"), -- usuário 3
("Luis", "luis", "luis@gmail.com", "$2a$10$NH9kta/XWXvyKjyUdPpA5.xJxiy76.g436BJ2Eg9J.kA.lQFoB7cq"),             -- usuário 4
("Felipe", "felipe", "felipe@gmail.com", "$2a$10$NH9kta/XWXvyKjyUdPpA5.xJxiy76.g436BJ2Eg9J.kA.lQFoB7cq"),       -- usuário 5
("Boeff", "boeff", "boeff@gmail.com", "$2a$10$NH9kta/XWXvyKjyUdPpA5.xJxiy76.g436BJ2Eg9J.kA.lQFoB7cq"),          -- usuário 6
("Teste", "teste", "teste@gmail.com", "$2a$10$NH9kta/XWXvyKjyUdPpA5.xJxiy76.g436BJ2Eg9J.kA.lQFoB7cq");          -- usuário 7

INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES
(1, 2),
(3, 1),
(1, 3),
(6, 4),
(6, 5),
(5, 4),
(5, 6);

INSERT INTO publicacoes (titulo, conteudo, autor_id)
VALUES
("Minha Primeira Publicação", "Hey", 7),
("Minha Segunda Publicação", "Natália!", 7),
("Minha Terceira Publicação", "World!", 7),
("Aloha", "Galeraaaaa!", 7),
("Publicação do Usuário 1", "Aloha", 1),
("Publicação do Usuário 2", "Ihul", 2),
("Publicação do Usuário 3", "Oba", 3);




eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTI3ODYxMjksInVzdWFyaW9JZCI6N30.MCU0mfN8osqkPcL3hWrZb7LWNYavhVZmVVvnFR-BhNg

{
    "email": "teste@gmail.com",
    "senha": "boeff"
}