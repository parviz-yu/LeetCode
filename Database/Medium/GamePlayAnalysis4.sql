/* Write your T-SQL query statement below */
-- First Option
SELECT ROUND(COUNT(*) * 1.0 / (SELECT COUNT(DISTINCT player_id) FROM activity), 2) AS 'fraction'
FROM (
    SELECT player_id, MIN(event_date) AS 'first_date'
    FROM activity
    GROUP BY player_id
) q1 INNER JOIN (
    SELECT 
        player_id,
        event_date AS 'tod',
        LEAD(event_date) OVER (PARTITION BY player_id ORDER BY event_date) AS 'tom'
    FROM activity
) q2 ON q1.player_id = q2.player_id AND q1.first_date = q2.tod
WHERE DATEDIFF(day, q2.tod, q2.tom) = 1;

-- Second Option
WITH first_logins AS (
    SELECT player_id, MIN(event_date) AS 'first_login'
    FROM activity
    GROUP BY player_id
), consec_logins AS (
    SELECT COUNT(a.player_id) AS 'num_logins'
    FROM first_logins f
    INNER JOIN activity a ON f.player_id = a.player_id AND DATEADD(day, 1, f.first_date) = a.event_date
)

SELECT ROUND(
    (SELECT c.num_logins FROM consec_logins c) * 1.0 / (SELECT COUNT(f.player_id) FROM first_logins f), 2
) AS 'fraction';