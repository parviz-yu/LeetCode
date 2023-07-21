-- First Option
DELETE p1 FROM Person p1,
    Person p2
WHERE
    p1.Email = p2.Email AND p1.Id > p2.Id

-- Second Option
DELETE FROM Person
WHERE id IN (
    SELECT p2.id
    FROM Person p1
    INNER JOIN Person p2 ON p1.email = p2.email AND p2.id > p1.id
)