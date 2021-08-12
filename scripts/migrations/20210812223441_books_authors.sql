-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.books_authors (
    book_id int8 NOT NULL,
    author_id int8 NOT NULL,
    user_id int8 NOT NULL,
    added_at timestamptz NOT NULL,
    FOREIGN KEY (author_id) REFERENCES public.authors (id),
    FOREIGN KEY (book_id) REFERENCES public.books (id),
    FOREIGN KEY (user_id) REFERENCES public.users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.books_authors;
-- +goose StatementEnd
