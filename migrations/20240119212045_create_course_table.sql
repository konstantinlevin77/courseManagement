-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.course (
                               id serial4 NOT NULL,
                               "name" varchar NOT NULL,
                               description text NOT NULL,
                               created_at timestamp NOT NULL,
                               updated_at timestamp NOT NULL,
                               CONSTRAINT course_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.course;
-- +goose StatementEnd
