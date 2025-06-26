create table video {
    id int primary key
    titile text
    duration int
    author_id int fk
}

create table author {
    id int primary key
    name text
}

/* написать запрос, который покажет авторов, у которых нет видео с duration > 20 */

SELECT name
LEFT JOIN video as v ON v.author_id = a.id AND v.duration > 20
FROM author as a
WHERE v.id IS NULL

/*
SELECT name
FROM author a
WHERE NOT EXISTS (
  SELECT 1
  FROM video v
  WHERE v.author_id = a.id AND v.duration > 20
);
*/