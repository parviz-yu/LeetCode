/* Write your T-SQL query statement below */
SELECT
  [name],
  ISNULL(SUM(distance), 0) AS 'travelled_distance'
FROM users u
LEFT JOIN rides r
  ON u.id = r.user_id
GROUP BY u.id,
         [name]
ORDER BY travelled_distance DESC, [name]