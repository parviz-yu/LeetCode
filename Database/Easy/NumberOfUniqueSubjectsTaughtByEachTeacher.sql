/* Write your T-SQL query statement below */
-- First approach
SELECT teacher_id, COUNT(cnt) AS cnt
FROM (
  SELECT teacher_id, COUNT(subject_id) AS cnt
  FROM teacher
  GROUP BY teacher_id, subject_id
) q
GROUP BY teacher_id;

/* Write your T-SQL query statement below */
-- Efficient approach
SELECT teacher_id, COUNT(DISTINCT subject_id) AS cnt
FROM teacher
GROUP BY teacher_id;