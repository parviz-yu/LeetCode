/* Write your T-SQL query statement below */
-- First Option
SELECT p.project_id, ROUND(AVG(CAST(e.experience_years AS DECIMAL(10,2))), 2) AS 'average_years'
FROM employee e
    INNER JOIN project p ON e.employee_id = p.employee_id
GROUP BY p.project_id

-- Second Option
SELECT p.project_id, ROUND(AVG(e.experience_years * 1.0), 2) AS 'average_years'
FROM employee e
    INNER JOIN project p ON e.employee_id = p.employee_id
GROUP BY p.project_id