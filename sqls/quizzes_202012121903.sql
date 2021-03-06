-- public.quizzes definition

-- Drop table

-- DROP TABLE public.quizzes;

CREATE TABLE public.quizzes (
	id serial NOT NULL,
	tag varchar(40) NOT NULL,
	url varchar(255) NOT NULL,
	mot varchar(40) NOT NULL,
	palavra varchar(40) NOT NULL,
	audio varchar(60) NOT NULL,
	uuid uuid NOT NULL DEFAULT uuid_generate_v4(),
	created_at timestamptz(0) NOT NULL DEFAULT now(),
	CONSTRAINT quizzes_pkey PRIMARY KEY (id)
);


INSERT INTO public.quizzes (tag,url,mot,palavra,audio,uuid,created_at,updated_at) VALUES
	 ('sentimentos','https://www.criandocomapego.com/wp-content/uploads/2018/11/fichas-das-emocoes-para-imprimir-alegria.jpg','heureux','alegre','Il a heureux','bebbb5c9-ffe6-4bce-b4ab-b98fe0504757','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('cores','https://cdn.pixabay.com/photo/2013/07/13/10/45/bucket-157730_960_720.png','bleu','azul','','4e2d6ae6-ec0b-4700-ab56-a3e32c53941f','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('ações','https://image.shutterstock.com/image-vector/child-drinking-water-vector-illustration-260nw-330125195.jpg','-','beber','','d2a71f4e-535f-404f-bc7e-3da8a8b5e106','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('vestuário','https://i.pinimg.com/originals/d2/16/e5/d216e5c5b7e7fef30c79057cd2d85222.jpg','sac','bolsa','le sac','33bc0a49-a7cb-4773-99f0-12ef8d73f5cb','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('animais','https://i.pinimg.com/736x/62/b7/fb/62b7fb57ec0e461b6daf6a928edc87ee.jpg','papillon ','borboleta','','88deffe2-fa69-4cec-aba7-694699213657','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('cores','https://cdn.pixabay.com/photo/2013/07/13/10/48/paint-pot-157814_960_720.png','blanche','branco','','f11af346-a529-4bc6-bdef-89cf06498aaa','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('ações','https://www.pngitem.com/pimgs/m/532-5322872_kids-play-png-transparent-cartoons-cartoon-images-of.png','-','brincar','','81f9221c-6709-4670-8cc6-d2b1313bf599','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('animais','https://image.freepik.com/free-vector/cute-little-dog-cartoon-isolated-white_143596-3.jpg','chien','cachorro','le chien','d1196953-242f-42df-8d7f-e1d77f839e12','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('sentimentos','https://www.criandocomapego.com/wp-content/uploads/2018/11/fichas-das-emocoes-para-imprimir-calma.jpg','calme','calmo','Il a calme','9dd732d4-7421-4e09-8ec9-fa5dff4d2064','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('vestuário','https://cdn.pixabay.com/photo/2016/03/31/19/24/clothes-1294974_960_720.png','pantalon','calça','le pantalon','381c3127-08b7-4ab8-989a-9577dcd12605','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05');
INSERT INTO public.quizzes (tag,url,mot,palavra,audio,uuid,created_at,updated_at) VALUES
	 ('ações','https://cdn.pixabay.com/photo/2017/01/31/17/26/boy-2025737_960_720.png','promenader','caminhar','','5a2eb123-8c72-4be9-88ac-3bf93e9d2817','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('alimentos','https://cdn.pixabay.com/photo/2018/02/21/23/23/steak-3171790_960_720.png','boeuf','carne','','8ed02b9c-be8b-45e9-bc5f-deaec603a5a3','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('objetos','https://cdn.pixabay.com/photo/2013/07/12/13/58/letter-147697_960_720.png','lettres','cartas','','edb2d288-9be7-4d7e-80e0-5b73f3dd38f1','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('animais','https://image.freepik.com/free-vector/cartoon-brown-horse-isolated-white_29190-5049.jpg','cheval','cavalo','','7479c993-2bb3-4134-b055-63566bfefbec','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('alimentos','https://thumbs.dreamstime.com/b/cenoura-dos-desenhos-animados-110455651.jpg','carotte','cenoura','','c9f56a1a-e95e-4ef6-a2b1-cebf85a90404','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('alimentos','https://i.pinimg.com/736x/ca/ff/ac/caffac214721104ab6a7a76ecda33aec.jpg','chocolat','chocolate','le chocolat','b280dff8-6f53-4ae3-9385-2c4ca56677c7','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('ações','https://st.depositphotos.com/2400497/2883/v/450/depositphotos_28833263-stock-illustration-cartoon-girl-eating.jpg','manger','comer','','5a2d9fb0-2a50-4d42-a371-c70916dae3ea','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('pessoas','https://cdn.pixabay.com/photo/2020/10/21/19/32/children-5674132_960_720.png','enfant','criança','','2dc298ec-0170-4dc6-871e-27fa514a9584','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('ações','https://image.freepik.com/free-vector/kid-drawing-vector-illustration-isolated_97632-299.jpg','dessiner','desenhar','','d169efe0-54d1-4335-8e55-fd884c56bce6','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('ações','https://i.pinimg.com/736x/2f/68/23/2f68236fdffd49cf8854f254aad5e910.jpg','-','dormir','','dbf3ecc6-ed24-46be-814e-a2e19ff7b88b','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05');
INSERT INTO public.quizzes (tag,url,mot,palavra,audio,uuid,created_at,updated_at) VALUES
	 ('animais','https://image.freepik.com/free-vector/happy-elephant-cartoon_160606-95.jpg','éléphant','elefante','','ff689310-47ea-4e98-9ece-266ae7e8588f','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('ações','https://thumbs.dreamstime.com/b/two-little-kids-talking-counting-fingers-two-little-kids-boy-girl-pink-talking-being-friends-180131114.jpg','parle','falar','','0224afa9-ad94-46aa-a184-ad92ff8ee6dd','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('animais','https://how2drawanimals.com/images/CartoonCat/draw-cartoon-kitten-cat-kitty-thumb.jpg','chat','gato','le chat','55115194-3db6-41c9-b8b6-b95c88f035ad','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('animais','https://image.freepik.com/free-vector/adorable-dolphin-cartoon_74769-37.jpg','dauphin','golfinho','','6051d3bf-bf5b-453c-99e6-a2070d524bc1','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('ações','https://image.freepik.com/free-vector/children-people-love_1308-39261.jpg','aime','gostar','j''aime','ea83cc2c-87f7-4a5e-be3b-5636d71a6269','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('pessoas','https://www.clipartkey.com/mpngs/m/137-1378551_commerce-vector-cartoon-business-man-hq-image-free.png','homme','homem','le homme','3adc5e90-0c66-44bd-b28a-eb5542f56963','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('frutas','https://thumbs.dreamstime.com/b/color-image-cartoon-oranges-white-background-fruits-vector-illustration-165217615.jpg','orange','laranja','l''orange','17fabf5b-fd71-4961-be2c-9763997ec699','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('alimentos','https://img.freepik.com/free-vector/milk-tetrabrick_51188-5.jpg?size=338&ext=jpg','lait','leite','','a5a24a80-dcdf-44e3-9a44-ecf18571c79d','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('adjetivos','https://images-na.ssl-images-amazon.com/images/I/710bPncUkKL._AC_SL1500_.jpg','légère','leve','','aff44d57-9cbf-4007-b82b-c6d06f21dbad','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('frutas','https://image.freepik.com/free-vector/green-lemon-fresh-isolated_99326-187.jpg','citron','limão','','f52a3c16-85ef-4eb3-a9ee-005125ca2126','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05');
INSERT INTO public.quizzes (tag,url,mot,palavra,audio,uuid,created_at,updated_at) VALUES
	 ('objetos','https://cdn.pixabay.com/photo/2017/01/31/00/09/book-2022464_960_720.png','livre','livro','','15e2b032-5178-483f-8204-66428b3cce43','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('frutas','https://image.freepik.com/free-vector/realistic-set-whole-cut-juicy-red-green-apples-isolated-white_1284-33173.jpg','pomme','maçã','','9a8378b3-b0ba-4993-84ab-06536895da4f','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('sentimentos','https://www.criandocomapego.com/wp-content/uploads/2018/11/fichas-das-emocoes-para-imprimir-medo.jpg','peur','medo','Il a peur','591e0fae-d2e0-4895-a3c7-cb576404fa68','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('pessoas','https://cdn2.f-cdn.com/contestentries/1458359/21383270/5c0fca8e25198_thumb900.jpg','fille','menina','','055158ab-477e-4b00-81b2-added11cdd39','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('pessoas','https://image.freepik.com/free-vector/cute-boy-cartoon_33070-2118.jpg','garçon','menino','le garçon','0d3097e6-24f0-48d6-90e4-c91e17a497df','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('pessoas','https://image.freepik.com/free-vector/brunette-woman-character_23-2147502479.jpg','femme','mulher','la femme','a1dc1997-405c-4047-be7e-e8a2e9e761b4','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('alimentos','https://cdn.pixabay.com/photo/2012/04/02/12/49/egg-24404_960_720.png','oeuf','ovo','','af67dd26-3fc9-4c8d-b5db-8e2976cc87b8','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('animais','https://images-na.ssl-images-amazon.com/images/I/710bPncUkKL._AC_SL1500_.jpg','poisson','peixe','','8633af8e-9c62-4fb5-9d2f-5bd049697265','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('cores','https://cdn.pixabay.com/photo/2013/07/13/10/48/paint-pot-157813_960_720.png','noir ','preto','noir','188254ed-0a44-482f-b18c-1f0d045d512a','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('alimentos','https://www.jing.fm/clipimg/detail/4-43044_loaf-bread-png-clip-art-clip-art-loaf.png','baguette','pão','','794bd978-411e-41c5-92ca-eb8bc4b3f06e','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05');
INSERT INTO public.quizzes (tag,url,mot,palavra,audio,uuid,created_at,updated_at) VALUES
	 ('alimentos','https://png.pngtree.com/element_our/20190602/ourlarge/pngtree-yellow-cheese-cartoon-illustration-image_1423401.jpg','fromage','queijo','','41d0e83e-4d92-4d95-8f59-f90192bc3304','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('sentimentos','https://www.criandocomapego.com/wp-content/uploads/2018/11/fichas-das-emocoes-para-imprimir-raiva.jpg','rage','raiva','Il a rage','2053ca92-df2d-4c29-b2b2-97cb8f6f9007','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('animais','https://image.freepik.com/free-vector/cute-mouse-cartoon-presenting_160606-84.jpg','souris','rato','','2252aef9-6bc0-4fda-aaaf-3b87ede9e3c8','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('alimentos','https://freesvg.org/img/mealplate.png','repas','refeição','','44751589-63dd-4407-af13-eef5f367488f','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('vestuário','https://cdn.pixabay.com/photo/2016/03/31/21/19/boy-1296343_960_720.png','vètement','roupa','','3a94313b-2e97-4a18-b7b6-2d882a7d765f','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('alimentos','https://i.pinimg.com/originals/90/c8/7c/90c87cb3d02f23c2a33e5b8786788e95.jpg','sandwich','sanduiche','Le sandwich','9eac927d-d99e-4525-bbc0-d334a3aedf3f','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('vestuário','https://image.shutterstock.com/image-vector/yellow-sneakers-vector-isolated-260nw-569243245.jpg','chaussure','sapato','','2647abbc-39fb-4228-826d-08ebcdac2d0d','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('ações','https://cdn.pixabay.com/photo/2013/07/12/17/00/panda-151674_960_720.png','assesoir','sentar','','5f018c0a-695c-4c17-8c12-af0f443f914b','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('sentimentos','https://image.freepik.com/free-vector/cute-boy-sleepy_33070-1825.jpg','sommeil','sono','Il a sommeil','0703f85e-8230-47fa-a3b1-ca9110d79edd','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('sentimentos','https://www.criandocomapego.com/wp-content/uploads/2018/11/fichas-das-emocoes-para-imprimir-tristeza.jpg','triste','triste','Il a triste','c292c2d4-ce8c-4832-8965-9366166b7534','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05');
INSERT INTO public.quizzes (tag,url,mot,palavra,audio,uuid,created_at,updated_at) VALUES
	 ('animais','https://image.freepik.com/free-vector/pack-baby-sharks-ocean_23-2148492976.jpg','requin','tubarão','le requin','dd16a685-bbfe-4489-b018-6fc64cd8a042','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('frutas','https://image.freepik.com/free-vector/grapes-cartoon-icon_90744-153.jpg','raisin','uva','le raisin','079f809c-9d08-4050-82a9-85e8237c2e64','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('cores','https://cdn.pixabay.com/photo/2013/07/13/10/45/bucket-157731_960_720.png','vert','verde','','53b10de0-5b16-4e10-aabb-c28bda3a7ca5','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('cores','https://www.clipartkey.com/mpngs/m/144-1449759_drop-clip-art-colours-red-splodge.png','rouge','vermelho','','b08755d3-1b77-4157-9784-b04d111b5eba','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('vestuário','https://cdn.pixabay.com/photo/2016/09/26/10/26/cute-cartoon-characters-1695612_960_720.png','robe','vestido','la robe','14275a4a-0f04-4b84-9063-33a0f495f724','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05'),
	 ('alimentos','https://image.freepik.com/free-vector/cute-smiling-happy-glass-water-drop-flat-cartoon-character-illustration-isolated-white-background-water-drop-character-concept_92289-1426.jpg','l''eau','água','','daf92611-9a02-4792-a080-fb98eaf0d9e0','2020-12-12 17:26:55-05','2020-12-12 17:27:48-05');
