-- Nombre de la tabla
ALTER TABLE proyecto_academico.proyecto_academico_rol_tercero_dependencia
    RENAME TO proyecto_academico_rol_persona_dependecia;

-- Agregar atributo rol_id
ALTER TABLE proyecto_academico.proyecto_academico_rol_persona_dependecia ADD COLUMN rol_id integer NOT NULL;

COMMENT ON COLUMN proyecto_academico.proyecto_academico_rol_persona_dependecia.rol_id IS 'Campo que referencia el rol que tiene la persona en WSO2';

-- Ajuste de tercero_id por persona_id
ALTER TABLE proyecto_academico.proyecto_academico_rol_persona_dependecia
    RENAME COLUMN tercero_id TO persona_id;

-- Actualización de constraints
ALTER TABLE proyecto_academico.proyecto_academico_rol_persona_dependecia
    RENAME CONSTRAINT pk_proyecto_academico_rol_tercero_dependencia TO pk_proyecto_academico_rol_persona_dependecia;

-- Actualización de secuencia
ALTER SEQUENCE proyecto_academico.proyecto_academico_rol_tercero_dependencia_id_seq RENAME TO proyecto_academico_rol_persona_dependecia_id_seq;