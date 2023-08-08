/* Write your T-SQL query statement below */
SELECT
    d.name AS 'Department',
    s.name AS 'Employee',
    s.salary AS 'Salary'
FROM (
    SELECT
        *,
        DENSE_RANK() OVER (PARTITION BY departmentId ORDER BY salary DESC) as 'salary_rank'
    FROM employee
) s INNER JOIN department d ON s.departmentId = d.id
WHERE s.salary_rank < 4;