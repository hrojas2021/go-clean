CREATE TABLE IF NOT EXISTS users (
   id uuid NOT NULL,
   name VARCHAR (128),
   username VARCHAR (128),
   password VARCHAR (128),
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   PRIMARY KEY (id)
);