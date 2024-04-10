INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Usuário1", "usuario1", "usuario1@gmail.com", "$2a$10$NH9kta/XWXvyKjyUdPpA5.xJxiy76.g436BJ2Eg9J.kA.lQFoB7cq"), -- usuário 1
("Usuário2", "usuario2", "usuario2@gmail.com", "$2a$10$NH9kta/XWXvyKjyUdPpA5.xJxiy76.g436BJ2Eg9J.kA.lQFoB7cq"), -- usuário 2
("Usuário3", "usuario3", "usuario3@gmail.com", "$2a$10$NH9kta/XWXvyKjyUdPpA5.xJxiy76.g436BJ2Eg9J.kA.lQFoB7cq"); -- usuário 3

INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES
(1, 2),
(3, 1),
(1, 3);