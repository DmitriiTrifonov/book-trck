-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.users
(
    id serial PRIMARY KEY,
    username varchar NOT NULL CHECK (username <> ''),
    password varchar NOT NULL CHECK (password <> ''),
    created_at timestamptz NOT NULL
);

CREATE TABLE public.books (
    id serial PRIMARY KEY,
    title varchar NOT NULL CHECK (title <> ''),
    isbn varchar,
    total_pages int8,
    year int8,
    created_at timestamptz NOT NULL
);

CREATE TABLE public.authors
(
    id serial PRIMARY KEY,
    first_name varchar NOT NULL CHECK (first_name <> ''),
    last_name varchar,
    middle_name varchar,
    created_at timestamptz NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.authors;

DROP table public.books;

DROP TABLE public.users;
-- +goose StatementEnd
