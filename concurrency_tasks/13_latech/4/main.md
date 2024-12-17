CREATE TABLE users (username text, password text, age integer)

CREATE INDEX h_idx ON users USING hash(age);

SELECT *
FROM users
WHERE age > 18;