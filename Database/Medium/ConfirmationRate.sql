/* Write your T-SQL query statement below */
-- First Option
SELECT
    s.user_id,
    ROUND(ISNULL(SUM(CASE WHEN c.action = 'confirmed' THEN 1 END) * 1.0 / (SELECT COUNT(*) FROM confirmations WHERE [user_id] = s.user_id), 0), 2) AS confirmation_rate
FROM signups s
    LEFT JOIN confirmations c ON s.user_id = c.user_id
GROUP BY s.user_id;

/* Write your T-SQL query statement below */
-- Faster Option
SELECT
    s.user_id,
    ROUND(
        SUM(CASE
        WHEN c.action = 'confirmed'THEN 1.00
        WHEN c.action IS NULL THEN 0
        ELSE 0 END) / COUNT(*), 2) AS confirmation_rate
FROM signups s
    LEFT JOIN confirmations c ON s.user_id = c.user_id
GROUP BY s.user_id