/* Write your T-SQL query statement below */
-- First approach
SELECT
    stock_name,
    SUM(CASE WHEN operation = 'Buy' THEN -1 * price ELSE price END) AS capital_gain_loss
FROM stocks
GROUP BY stock_name;

/* Write your T-SQL query statement below */
-- Second approach
WITH sells_buys AS(
    SELECT
        stock_name,
        SUM(CASE WHEN operation = 'Buy' THEN price END) AS buys,
        SUM(CASE WHEN operation = 'Sell' THEN price END) AS sells
    FROM Stocks 
    GROUP BY stock_name
)

SELECT stock_name, (sells - buys) AS capital_gain_loss 
FROM sells_buys;