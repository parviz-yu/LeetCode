-- Write your PostgreSQL query statement below
SELECT p.product_id,
    p.product_name
FROM (
        SELECT product_id
        FROM sales
        GROUP BY product_id
        HAVING MIN(sale_date) >= '2019-01-01'
            AND MAX(sale_date) <= '2019-03-31'
    ) t
    INNER JOIN product p ON t.product_id = p.product_id