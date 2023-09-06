insert into schema.users (name, is_admin, email, password) values
('The Administrator', true, 'admin@example.com', 'admin'),
('The User', false, 'user@example.com', 'user');

insert into schema.comments (id_user, txt) values
(1, 'Hey from Admin!'),
(2, 'Hey from User!'),
(2, 'Hey from User again!');
