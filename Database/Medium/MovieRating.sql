/* Write your T-SQL query statement below */
SELECT (
    SELECT TOP 1 u.name AS results
    FROM users u
        INNER JOIN movierating mr ON u.user_id = mr.user_id
    GROUP BY u.name
    ORDER BY COUNT(mr.rating) DESC, u.name
) AS results
UNION ALL
SELECT (
    SELECT TOP 1 m.title AS results
    FROM movies m
        INNER JOIN movierating mr ON m.movie_id = mr.movie_id
    WHERE YEAR(mr.created_at) = 2020 AND MONTH(mr.created_at) = 2
    GROUP BY m.title
    ORDER BY AVG(mr.rating * 1.0) DESC, m.title
)