/* Write your T-SQL query statement below */

-- First Option
SELECT w2.id AS 'Id'
FROM Weather w1, Weather w2
WHERE  DATEDIFF(day, w1.recordDate, w2.recordDate) = 1 AND w2.temperature > w1.temperature

-- Second Option
SELECT w2.id AS 'Id'
FROM Weather w1
    INNER JOIN Weather w2 ON DATEDIFF(day, w1.recordDate, w2.recordDate) = 1
WHERE w2.temperature > w1.temperature