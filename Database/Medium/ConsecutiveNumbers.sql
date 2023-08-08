/* Write your T-SQL query statement below */
SELECT curr AS 'ConsecutiveNums'
FROM (
  SELECT
    num AS 'curr',
    LAG(num) OVER(order by id) AS 'prev',
    LEAD(num) OVER(order by id) AS 'next'
  FROM logs
) q
WHERE curr = prev AND curr = [next]
GROUP BY curr;