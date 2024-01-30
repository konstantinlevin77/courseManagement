-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.instructor_access (
                                          id serial4 NOT NULL,
                                          instructor_id int4 NOT NULL,
                                          course_id int4 NOT NULL,
                                          created_at timestamp NOT NULL,
                                          updated_at timestamp NOT NULL,
                                          CONSTRAINT instructor_access_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.instructor_access;
-- +goose StatementEnd
