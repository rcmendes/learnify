CREATE TABLE tb_categories (
	id SERIAL NOT NULL,
	uuid uuid NOT NULL,
	name varchar(60) NOT NULL,
	description varchar(255) NULL,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at timestamp NULL,
	CONSTRAINT tb_categories_pkey PRIMARY KEY (id),
	CONSTRAINT tb_categories_uuid_key UNIQUE (uuid),
	CONSTRAINT tb_categories_name_key UNIQUE (name)
);