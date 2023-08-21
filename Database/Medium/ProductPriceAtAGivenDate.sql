/* Write your T-SQL query statement below */
SELECT
    p1.product_id,
    ISNULL(p2.new_price, 10) AS price
FROM (
    SELECT product_id
    FROM products
    GROUP BY product_id
) p1 LEFT JOIN (
    SELECT p1.product_id, p1.new_price
    FROM products p1
        INNER JOIN (
            SELECT product_id, MAX(change_date) AS last_price
            FROM products
            WHERE change_date <= '2019-08-16'
            GROUP BY product_id
        ) p2 ON p1.product_id = p2.product_id AND p1.change_date = p2.last_price
) p2 ON p1.product_id = p2.product_id;
