/* Write your T-SQL query statement below */
SELECT ROUND(SUM(TIV_2016), 2) AS tiv_2016
FROM (
    SELECT *,
        COUNT(*) OVER (PARTITION BY TIV_2015) AS cnt1,
        COUNT(*) OVER (PARTITION BY LAT, LON) AS cnt2
    FROM insurance
) sub_query
WHERE cnt1 > 1 AND cnt2 = 1;