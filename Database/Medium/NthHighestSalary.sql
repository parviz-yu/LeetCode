-- First Option
CREATE FUNCTION getNthHighestSalary(@N INT) RETURNS INT AS
BEGIN
    RETURN (
        /* Write your T-SQL query statement below. */
        SELECT salary
        FROM employee
        GROUP BY salary
        ORDER BY salary DESC
        OFFSET @N - 1 ROWS FETCH NEXT 1 ROW ONLY
    );
END

-- Second Option
CREATE FUNCTION getNthHighestSalary(@N INT) RETURNS INT AS
BEGIN
    RETURN (
      SELECT MAX(SALARY) AS getNthHighestSalary
      FROM (
            SELECT SALARY,  DENSE_RANK() OVER (ORDER BY SALARY DESC) AS ROWNO
            FROM EMPLOYEE
            ) E
      WHERE E.ROWNO = @N
        
    );
END