/* Write your T-SQL query statement below */
-- First approach -> Gaps and Islands Problem
WITH groupped AS (
  SELECT
    *,
    COUNT(dif) OVER(PARTITION BY dif) AS num_in_group
  FROM (
    SELECT
      *,
      id - DENSE_RANK() OVER(ORDER BY id) AS dif
    FROM (
      SELECT *
      FROM stadium
      WHERE people > 99
    ) filtered
  ) diff
)

SELECT id, visit_date, people
FROM groupped
WHERE num_in_group >= 3
ORDER BY visit_date;

/* Write your T-SQL query statement below */
-- Second approach
WITH extended AS (
  SELECT
    id,
    visit_date,
    people,
    LEAD(people, 1) OVER(ORDER BY id) [next],
    LEAD(people, 2) OVER(ORDER BY id) [next2],
    LAG(people, 1) OVER(ORDER BY id) [prev],
    LAG(people, 2) OVER(ORDER BY id) [prev2]
  FROM stadium
)

SELECT id, visit_date, people
FROM extended e
WHERE (e.people >= 100 AND e.next >= 100 AND e.next2 >= 100)
  OR (e.people >= 100 AND e.next >= 100 AND e.prev >= 100)
  OR (e.people >= 100 AND e.prev >= 100 AND e.prev2 >= 100);

