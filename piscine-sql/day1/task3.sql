SELECT DISTINCT person_id FROM person_visits 
WHERE visit_date >= '2022-01-06' AND visit_date <= '2022-01-09' OR pizzeria_id = 2
ORDER by person_id DESC
