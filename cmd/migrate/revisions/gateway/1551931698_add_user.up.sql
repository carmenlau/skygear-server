CREATE TABLE "user" (
	id uuid PRIMARY KEY,
	created_at timestamp WITHOUT TIME ZONE NOT NULL,
	updated_at timestamp WITHOUT TIME ZONE NOT NULL,
	email text NOT NULL,
	password text NOT NULL
);
