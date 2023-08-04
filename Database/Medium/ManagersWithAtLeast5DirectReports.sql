/* Write your T-SQL query statement below */
SELECT m.name
FROM employee m
    INNER JOIN employee e ON m.id = e.managerId
GROUP BY m.id, m.name
HAVING COUNT(e.name) > 4