-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.instructor (
                                   id serial4 NOT NULL,
                                   first_name varchar NOT NULL,
                                   last_name varchar NOT NULL,
                                   email varchar NOT NULL,
                                   "password" varchar NOT NULL,
                                   phone_number varchar NOT NULL,
                                   created_at time NOT NULL,
                                   updated_at time NOT NULL,
                                   CONSTRAINT instructor_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.instructor
-- +goose StatementEnd
