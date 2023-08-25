/* Write your T-SQL query statement below */
WITH uniq AS (
    SELECT DISTINCT visited_on
    FROM customer
    WHERE visited_on >= DATEADD(day, 6, (SELECT MIN(visited_on) FROM customer))   
)

SELECT u.visited_on, SUM(c.amount) AS amount, ROUND(SUM(c.amount)* 1.0 / 7, 2) AS average_amount
FROM uniq u
    LEFT JOIN customer c ON u.visited_on >= c.visited_on AND u.visited_on <= DATEADD(day, 6, c.visited_on)
GROUP BY u.visited_on
ORDER BY u.visited_on;

