/* Write your T-SQL query statement below */
SELECT u.name, SUM(amount) AS 'balance'
FROM transactions t
    INNER JOIN users u ON t.account = u.account
GROUP BY u.name
HAVING SUM(amount) > 10000