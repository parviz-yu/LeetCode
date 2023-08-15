/* Write your T-SQL query statement below */
SELECT 
    u.user_id AS 'buyer_id',
    u.join_date,
    ISNULL(sub.ors, 0) AS 'orders_in_2019'
FROM users u 
    LEFT JOIN (
        SELECT o.buyer_id, COUNT(o.item_id) AS 'ors'
        FROM orders o
        WHERE YEAR(o.order_date) = 2019
        GROUP BY o.buyer_id
    ) sub ON u.user_id = sub.buyer_id;