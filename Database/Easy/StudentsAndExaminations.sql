/* Write your T-SQL query statement below */
SELECT 
  s.student_id, 
  s.student_name, 
  sub.subject_name, 
  ISNULL(query.cnt, 0) AS 'attended_exams' 
FROM 
  students s
  CROSS JOIN subjects sub 
  LEFT JOIN (
    SELECT 
      student_id, 
      subject_name, 
      COUNT(subject_name) AS 'cnt' 
    FROM 
      examinations 
    GROUP BY 
      student_id, 
      subject_name
  ) query ON s.student_id = query.student_id 
  AND sub.subject_name = query.subject_name 
ORDER BY 
  s.student_id, 
  sub.subject_name
