--
-- PostgreSQL database dump
--

-- Dumped from database version 15.7 (Debian 15.7-1.pgdg120+1)
-- Dumped by pg_dump version 16.3

-- Started on 2024-08-25 15:26:29

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 217 (class 1259 OID 16399)
-- Name: notes; Type: TABLE; Schema: public; Owner: baseuser
--

CREATE TABLE public.notes (
    id bigint NOT NULL,
    title character varying(300) NOT NULL,
    description text DEFAULT ''::text,
    user_id bigint
);


ALTER TABLE public.notes OWNER TO baseuser;

--
-- TOC entry 216 (class 1259 OID 16398)
-- Name: notes_id_seq; Type: SEQUENCE; Schema: public; Owner: baseuser
--

CREATE SEQUENCE public.notes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.notes_id_seq OWNER TO baseuser;

--
-- TOC entry 3361 (class 0 OID 0)
-- Dependencies: 216
-- Name: notes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baseuser
--

ALTER SEQUENCE public.notes_id_seq OWNED BY public.notes.id;


--
-- TOC entry 215 (class 1259 OID 16390)
-- Name: users; Type: TABLE; Schema: public; Owner: baseuser
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    login character varying(150) NOT NULL,
    password character varying(255) NOT NULL
);


ALTER TABLE public.users OWNER TO baseuser;

--
-- TOC entry 214 (class 1259 OID 16389)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: baseuser
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO baseuser;

--
-- TOC entry 3362 (class 0 OID 0)
-- Dependencies: 214
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: baseuser
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 3205 (class 2604 OID 16402)
-- Name: notes id; Type: DEFAULT; Schema: public; Owner: baseuser
--

ALTER TABLE ONLY public.notes ALTER COLUMN id SET DEFAULT nextval('public.notes_id_seq'::regclass);


--
-- TOC entry 3204 (class 2604 OID 16393)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: baseuser
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 3212 (class 2606 OID 16407)
-- Name: notes notes_pkey; Type: CONSTRAINT; Schema: public; Owner: baseuser
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_pkey PRIMARY KEY (id);


--
-- TOC entry 3208 (class 2606 OID 16414)
-- Name: users users_login_key; Type: CONSTRAINT; Schema: public; Owner: baseuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_login_key UNIQUE (login);


--
-- TOC entry 3210 (class 2606 OID 16397)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: baseuser
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- TOC entry 3213 (class 2606 OID 16408)
-- Name: notes notes_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: baseuser
--

ALTER TABLE ONLY public.notes
    ADD CONSTRAINT notes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


-- Completed on 2024-08-25 15:26:29

--
-- PostgreSQL database dump complete
--

