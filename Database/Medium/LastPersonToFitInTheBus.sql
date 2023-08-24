/* Write your T-SQL query statement below */
WITH cumm_t AS (
    SELECT 
        person_name,
        [weight],
        SUM([weight]) OVER(ORDER BY turn) AS cum
    FROM queue
)

SELECT TOP 1 person_name
FROM cumm_t
WHERE cum <= 1000
ORDER BY cum DESC;

-- # Write your MySQL query statement below
SELECT q1.person_name
FROM queue q1
    INNER JOIN queue q2 ON q1.turn >= q2.turn
GROUP BY q1.turn
HAVING SUM(q2.weight) <= 1000
ORDER BY SUM(q2.weight) DESC
LIMIT 1;