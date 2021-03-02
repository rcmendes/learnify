CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


-- public.quizzes definition

-- Drop table

-- DROP TABLE public.quizzes;

-- public.quizzes definition

-- Drop table

-- DROP TABLE public.quizzes;

CREATE TABLE public.quizzes (
	id uuid NOT NULL,
	category_id uuid NOT NULL,
	palavra varchar(60) NOT NULL,
	mot varchar(60) NOT NULL,
	image_filename varchar(255) NOT NULL,
	audio_filename varchar(255) NOT NULL,
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	updated_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT quizzes_pk PRIMARY KEY (id)
);

CREATE TABLE public.categories (
	id uuid NOT NULL,
	name varchar(40) not null unique,
	description varchar(255) null,
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	updated_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT categories_pk PRIMARY KEY (id)
);

CREATE TABLE public.games (
	id uuid NOT NULL,
	player_id uuid NOT NULL,
	status smallint NOT null default 1,
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	updated_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT game_pk PRIMARY KEY (id)
);

CREATE TABLE public.game_quizzes (
	game_id uuid NOT NULL,
	quiz_id uuid not null,
	updated_at timestamptz(0) NOT NULL DEFAULT now(),
	status smallint not null default 1,
	CONSTRAINT game_quizzes_pk PRIMARY KEY (game_id, quiz_id)
);

CREATE TABLE public.players (
	id uuid NOT NULL,
	name varchar(40) NOT NULL,
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT player_pk PRIMARY KEY (id)
);


ALTER TABLE public.quizzes ADD "uuid" uuid not NULL default uuid_generate_v4();
ALTER TABLE public.quizzes ADD updated_at timestamptz(0) NOT NULL DEFAULT now();



insert into players(id, name) values (uuid_generate_v4(), 'Guga')


select id, tag, palavra, url, mot from quizzes q where q.tag ='animais' order by q.tag;

select id, tag, palavra, url from quizzes q where q.tag ='animais' order by q.tag;




select * from players where id like '20f2777e-6bf8-409b-affc-9f7731eea714' as uuid
