insert into schema.users (id, name, is_admin, email, password) values
(1, 'The Administrator', true, 'admin@example.com', 'admin'),
(2, 'The User', false, 'user@example.com', 'user');

insert into schema.comments (id, id_user, txt) values
(1, 1, 'Hey from Admin!'),
(2, 2, 'Hey from User!'),
(3, 2, 'Hey from User again!');
