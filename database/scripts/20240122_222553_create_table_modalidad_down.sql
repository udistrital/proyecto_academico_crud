ALTER TABLE proyecto_academico.proyecto_academico_institucion
DROP CONSTRAINT IF EXISTS fk_modalidad_proyecto_academico_institucion CASCADE;

ALTER TABLE proyecto_academico.proyecto_academico_institucion
DROP COLUMN IF EXISTS modalidad_id;

DROP TABLE IF EXISTS proyecto_academico.modalidad CASCADE;