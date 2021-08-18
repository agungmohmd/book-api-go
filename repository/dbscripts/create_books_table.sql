CREATE TABLE public.books
(
    id integer NOT NULL DEFAULT nextval('books_id_seq'::regclass),
    bookname character varying COLLATE pg_catalog."default" NOT NULL,
    stock numeric NOT NULL
)