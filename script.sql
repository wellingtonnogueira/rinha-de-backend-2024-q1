CREATE table public.client (
	id integer NOT NULL,
	"name" varchar(256) NOT NULL,
	"limit" integer NOT NULL,
	CONSTRAINT client_pk PRIMARY KEY (id)
);

CREATE TABLE public.balance (
	client_id integer NOT NULL,
	total_balance integer NOT NULL,
	CONSTRAINT balance_pk PRIMARY KEY (client_id),
	CONSTRAINT balance_client_fk FOREIGN KEY (client_id) REFERENCES public.client(id)
);

CREATE TABLE public."statement" (
	id integer NOT NULL,
	client_id integer NOT NULL,
	statement_type varchar(1) NOT NULL,
	value integer NOT NULL,
	description varchar(256) NULL,
	CONSTRAINT statement_pk PRIMARY KEY (id),
	CONSTRAINT statement_check CHECK (statement_type in ('c', 'd'))
);

INSERT INTO public.client
(id, "name", "limit") values 
(1, 'Cliente 1', 100000),
(2, 'Cliente II', 80000),
(3, 'Cliente 3-III', 1000000),
(4, 'Cliente 2+2', 10000000),
(5, 'Cliente V', 500000);

insert into public.balance 
(client_id, total_balance) values
(1, 0),
(2, 0),
(3, 0),
(4, 0),
(5, 0);

commit;


INSERT INTO public."statement"
(id, client_id, statement_type, value, description)
VALUES(1, 1, 'c', 110, 'teste');

truncate "statement" ;
