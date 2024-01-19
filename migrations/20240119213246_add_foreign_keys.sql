-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.instructor_access ADD CONSTRAINT instructor_access_fk FOREIGN KEY (course_id) REFERENCES public.course(id);
ALTER TABLE public.instructor_access ADD CONSTRAINT instructor_access_instructor_fk FOREIGN KEY (instructor_id) REFERENCES public.instructor(id);

ALTER TABLE public.student_access ADD CONSTRAINT student_access_course_id_fk FOREIGN KEY (course_id) REFERENCES public.course(id);
ALTER TABLE public.student_access ADD CONSTRAINT student_access_student_id_fk FOREIGN KEY (student_id) REFERENCES public.student(id);

ALTER TABLE public."class" ADD CONSTRAINT class_fk FOREIGN KEY (course_id) REFERENCES public.course(id);

ALTER TABLE public.class_material ADD CONSTRAINT class_material_fk FOREIGN KEY (class_id) REFERENCES public."class"(id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.instructor_access DROP CONSTRAINT instructor_access_fk;
ALTER TABLE public.instructor_access DROP CONSTRAINT instructor_access_instructor_fk;

ALTER TABLE public.student_access DROP CONSTRAINT student_access_course_id_fk;
ALTER TABLE public.student_access DROP CONSTRAINT student_access_student_id_fk;

ALTER TABLE public."class" DROP CONSTRAINT class_fk;

ALTER TABLE public.class_material DROP CONSTRAINT class_material_fk;
-- +goose StatementEnd
