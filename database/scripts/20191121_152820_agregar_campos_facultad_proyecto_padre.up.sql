-- Campo facultad_id
ALTER TABLE proyecto_academico.proyecto_academico_institucion ADD COLUMN facultad_id integer NOT NULL DEFAULT 0;

-- Relación proyecto padre
ALTER TABLE proyecto_academico.proyecto_academico_institucion ADD COLUMN proyecto_padre_id integer;


-- object: fk_proyecto_padre_proyecto_academico_institucion | type: CONSTRAINT --
-- ALTER TABLE proyecto_academico.proyecto_academico_institucion DROP CONSTRAINT IF EXISTS fk_proyecto_padre_proyecto_academico_institucion CASCADE;
ALTER TABLE proyecto_academico.proyecto_academico_institucion ADD CONSTRAINT fk_proyecto_padre_proyecto_academico_institucion FOREIGN KEY (proyecto_padre_id)
REFERENCES proyecto_academico.proyecto_academico_institucion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;