-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.student_access (
                                       id serial NOT NULL,
                                       student_id int NOT NULL,
                                       course_id int NOT NULL,
                                       created_at timestamp NOT NULL,
                                       updated_at timestamp NOT NULL,
                                       CONSTRAINT student_access_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.student_access;
-- +goose StatementEnd
