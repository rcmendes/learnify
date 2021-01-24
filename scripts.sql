CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE public.assets (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	category varchar(40) NOT NULL,
	palavra varchar(60) NOT NULL,
	mot varchar(60) NOT NULL,
	image_filename varchar(255) NOT NULL,
	audio_filename varchar(255) NOT NULL,
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	updated_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT assets_pk PRIMARY KEY (id)
);


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


ALTER TABLE public.quizzes ADD "uuid" uuid not NULL default uuid_generate_v4();
ALTER TABLE public.quizzes ADD updated_at timestamptz(0) NOT NULL DEFAULT now();



insert into player(id, name) values (uuid_generate_v4(), 'Guga')


select id, tag, palavra, url, mot from quizzes q where q.tag ='animais' order by q.tag;

select id, tag, palavra, url from quizzes q where q.tag ='animais' order by q.tag;











insert into assets(image_filename, category, mot, palavra, audio_filename) VALUES
('alegre.jpg','sentimentos','heureux','alegre','Il a heureux'),
('raiva.jpg','sentimentos','rage','raiva','Il a rage'),
('triste.jpg','sentimentos','triste','triste','Il a triste'),
('medo.jpg','sentimentos','peur','medo','Il a peur'),
('sono.jpg','sentimentos','sommeil','sono','Il a sommeil'),
('calme.jpg','sentimentos','calme','calmo','Il a calme'),
('gato.jpg','animais','chat','gato','le chat'),
('','vestuário','robe','vestido','la robe'),
('https://cdn3.vectorstock.com/i/1000x1000/75/47/black-color-splash-vector-1097547.jpg','cores','noir ','preto','noir'),
('https://thumbs.dreamstime.com/b/color-image-cartoon-oranges-white-background-fruits-vector-illustration-165217615.jpg','frutas','orange','laranja','l''orange'),
('https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcQhVYavbmxaprx3dvCGkaD6grjgUQlnF-lwbA&usqp=CAU','animais','requin','tubarão','le requin'),
('https://lh3.googleusercontent.com/proxy/GQnxCEACtRpddIZS4JY83amYRbTofbiz97O5weKMr900gS91HPr03k1OuxAbu0awwNDyipb8df8LS18pPeOBVakj5IR0OW9DK4k0W7-ij000Vh-9RqPllTpGoKk08QoJ5AuaeQ','vestuário','vètement','roupa',''),
('https://i.pinimg.com/originals/d2/16/e5/d216e5c5b7e7fef30c79057cd2d85222.jpg','vestuário','sac','bolsa','le sac'),
('https://image.freepik.com/free-vector/cute-boy-cartoon_33070-2118.jpg','pessoas','garçon','menino','le garçon'),
('https://image.freepik.com/free-vector/brunette-woman-character_23-2147502479.jpg','pessoas','femme','mulher','la femme'),
('https://www.netclipart.com/pp/m/110-1103220_paintball-splat-cliparts-paint-splash-cartoon.png','cores','vert','verde',''),
('https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcSIStIzlQlMJTKljn8WP1H23TZDhIC4DdrZ9Q&usqp=CAU','cores','bleu','azul',''),
('https://freesvg.org/img/mealplate.png','alimentos','repas','refeição',''),
('https://image.freepik.com/free-vector/cute-little-dog-cartoon-isolated-white_143596-3.jpg','animais','chien','cachorro','le chien'),
('https://www.clipartkey.com/mpngs/m/137-1378551_commerce-vector-cartoon-business-man-hq-image-free.png','pessoas','homme','homem','le homme'),
('https://st.depositphotos.com/2400497/2883/v/450/depositphotos_28833263-stock-illustration-cartoon-girl-eating.jpg','ações','manger','comer',''),
('https://i.dlpng.com/static/png/6540632_preview.png','cores','blanche','branco',''),
('https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcR48hDZWbUuKtN3CARMhGJSuNSbaoRivbwNmQ&usqp=CAU','frutas','pomme','maçã',''),
('https://lh3.googleusercontent.com/proxy/5LhEMMQv-oEvlCa0l5wJzxmtm7NsOniMskJ0Q_Fl9-M2KQffbEAHKOFp2Wo1at1JwMMd9jKNNOQqbMO13kT012GEz9thKumcE5rYsoJq6De1odriP0lqjH2jcsfFgCg','objetos','livre','livro',''),
('https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRMYzZ1HDRi41Sm1CsD89cOfQZL_wX-UYZPfg&usqp=CAU','frutas','citron','limão',''),
('https://img.freepik.com/free-vector/milk-tetrabrick_51188-5.jpg?size=338&ext=jpg','alimentos','lait','leite',''),
('https://image.shutterstock.com/image-vector/child-drinking-water-vector-illustration-260nw-330125195.jpg','ações','-','beber',''),
('https://i.pinimg.com/736x/2f/68/23/2f68236fdffd49cf8854f254aad5e910.jpg','ações','-','dormir',''),
('https://www.pngitem.com/pimgs/m/532-5322872_kids-play-png-transparent-cartoons-cartoon-images-of.png','ações','-','brincar',''),
('https://cdn2.vectorstock.com/i/1000x1000/46/11/little-boy-cartoon-walking-vector-4414611.jpg','ações','-','caminhar',''),
('https://cdn5.vectorstock.com/i/1000x1000/12/44/kid-sitting-on-chair-vector-20551244.jpg','ações','-','sentar',''),
('https://cdn4.vectorstock.com/i/1000x1000/05/33/flat-denim-jeans-icon-vector-18960533.jpg','vestuário','pantalon','calça','le pantalon'),
('https://image.shutterstock.com/image-vector/yellow-sneakers-vector-isolated-260nw-569243245.jpg','vestuário','chaussure','sapato',''),
('','objetos','lettres','cartas',''),
('https://www.clipartkey.com/mpngs/m/144-1449759_drop-clip-art-colours-red-splodge.png','cores','rouge','vermelho',''),
('https://images-na.ssl-images-amazon.com/images/I/710bPncUkKL._AC_SL1500_.jpg','animais','poisson','peixe',''),
('https://www.netclipart.com/pp/m/27-270944_cooked-turkey-clipart-black-and-white-cooked-meat.png','alimentos','boeuf','carne',''),
('','adjetivos','leve','légère',''),
('https://fsb.zobj.net/crop.php?r=67k5Vxa6DCKDWWgyQNyW3GnT5upfIpkb9Qg1gWle4ZBJOD1xWq9bMGjh4odEw0KAZ_e5L240kFb_3XZNxbKQB4NFO29FlWPELo-lpDZnfPMkpsu6vCdbLVPij1gJ5NFouLCsoP_m7mkp6Q4m','ações','aime','gostar','j''aime'),
('https://cdn2.f-cdn.com/contestentries/1458359/21383270/5c0fca8e25198_thumb900.jpg','pessoas','fille','menina',''),
('https://image.freepik.com/free-vector/kid-drawing-vector-illustration-isolated_97632-299.jpg','ações','dessiner','desenhar',''),
('https://thumbs.dreamstime.com/b/two-little-kids-talking-counting-fingers-two-little-kids-boy-girl-pink-talking-being-friends-180131114.jpg','ações','parle','falar',''),
('https://i.pinimg.com/originals/90/c8/7c/90c87cb3d02f23c2a33e5b8786788e95.jpg','alimentos','sandwich','sanduiche','Le sandwich'),
('https://png.pngtree.com/element_our/20190602/ourlarge/pngtree-yellow-cheese-cartoon-illustration-image_1423401.jpg','alimentos','fromage','queijo',''),
('https://thumbs.dreamstime.com/b/cenoura-dos-desenhos-animados-110455651.jpg','alimentos','carotte','cenoura',''),
('https://image.freepik.com/free-vector/grapes-cartoon-icon_90744-153.jpg','frutas','raisin','uva','le raisin'),
('https://i.pinimg.com/736x/ca/ff/ac/caffac214721104ab6a7a76ecda33aec.jpg','alimentos','chocolat','chocolate','le chocolat'),
('','alimentos','oeuf','ovo',''),
('https://cdn4.vectorstock.com/i/1000x1000/50/33/multi-ethnic-group-of-children-forming-a-circle-vector-6725033.jpg','pessoas','enfant','criança',''),
('https://www.jing.fm/clipimg/detail/4-43044_loaf-bread-png-clip-art-clip-art-loaf.png','alimentos','baguette','pão',''),
('https://image.freepik.com/free-vector/cute-smiling-happy-glass-water-drop-flat-cartoon-character-illustration-isolated-white-background-water-drop-character-concept_92289-1426.jpg','alimentos','l''eau','água',''),
('https://image.freepik.com/free-vector/cartoon-brown-horse-isolated-white_29190-5049.jpg','animais','cheval','cavalo',''),
('https://i.pinimg.com/736x/62/b7/fb/62b7fb57ec0e461b6daf6a928edc87ee.jpg','animais','papillon ','borboleta',''),
('https://image.freepik.com/free-vector/adorable-dolphin-cartoon_74769-37.jpg','animais','dauphin','golfinho',''),
('https://image.freepik.com/free-vector/cute-mouse-cartoon-presenting_160606-84.jpg','animais','souris','rato',''),
('https://image.freepik.com/free-vector/happy-elephant-cartoon_160606-95.jpg','animais','éléphant','elefante','');