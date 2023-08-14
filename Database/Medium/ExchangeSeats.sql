/* Write your T-SQL query statement below */
-- First Option
SELECT
    id, 
    CASE WHEN id % 2 = 1 AND [next] IS NOT NULL THEN [next]
    WHEN id % 2 = 0 THEN prev
    ELSE student
    END AS 'student'
FROM (
    SELECT
        id,
        student,
        LAG(student, 1) OVER (ORDER BY id) AS 'prev',
        LEAD(student, 1) OVER (ORDER BY id) AS 'next'
    FROM seat
) sub_query;

/* Write your T-SQL query statement below */
-- Second Option
SELECT
    CASE WHEN id = (SELECT MAX(id) FROM seat) AND id % 2 = 1 THEN id
    WHEN id % 2 = 1 THEN id + 1
    WHEN id % 2 = 0 THEN id - 1 
    END AS 'id',
    student
from seat
order by id;