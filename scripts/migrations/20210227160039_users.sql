-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.users
(
    id int8 PRIMARY KEY,
    username varchar NOT NULL CHECK (username <> ''),
    books_ids int8[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.users;
-- +goose StatementEnd
