/* Write your T-SQL query statement below */
SELECT eu.unique_id, e.name
FROM employees e
    LEFT JOIN employeeuni eu ON e.id = eu.id