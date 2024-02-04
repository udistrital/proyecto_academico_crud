CREATE TABLE proyecto_academico.modalidad(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp,
	CONSTRAINT pk_modalidad PRIMARY KEY (id),
	CONSTRAINT uq_nombre_modalidad UNIQUE (nombre)

);

COMMENT ON TABLE proyecto_academico.modalidad IS 'Tabla que almacena las modalidades para un programa académico';

COMMENT ON COLUMN proyecto_academico.modalidad.id IS 'Identificador de la tabla paramétrica modalidad';

COMMENT ON COLUMN proyecto_academico.modalidad.nombre IS 'Nombre de la modalidad (ej: investigacion, profundizacion)';

COMMENT ON COLUMN proyecto_academico.modalidad.descripcion IS 'Descripción opcional del parámetro';

COMMENT ON COLUMN proyecto_academico.modalidad.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';

COMMENT ON COLUMN proyecto_academico.modalidad.activo IS 'Campo que indica si el parámetro está activo';

COMMENT ON COLUMN proyecto_academico.modalidad.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';

COMMENT ON COLUMN proyecto_academico.modalidad.fecha_creacion IS 'Fecha de creacion de una modalidad';

COMMENT ON COLUMN proyecto_academico.modalidad.fecha_modificacion IS 'Fecha de modificacion de una modalidad ';

COMMENT ON CONSTRAINT pk_modalidad ON proyecto_academico.modalidad  IS 'Llave primaria de la tabla modalidad';

COMMENT ON CONSTRAINT uq_nombre_modalidad ON proyecto_academico.modalidad  IS 'Restricción para que no se repita el nombre de las modalidades';

ALTER TABLE proyecto_academico.proyecto_academico_institucion
ADD COLUMN modalidad_id integer;

ALTER TABLE proyecto_academico.proyecto_academico_institucion ADD CONSTRAINT fk_modalidad_proyecto_academico_institucion FOREIGN KEY (modalidad_id)
REFERENCES proyecto_academico.modalidad (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;