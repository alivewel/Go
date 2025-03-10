--  База данных: PostgreSQL

CREATE TABLE users (
    id serial,
    username varchar(255),
    email varchar(255),
    is_active boolean,
    created_at timestamp
);

# SELECT * FROM users;
 id |   username    | email | is_active | created_at 
----+---------------+-------+-----------+------------
  1 | первый_юзер   |       |           | 
  2 | второй вариант|       |           | 
  3 | третий столб  |       |           | 
  4 | четвертый путь|       |           | 
(4 rows)

CREATE TABLE profiles (
    id serial,
    user_id integer,
    first_name varchar(255),
    last_name varchar(255),
    gender integer, -- значения: 0, 1, 2
    birthdate date
);

# SELECT * FROM profiles;
 id | user_id | first_name | last_name  | birthdate | gender 
----+---------+------------+------------+-----------+--------
  1 |       1 | name1      | last_name1 |           |      0
  2 |       1 | name2      | last_name2 |           |      1
  3 |       2 | name3      | last_name3 |           |      2
(3 rows)


-- Написать два запроса:
-- Запрос 1: Вывести все id пользователей у которых нет профилей
-- Запрос 2: Для пользователей у которых более одного профиля вывести id пользователя и id профиля

-- Запрос 1:

SELECT u.id
FROM users u
LEFT JOIN profiles p on p.user_id = u.id
WHERE p.user_id is NULL;

-- Запрос 2:

1 способ:

SELECT u.id, p.user_id
FROM users u
JOIN profiles p on p.user_id = u.id
WHERE u.id IN (
    SELECT u.id
    FROM users u
    JOIN profiles p on p.user_id = u.id
    GROUP BY u.id
    HAVING COUNT(p.id) > 1
);

1) Используем JOIN, а не LEFT JOIN потому что нам требуется вернуть users с существующими профилями profiles (без NULL)
2) В WHERE используем IN, а не = , чтобы соответствовать нескольким идентификаторам пользователей, возвращаемым подзапросом. Хотя = тоже работает.
3) Порядок операторов: сначала GROUP BY, потом HAVING.
4) Используем COUNT(p.id), а не COUNT(*).

2 способ:

SELECT p.user_id, p.id as p_id
FROM profiles p
JOIN (
    SELECT user_id
    FROM profiles
    GROUP BY user_id
    HAVING COUNT(*) > 1
) subquery ON p.user_id = subquery.user_id;

1) Здесь для JOIN используем подзапрос.
