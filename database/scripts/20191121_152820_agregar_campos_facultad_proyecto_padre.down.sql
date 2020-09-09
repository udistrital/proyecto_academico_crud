-- Eliminar relación
ALTER TABLE proyecto_academico.proyecto_academico_institucion DROP CONSTRAINT IF EXISTS fk_proyecto_padre_proyecto_academico_institucion CASCADE;

-- Eliminar campos
ALTER TABLE proyecto_academico.proyecto_academico_institucion 
DROP COLUMN IF EXISTS proyecto_padre_id;

ALTER TABLE proyecto_academico.proyecto_academico_institucion 
DROP COLUMN IF EXISTS facultad_id;