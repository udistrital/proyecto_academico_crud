-- Nombre de la tabla
ALTER TABLE proyecto_academico.proyecto_academico_rol_persona_dependecia
    RENAME TO proyecto_academico_rol_tercero_dependencia;

-- Eliminación de atributo rol_id
ALTER TABLE proyecto_academico.proyecto_academico_rol_tercero_dependencia DROP COLUMN IF EXISTS rol_id;

-- Ajuste de persona_id por tercero_id
ALTER TABLE proyecto_academico.proyecto_academico_rol_tercero_dependencia
    RENAME COLUMN persona_id TO tercero_id;

-- Actualización de constraints
ALTER TABLE proyecto_academico.proyecto_academico_rol_tercero_dependencia
    RENAME CONSTRAINT pk_proyecto_academico_rol_persona_dependecia TO pk_proyecto_academico_rol_tercero_dependencia;

-- Actualización de secuencia
ALTER SEQUENCE proyecto_academico.proyecto_academico_rol_persona_dependecia_id_seq RENAME TO proyecto_academico_rol_tercero_dependencia_id_seq;
