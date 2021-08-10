-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.users
(
    id serial PRIMARY KEY,
    username varchar NOT NULL CHECK (username <> ''),
    password varchar NOT NULL CHECK (password <> '')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.users;
-- +goose StatementEnd
