/* Write your T-SQL query statement below */
WITH t AS (
  SELECT requester_id AS 'id', COUNT(requester_id) 'num'
  FROM (
    SELECT requester_id
    FROM RequestAccepted

    UNION ALL

    SELECT accepter_id
    FROM RequestAccepted
  ) all_rel
  GROUP BY requester_id
)

SELECT TOP 1 id, num
FROM t
ORDER BY num DESC;