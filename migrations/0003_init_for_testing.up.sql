INSERT INTO users
(
	passport_number,
	name
)
VALUES
	('1234 567890', 'elina'),
	('1234 567890', 'pavel'),
	('0987 654321', 'nikita')
ON CONFLICT DO NOTHING;