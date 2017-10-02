CREATE TABLE `students` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(300) NOT NULL,
  `score` int NOT NULL
);

INSERT INTO `students` (`name`, `score`) 
VALUES ('Oleh Hudenko', '100');
VALUES ('Oleh1 Hudenko1', '99');
VALUES ('Oleh2 Hudenko2', '98');
VALUES ('Oleh3 Hudenko3', '97');
