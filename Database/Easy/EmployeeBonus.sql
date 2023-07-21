/* Write your T-SQL query statement below */
-- First Option
SELECT e.name, b.bonus
FROM employee e
    LEFT JOIN bonus b ON e.empId = b.empId
WHERE bonus < 1000 OR bonus IS NULL

-- Second Option
SELECT e.name, b.bonus
FROM employee e
    LEFT JOIN bonus b ON e.empId = b.empId
WHERE ISNULL(b.bonus, 0) < 1000