-- Campo relacion nivel formacion padre
ALTER TABLE proyecto_academico.nivel_formacion ADD COLUMN nivel_formacion_padre_id integer;


-- object: fk_nivel_formacion_padre_nivel_formacion | type: CONSTRAINT --
-- ALTER TABLE proyecto_academico.nivel_formacion DROP CONSTRAINT IF EXISTS fk_nivel_formacion_padre_nivel_formacion CASCADE;
ALTER TABLE proyecto_academico.nivel_formacion ADD CONSTRAINT fk_nivel_formacion_padre_nivel_formacion FOREIGN KEY (nivel_formacion_padre_id)
REFERENCES proyecto_academico.nivel_formacion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;