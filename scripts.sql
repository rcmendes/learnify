-- public.quizzes definition

-- Drop table

-- DROP TABLE public.quizzes;

CREATE TABLE public.quizzes (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	category varchar(40) NOT NULL,
	palavra varchar(60) NOT NULL,
	mot varchar(60) NOT NULL,
	image_filename varchar(255) NOT NULL,
	audio_filename varchar(255) NOT NULL,
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	updated_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT quizzes_pk PRIMARY KEY (id)
);

CREATE TABLE public.game_quiz (
	game_id uuid NOT NULL,
	quiz_id uuid not null,
	updated_at timestamptz(0) NOT NULL DEFAULT now(),
	success bool null,
	CONSTRAINT game_quizzes_pk PRIMARY KEY (game_id, quiz_id)
);

CREATE TABLE public.game (
	id uuid NOT NULL,
	player_id uuid NOT NULL,
	status char NOT null default 'C',
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	updated_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT game_pk PRIMARY KEY (id)
);

CREATE TABLE public.player (
	id uuid NOT NULL,
	name varchar(100) NOT NULL,
	status char NOT null default 'C',
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	updated_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT player_pk PRIMARY KEY (id)
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
ALTER TABLE public.quizzes ADD "uuid" uuid not NULL default uuid_generate_v4();
ALTER TABLE public.quizzes ADD updated_at timestamptz(0) NOT NULL DEFAULT now();



insert into player(id, name) values (uuid_generate_v4(), 'Guga')


select id, tag, palavra, url, mot from quizzes q where q.tag ='animais' order by q.tag;

select id, tag, palavra, url from quizzes q where q.tag ='animais' order by q.tag;