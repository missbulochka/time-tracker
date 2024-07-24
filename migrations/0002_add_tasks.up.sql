INSERT INTO tasks (task_name)
VALUES
	('research'),
	('coffee_break'),
	('coding'),
	('smoothie_break')
ON CONFLICT DO NOTHING;