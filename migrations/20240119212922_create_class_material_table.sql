-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.class_material (
                                       id serial4 NOT NULL,
                                       class_id int4 NOT NULL,
                                       "type" varchar NOT NULL,
                                       description varchar NOT NULL,
                                       "path" varchar NOT NULL,
                                       created_at time NOT NULL,
                                       updated_at time NOT NULL,
                                       CONSTRAINT class_material_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.class_material;
-- +goose StatementEnd
