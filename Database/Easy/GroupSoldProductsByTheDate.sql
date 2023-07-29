# Write your MySQL query statement below
SELECT sell_date, COUNT(DISTINCT product) AS num_sold,
GROUP_CONCAT(DISTINCT product ORDER BY product) AS products
FROM activities
GROUP BY sell_date
ORDER BY sell_date;

/* Write your T-SQL query statement below */
SELECT 
  sell_date, 
  COUNT(DISTINCT product) AS 'num_sold', 
  STRING_AGG(product, ',') WITHIN GROUP (ORDER BY product) AS 'products' 
FROM 
  (
    SELECT 
      DISTINCT sell_date, 
      product 
    FROM 
      activities
  ) ac 
GROUP BY 
  ac.sell_date 
ORDER BY 
  ac.sell_date
