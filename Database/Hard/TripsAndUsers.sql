/* Write your T-SQL query statement below */
-- Effective approach
SELECT
    request_at AS [Day],
    ROUND(SUM(CASE WHEN [status] IN ('cancelled_by_driver', 'cancelled_by_client') THEN 1 ELSE 0 END) * 1.0 / COUNT(status), 2) AS 'Cancellation Rate'
FROM trips t
    INNER JOIN users c ON t.client_id = c.users_id
    INNER JOIN users d ON t.driver_id = d.users_id
WHERE c.banned = 'No' AND d.banned = 'No' AND request_at BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY request_at;

/* Write your T-SQL query statement below */
-- First Approach
WITH unbanned_clients AS (
    SELECT users_id
    FROM users
    WHERE [role] = 'client' AND banned = 'No'
), unbanned_drivers AS (
    SELECT users_id
    FROM users
    WHERE [role] = 'driver' AND banned = 'No'
)

SELECT
    DISTINCT request_at AS 'Day',
    ROUND(cancelNum * 1.0 / (completeNum + cancelNum), 2) AS 'Cancellation Rate'
FROM (
    SELECT 
    request_at,
    SUM(CASE WHEN [status]='completed' THEN 1 ELSE 0 END) OVER(PARTITION BY request_at) AS completeNum,
    SUM(CASE WHEN [status] IN ('cancelled_by_client', 'cancelled_by_driver') THEN 1 ELSE 0 END) OVER(PARTITION BY request_at) AS cancelNum
FROM trips t
    INNER JOIN unbanned_clients c ON t.client_id = c.users_id
    INNER JOIN unbanned_drivers d ON t.driver_id = d.users_id
WHERE request_at BETWEEN '2013-10-01' AND '2013-10-03'
) q;