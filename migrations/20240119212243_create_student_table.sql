-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.student (
                                id serial4 NOT NULL,
                                first_name varchar NOT NULL,
                                last_name varchar NOT NULL,
                                email varchar NOT NULL,
                                phone_number varchar NOT NULL,
                                "password" varchar NOT NULL,
                                created_at time NOT NULL,
                                updated_at time NOT NULL,
                                CONSTRAINT student_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.student;
-- +goose StatementEnd
