/* Write your T-SQL query statement below */
-- First Option
SELECT 
  (
    SELECT 
      salary
    FROM 
      employee 
    GROUP BY 
      salary 
    ORDER BY 
      salary DESC OFFSET 1 ROW FETCH NEXT 1 ROW ONLY
) AS 'SecondHighestSalary' 

-- Second Option
SELECT 
  MAX(salary) AS 'SecondHighestSalary' 
FROM 
  employee 
WHERE 
  salary < (
    SELECT 
      MAX(salary) 
    FROM 
      employee
  )