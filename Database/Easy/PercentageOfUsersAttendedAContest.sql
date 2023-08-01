# Write your MySQL query statement below
SELECT contest_id, ROUND(COUNT(r.user_id) * 100.0 / (SELECT COUNT(*) FROM users), 2) AS 'percentage'
FROM register r
    LEFT JOIN users u ON u.user_id = r.user_id
GROUP BY contest_id
ORDER BY percentage DESC, contest_id

/* Write your T-SQL query statement below */
DECLARE @len int = (SELECT COUNT(*) FROM users)

SELECT contest_id, ROUND(COUNT(r.user_id) * 100.0 / @len, 2) AS 'percentage'
FROM register r
GROUP BY contest_id
ORDER BY [percentage] DESC, contest_id