-- Eliminar relación
ALTER TABLE proyecto_academico.nivel_formacion DROP CONSTRAINT IF EXISTS fk_nivel_formacion_padre_nivel_formacion CASCADE;

-- Eliminar campos
ALTER TABLE proyecto_academico.nivel_formacion
DROP COLUMN IF EXISTS nivel_formacion_padre_id;
