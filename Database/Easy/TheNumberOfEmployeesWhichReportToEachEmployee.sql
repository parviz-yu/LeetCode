/* Write your T-SQL query statement below */
SELECT
    m.employee_id,
    m.name,
    COUNT(e.name) AS 'reports_count',
    ROUND(AVG(e.age * 1.0), 0) AS 'average_age' 
FROM employees m
    INNER JOIN employees e ON m.employee_id = e.reports_to
GROUP BY m.employee_id, m.name
ORDER BY m.employee_id;