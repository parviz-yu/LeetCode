/* Write your T-SQL query statement below */
WITH department_max_salary AS (
    SELECT departmentId, MAX(salary) AS 'max_salary'
    FROM employee
    GROUP BY departmentId
)

SELECT
    d.name AS 'Department',
    e.name AS 'Employee',
    dms.max_salary AS 'Salary'
FROM employee e
    INNER JOIN department_max_salary dms ON e.departmentId = dms.departmentId AND e.salary = dms.max_salary
    INNER JOIN department d ON d.id = e.departmentId;