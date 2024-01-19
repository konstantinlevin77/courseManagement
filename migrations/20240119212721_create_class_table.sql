-- +goose Up
-- +goose StatementBegin
CREATE TABLE public."class" (
                                id serial NOT NULL,
                                course_id int NOT NULL,
                                "name" varchar NOT NULL,
                                "date" date NOT NULL,
                                description text NOT NULL,
                                created_at time NOT NULL,
                                updated_at time NOT NULL,
                                CONSTRAINT class_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.class;
-- +goose StatementEnd
