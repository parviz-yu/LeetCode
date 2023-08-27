/* Write your T-SQL query statement below */
-- First approach
SELECT e.employee_id
FROM employees e
    LEFT JOIN employees m ON e.manager_id = m.employee_id
WHERE e.salary < 30000 AND m.employee_id IS NULL AND e.manager_id IS NOT NULL
ORDER BY e.employee_id;

/* Write your T-SQL query statement below */
-- Efficient approach
SELECT employee_id
FROM Employees
WHERE manager_id NOT IN (
    SELECT employee_id
    FROM Employees
) AND salary <30000 ORDER BY employee_id;