/* Write your T-SQL query statement below */
-- First Option
SELECT employee_id, department_id
FROM employee
WHERE primary_flag = 'Y'
UNION
SELECT employee_id, department_id
FROM employee
WHERE employee_id IN (
  SELECT employee_id
  FROM employee
  GROUP BY employee_id
  HAVING COUNT(department_id) = 1
);


/* Write your T-SQL query statement below */
-- Second Option
SELECT employee_id, department_id
FROM (
  SELECT *,
    COUNT(department_id) OVER (PARTITION BY employee_id ORDER BY primary_flag DESC) AS 'r'
  FROM employee
) q
WHERE r = 1 OR (r > 1 AND primary_flag = 'Y');
