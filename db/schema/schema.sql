-- DROP TABLE IF EXISTS public.order_item;
-- DROP TABLE IF EXISTS public."order";

CREATE TABLE IF NOT EXISTS public."order"
(
    order_uid character varying(32) COLLATE pg_catalog."default" NOT NULL,
    track_number character varying(32) COLLATE pg_catalog."default" NOT NULL,
    entry character varying(4) COLLATE pg_catalog."default" NOT NULL,
    locale character varying(2) COLLATE pg_catalog."default" NOT NULL,
    internal_signature character varying(32) COLLATE pg_catalog."default" NOT NULL,
    customer_id character varying(32) COLLATE pg_catalog."default" NOT NULL,
    delivery_service character varying(32) COLLATE pg_catalog."default" NOT NULL,
    shardkey character varying(4) COLLATE pg_catalog."default" NOT NULL,
    sm_id integer NOT NULL,
    oof_shard character varying(4) COLLATE pg_catalog."default" NOT NULL,
    delivery_name character varying(64) COLLATE pg_catalog."default" NOT NULL,
    delivery_phone character varying(11) COLLATE pg_catalog."default" NOT NULL,
    delivery_zip character varying(7) COLLATE pg_catalog."default" NOT NULL,
    delivery_city character varying(32) COLLATE pg_catalog."default" NOT NULL,
    delivery_address character varying(128) COLLATE pg_catalog."default" NOT NULL,
    delivery_region character varying(32) COLLATE pg_catalog."default" NOT NULL,
    delivery_email character varying(64) COLLATE pg_catalog."default" NOT NULL,
    payment_transaction character varying(32) COLLATE pg_catalog."default" NOT NULL,
    payment_request_id character varying(32) COLLATE pg_catalog."default" NOT NULL,
    payment_currency character varying(3) COLLATE pg_catalog."default" NOT NULL,
    payment_provider character varying(16) COLLATE pg_catalog."default" NOT NULL,
    payment_amount integer NOT NULL,
    payment_payment_dt integer NOT NULL,
    payment_bank character varying(16) COLLATE pg_catalog."default" NOT NULL,
    payment_delivery_cost integer NOT NULL,
    payment_goods_total integer NOT NULL,
    payment_custom_fee integer NOT NULL,
    date_created timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT order_pkey PRIMARY KEY (order_uid)
);

CREATE TABLE IF NOT EXISTS public.order_item
(
    order_uid character varying(32) COLLATE pg_catalog."default" NOT NULL,
    chrt_id integer NOT NULL,
    track_number character varying(32) COLLATE pg_catalog."default" NOT NULL,
    price integer NOT NULL,
    rid character varying(32) COLLATE pg_catalog."default" NOT NULL,
    name character varying(32) COLLATE pg_catalog."default" NOT NULL,
    sale integer NOT NULL,
    size character varying(6) COLLATE pg_catalog."default" NOT NULL,
    total_price integer NOT NULL,
    nm_id integer NOT NULL,
    brand character varying(32) COLLATE pg_catalog."default" NOT NULL,
    status integer NOT NULL,
    CONSTRAINT order_item_pkey PRIMARY KEY (order_uid, chrt_id)
);
