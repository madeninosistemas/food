-- Active: 1697408358107@@192.168.1.147@5432@DBDelivery@public
-- Active: 1697408358107@@192.168.1.135@5432@DBDelivery@public

CREATE TABLE mesa(
	id				UUID			PRIMARY KEY,
	nombre			VARCHAR(50)		NOT NULL,
	url				VARCHAR(200)	NOT NULL,
	images			VARCHAR(250)	NOT NULL,
	activo			BOOLEAN			NOT NULL DEFAULT TRUE,
	created_at		INTEGER			NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	updated_at		INTEGER
);

CREATE TABLE users(
	id 		    	UUID 			DEFAULT uuid_generate_v4() PRIMARY KEY,
	nombre			VARCHAR(100)	NOT NULL,
	email			VARCHAR(254)	NOT NULL,
	password		VARCHAR(100)	NOT NULL,
	is_admin		BOOLEAN 		NOT NULL DEFAULT TRUE,
	images			JSONB			NOT NULL,
	details	    	JSONB 	    	NOT NULL,
	created_at		INTEGER			NOT NULL  DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
	updated_at 		INTEGER,
	CONSTRAINT users_email_uk UNIQUE (email)
);

CREATE TABLE Promocion(
	id				UUID			PRIMARY KEY,
	nombre			VARCHAR(100)	NOT NULL,
	description		TEXT			NOT NULL,
	image			VARCHAR(256)	NOT NULL,
	precio			NUMERIC(10,2)	NOT NULL,
	features		JSONB			NOT NULL,
	activo			BOOLEAN			NOT NULL,
	created			INTEGER			NOT NULL,
	updated			INTEGER			NOT NULL
);


CREATE TABLE Category(
    id              UUID            PRIMARY KEY,
    nombre		    VARCHAR(100)    NOT NULL UNIQUE,
    description     TEXT            NOT NULL,
	images			VARCHAR(250)	NOT NULL,
    activo          BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at      INTEGER         NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at      INTEGER
);

CREATE TABLE Productos (
    id            UUID            PRIMARY KEY,
    idcategoria  UUID            NOT NULL,
    nombre        VARCHAR(128)    NOT NULL UNIQUE,
    precio        NUMERIC(10,2)   NOT NULL,
    imagen        VARCHAR(250)    NOT NULL,
    descripcion   TEXT            NOT NULL,
    activo        BOOLEAN         NOT NULL DEFAULT TRUE,
    time          TIMESTAMP       NOT NULL,
    calorias      NUMERIC(6,2)    NOT NULL,
    created_at    INTEGER         NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::INT,
    updated_at    INTEGER,
    FOREIGN KEY(idcategoria) REFERENCES Category(id)
);
