-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.1
-- PostgreSQL version: 9.5
-- Project Site: pgmodeler.io
-- Model Author: ---

-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: proyecto_academico | type: DATABASE --
-- -- DROP DATABASE IF EXISTS proyecto_academico;
-- CREATE DATABASE proyecto_academico;
-- -- ddl-end --
-- 

-- object: proyecto_academico | type: SCHEMA --
-- DROP SCHEMA IF EXISTS proyecto_academico CASCADE;
CREATE SCHEMA proyecto_academico;
-- ddl-end --

-- SET search_path TO pg_catalog,public,proyecto_academico;
-- ddl-end --

-- object: proyecto_academico.proyecto_academico_institucion | type: TABLE --
-- DROP TABLE IF EXISTS proyecto_academico.proyecto_academico_institucion CASCADE;
CREATE TABLE proyecto_academico.proyecto_academico_institucion(
	id serial NOT NULL,
	codigo varchar(6) NOT NULL,
	nombre character varying(250) NOT NULL,
	codigo_snies varchar(10) NOT NULL,
	duracion numeric(3,0) NOT NULL,
	correo_electronico varchar(30) NOT NULL,
	numero_creditos numeric(4) NOT NULL,
	ciclos_propedeuticos boolean NOT NULL,
	numero_acto_administrativo numeric(6) NOT NULL,
	enlace_acto_administrativo text NOT NULL,
	competencias text NOT NULL,
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	unidad_tiempo_id integer NOT NULL,
	ano_acto_administrativo varchar(4) NOT NULL,
	dependencia_id integer NOT NULL,
	area_conocimiento_id integer NOT NULL,
	nucleo_base_id integer NOT NULL,
	metodologia_id integer,
	nivel_formacion_id integer,
	CONSTRAINT pk_programa_academico PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE proyecto_academico.proyecto_academico_institucion IS 'Programas acadÈmicos de las intituciones de educaciÛn regstrados en el SNIES, SIET u otros (cuando es educaciÛn informal)';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.id IS 'Identificadorde la tabla programa_academico';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.codigo IS 'Campo para el codigo interno del proyecto curricular ';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.nombre IS 'Nombre del programa acadÈmico';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.codigo_snies IS 'Corresponde al cÛdigo del registro en el SNIES o el SIET(para los programas de Educacion para el Trabajo y el Desarrollo Humano), para los programas de educaciÛn informal el cÛdigo es 0';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.duracion IS 'DuraciÛn del programa acadÈmico';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.correo_electronico IS 'Campo para almacenar el correo elÈctronico del proyecto curricular ';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.numero_creditos IS 'Campo para almacenar el numero de credistos del proyecto curricular';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.ciclos_propedeuticos IS 'Campo para almacenar si el proyecto curricular es desarrollado por ciclos_propedeuticos';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.numero_acto_administrativo IS 'Campo para almacenar el nuemero del acuerdo del concejo superior';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.enlace_acto_administrativo IS 'Campo para registrar el enlace del documento de acto administrativo';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.competencias IS 'Campo para almacenar las compentencias del proyecto curricular ';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.codigo_abreviacion IS 'Campo para guardar la abreviaciÛn  del proyecto curricular ';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.activo IS 'Indica el estado del programa';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.fecha_creacion IS 'Campo para el registro de la fecha de creaciÛn del registro';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.fecha_modificacion IS 'Campo para el registro de la fecha de modificacion del registro';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.unidad_tiempo_id IS 'Unidad de tiempo para la duraciÛn la formaciÛn acadÈmica (ej: AÒos, Semestres). Referencia a la tabla unidad_tiempo';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.ano_acto_administrativo IS 'Campo que almacena el aÒo en el que se efectuo la resoluciÛn del proyecto curricular por parte del concejo superior';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.dependencia_id IS 'Campo que referencia la dependencia del proyecto curricular como facultad ';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.area_conocimiento_id IS 'Campo que referencia al esquema del core para el area de conocimiento defino por el SNIES';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_institucion.nucleo_base_id IS 'Campo que referencia el nucleo base conocimiento del esquema core segun el SNIES';
-- ddl-end --
COMMENT ON CONSTRAINT pk_programa_academico ON proyecto_academico.proyecto_academico_institucion  IS 'Llave primaria de la tabla programa_academico';
-- ddl-end --

-- object: proyecto_academico.enfasis | type: TABLE --
-- DROP TABLE IF EXISTS proyecto_academico.enfasis CASCADE;
CREATE TABLE proyecto_academico.enfasis(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_enfasis PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE proyecto_academico.enfasis IS '…nfasis de las carreras';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.enfasis.fecha_creacion IS 'Fecha de creacion de un enfasis';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.enfasis.fecha_modificacion IS 'Fecha de modificacion de un enfasis';
-- ddl-end --

-- object: proyecto_academico.metodologia | type: TABLE --
-- DROP TABLE IF EXISTS proyecto_academico.metodologia CASCADE;
CREATE TABLE proyecto_academico.metodologia(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_metodologia PRIMARY KEY (id),
	CONSTRAINT uq_nombre_metodologia UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE proyecto_academico.metodologia IS 'Tabla que almacena las metodologÌas para un programa acadÈmico';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.metodologia.id IS 'Identificador de la tabla paramÈtrica metodologia';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.metodologia.nombre IS 'Nombre de la metodologÌa (ej: presencial, virtual, a distancia)';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.metodologia.descripcion IS 'DescripciÛn opcional del par·metro';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.metodologia.codigo_abreviacion IS 'CÛdigo de abreviaciÛn, sigla, acrÛnimo u otra representaciÛn corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.metodologia.activo IS 'Campo que indica si el par·metro est· activo';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.metodologia.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicaciÛn ni por BD. AllÌ se almacena permitiendo realizar subitems mediante la precisiÛn.';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.metodologia.fecha_creacion IS 'Fecha de creacion de una metodologia';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.metodologia.fecha_modificacion IS 'Fecha de modificacion de una metodologia ';
-- ddl-end --
COMMENT ON CONSTRAINT pk_metodologia ON proyecto_academico.metodologia  IS 'Llave primaria de la tabla metodologia';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_metodologia ON proyecto_academico.metodologia  IS 'RestricciÛn para que no se repita el nombre de las metodologÌas';
-- ddl-end --

-- object: proyecto_academico.nivel_formacion | type: TABLE --
-- DROP TABLE IF EXISTS proyecto_academico.nivel_formacion CASCADE;
CREATE TABLE proyecto_academico.nivel_formacion(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_nivel_formacion PRIMARY KEY (id),
	CONSTRAINT uq_nombre_nivel_formacion UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE proyecto_academico.nivel_formacion IS 'Tabla paramÈtrica con los niveles de formaciÛn';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.nivel_formacion.id IS 'Identificador de la tabla nivel_formacion';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.nivel_formacion.nombre IS 'Nombre del nivel de formaciÛn (ej: TecnolÛgica, Universitaria, EspecializaciÛn)';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.nivel_formacion.descripcion IS 'DescripciÛn opcional del par·metro';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.nivel_formacion.codigo_abreviacion IS 'CÛdigo de abreviaciÛn, sigla, acrÛnimo u otra representaciÛn corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.nivel_formacion.activo IS 'Campo que indica si el par·metro est· activo';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.nivel_formacion.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicaciÛn ni por BD. AllÌ se almacena permitiendo realizar subitems mediante la precisiÛn.';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.nivel_formacion.fecha_creacion IS 'Fecha de creacion de un nivel de formaciÛn';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.nivel_formacion.fecha_modificacion IS 'Fecha de modificacion de un nivel de formaciÛn';
-- ddl-end --
COMMENT ON CONSTRAINT pk_nivel_formacion ON proyecto_academico.nivel_formacion  IS 'Llave primaria de la tabla nivel_formacion';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_nivel_formacion ON proyecto_academico.nivel_formacion  IS 'RestricciÛn para que no se repita el nombre de los niveles de formaciÛn';
-- ddl-end --

-- object: proyecto_academico.titulacion | type: TABLE --
-- DROP TABLE IF EXISTS proyecto_academico.titulacion CASCADE;
CREATE TABLE proyecto_academico.titulacion(
	id serial NOT NULL,
	nombre character varying(250) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	tipo_titulacion_id integer,
	proyecto_academico_institucion_id integer,
	CONSTRAINT pk_titulacion PRIMARY KEY (id),
	CONSTRAINT uq_nombre_titulacion UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE proyecto_academico.titulacion IS 'Tabla paramÈtrica con las titulaciones de los programas acadÈmicos';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.titulacion.id IS 'Identificador de la tabla titulacion';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.titulacion.nombre IS 'Nombre del tÌtulo otorgado por la instituciÛn';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.titulacion.descripcion IS 'DescripciÛn opcional del par·metro';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.titulacion.codigo_abreviacion IS 'CÛdigo de abreviaciÛn, sigla, acrÛnimo u otra representaciÛn corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.titulacion.activo IS 'Campo que indica si el par·metro est· activo';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.titulacion.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicaciÛn ni por BD. AllÌ se almacena permitiendo realizar subitems mediante la precisiÛn.';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.titulacion.fecha_creacion IS 'Fecha de creacion de una titulacion';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.titulacion.fecha_modificacion IS 'Fecha de modificacion de titulacion';
-- ddl-end --
COMMENT ON CONSTRAINT pk_titulacion ON proyecto_academico.titulacion  IS 'Llave primaria de la tabla titulacion';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_titulacion ON proyecto_academico.titulacion  IS 'RestricciÛn para que no se repita el nombre de las titulaciones';
-- ddl-end --

-- object: fk_metodologia_proyecto_academico_institucion | type: CONSTRAINT --
-- ALTER TABLE proyecto_academico.proyecto_academico_institucion DROP CONSTRAINT IF EXISTS fk_metodologia_proyecto_academico_institucion CASCADE;
ALTER TABLE proyecto_academico.proyecto_academico_institucion ADD CONSTRAINT fk_metodologia_proyecto_academico_institucion FOREIGN KEY (metodologia_id)
REFERENCES proyecto_academico.metodologia (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_nivel_formacion_proyecto_academico_institucion | type: CONSTRAINT --
-- ALTER TABLE proyecto_academico.proyecto_academico_institucion DROP CONSTRAINT IF EXISTS fk_nivel_formacion_proyecto_academico_institucion CASCADE;
ALTER TABLE proyecto_academico.proyecto_academico_institucion ADD CONSTRAINT fk_nivel_formacion_proyecto_academico_institucion FOREIGN KEY (nivel_formacion_id)
REFERENCES proyecto_academico.nivel_formacion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: proyecto_academico.proyecto_academico_enfasis | type: TABLE --
-- DROP TABLE IF EXISTS proyecto_academico.proyecto_academico_enfasis CASCADE;
CREATE TABLE proyecto_academico.proyecto_academico_enfasis(
	id serial NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	proyecto_academico_institucion_id integer,
	enfasis_id integer,
	CONSTRAINT pk_proyecto_academico_enfasis PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_enfasis.activo IS 'Campo para el registro si esta activo ';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_enfasis.fecha_creacion IS 'Fecha de creacion de la tabla ';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_enfasis.fecha_modificacion IS 'Fecha de modificacion de la tabla';
-- ddl-end --

-- object: fk_proyecto_academico_institucion_proyecto_academico__4404 | type: CONSTRAINT --
-- ALTER TABLE proyecto_academico.proyecto_academico_enfasis DROP CONSTRAINT IF EXISTS fk_proyecto_academico_institucion_proyecto_academico__4404 CASCADE;
ALTER TABLE proyecto_academico.proyecto_academico_enfasis ADD CONSTRAINT fk_proyecto_academico_institucion_proyecto_academico__4404 FOREIGN KEY (proyecto_academico_institucion_id)
REFERENCES proyecto_academico.proyecto_academico_institucion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_enfasis_proyecto_academico_enfasis | type: CONSTRAINT --
-- ALTER TABLE proyecto_academico.proyecto_academico_enfasis DROP CONSTRAINT IF EXISTS fk_enfasis_proyecto_academico_enfasis CASCADE;
ALTER TABLE proyecto_academico.proyecto_academico_enfasis ADD CONSTRAINT fk_enfasis_proyecto_academico_enfasis FOREIGN KEY (enfasis_id)
REFERENCES proyecto_academico.enfasis (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: proyecto_academico.registro_calificado_acreditacion | type: TABLE --
-- DROP TABLE IF EXISTS proyecto_academico.registro_calificado_acreditacion CASCADE;
CREATE TABLE proyecto_academico.registro_calificado_acreditacion(
	id serial NOT NULL,
	numero_acto_administrativo numeric(8) NOT NULL,
	ano_acto_administrativo_id varchar(4) NOT NULL,
	fecha_creacion_acto_administrativo timestamp NOT NULL,
	vigencia_acto_administrativo text NOT NULL,
	vencimiento_acto_administrativo timestamp NOT NULL,
	enlace_acto text NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	proyecto_academico_institucion_id integer,
	tipo_registro_id integer,
	CONSTRAINT pk_registro_calificado_alta_calidad PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE proyecto_academico.registro_calificado_acreditacion IS 'Tabla para el registro de la informaciÛn de los registro califado y los registros de alta calidad de los proyectos curriculares';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.registro_calificado_acreditacion.id IS 'Campo para el identificador del registro';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.registro_calificado_acreditacion.numero_acto_administrativo IS 'Campo para el registro del numero de resoluciÛn del registro calificado';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.registro_calificado_acreditacion.ano_acto_administrativo_id IS 'Campo que registra el aÒo de la resoluciÛn del registro calificado';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.registro_calificado_acreditacion.fecha_creacion_acto_administrativo IS 'Campo para el registro de la fecha de creaciÛn de la resoluciÛn del registro calificado';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.registro_calificado_acreditacion.vigencia_acto_administrativo IS 'Campo para el registro de la vigencia del registro calificado, se ingresa una fecha en meses y aÒos el dia se definie por cliente en 0.';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.registro_calificado_acreditacion.vencimiento_acto_administrativo IS 'Campo para el registro de la fecha de vencimiento del registro calificado, este se calcula con los campos de fecha_creacion_registro_calificado y el campo vigencia_registro_calificado';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.registro_calificado_acreditacion.enlace_acto IS 'Campo para registrar el enlace del documento de acto administrativo';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.registro_calificado_acreditacion.activo IS 'Campo para el registro activo';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.registro_calificado_acreditacion.fecha_creacion IS 'Fecha de creacion de un registro';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.registro_calificado_acreditacion.fecha_modificacion IS 'Fecha de modificacion de registro';
-- ddl-end --

-- object: fk_proyecto_academico_institucion_registro_calificado_ac_4408 | type: CONSTRAINT --
-- ALTER TABLE proyecto_academico.registro_calificado_acreditacion DROP CONSTRAINT IF EXISTS fk_proyecto_academico_institucion_registro_calificado_ac_4408 CASCADE;
ALTER TABLE proyecto_academico.registro_calificado_acreditacion ADD CONSTRAINT fk_proyecto_academico_institucion_registro_calificado_ac_4408 FOREIGN KEY (proyecto_academico_institucion_id)
REFERENCES proyecto_academico.proyecto_academico_institucion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: proyecto_academico.tipo_registro | type: TABLE --
-- DROP TABLE IF EXISTS proyecto_academico.tipo_registro CASCADE;
CREATE TABLE proyecto_academico.tipo_registro(
	id serial NOT NULL,
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_registro PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE proyecto_academico.tipo_registro IS 'Tabla parametrica del tipo de registro,  acreditaciÛn, renovaciÛn, alta calidad';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.tipo_registro.fecha_creacion IS 'Fecha de creacion de un tipo de registro';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.tipo_registro.fecha_modificacion IS 'Fecha de modificacion de un tipo de registro';
-- ddl-end --

-- object: proyecto_academico.proyecto_academico_rol_persona_dependecia | type: TABLE --
-- DROP TABLE IF EXISTS proyecto_academico.proyecto_academico_rol_persona_dependecia CASCADE;
CREATE TABLE proyecto_academico.proyecto_academico_rol_persona_dependecia(
	id serial NOT NULL,
	persona_id integer NOT NULL,
	dependencia_id integer NOT NULL,
	rol_id integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	proyecto_academico_institucion_id integer,
	CONSTRAINT pk_proyecto_academico_dependecia PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE proyecto_academico.proyecto_academico_rol_persona_dependecia IS 'Tabla de rompiendo entre el proyecto academico y el esquema de personas, tambien se asocia el rol de la persona I.E coordinador , asistente';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_rol_persona_dependecia.id IS 'campo para el registro del consecutivo de los registros';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_rol_persona_dependecia.persona_id IS 'Campo que referencia a la tabla personas del esquema core_new';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_rol_persona_dependecia.dependencia_id IS 'Campo para relacionar con la tabla dependencia en el esquema espacio_logico, relaciona la depencia (proyecto curricular) con la vinculaciÛn de la persona.';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_rol_persona_dependecia.rol_id IS 'Campo que referencia el rol que tiene la persona en WSO2';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_rol_persona_dependecia.activo IS 'Campo que indica si el par·metro est· activo';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_rol_persona_dependecia.fecha_creacion IS 'Fecha de creacion del registro';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.proyecto_academico_rol_persona_dependecia.fecha_modificacion IS 'Fecha de modificaciÛn de un registro';
-- ddl-end --

-- object: fk_tipo_registro_registro_calificado_acreditacion | type: CONSTRAINT --
-- ALTER TABLE proyecto_academico.registro_calificado_acreditacion DROP CONSTRAINT IF EXISTS fk_tipo_registro_registro_calificado_acreditacion CASCADE;
ALTER TABLE proyecto_academico.registro_calificado_acreditacion ADD CONSTRAINT fk_tipo_registro_registro_calificado_acreditacion FOREIGN KEY (tipo_registro_id)
REFERENCES proyecto_academico.tipo_registro (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: proyecto_academico.tipo_titulacion | type: TABLE --
-- DROP TABLE IF EXISTS proyecto_academico.tipo_titulacion CASCADE;
CREATE TABLE proyecto_academico.tipo_titulacion(
	id serial NOT NULL,
	nombre character varying(250) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_titulacion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE proyecto_academico.tipo_titulacion IS 'Tabla paramÈtrica con las tipos de titulaciones de los programas acadÈmicos ejemplo SNIES, Genero';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.tipo_titulacion.id IS 'Identificador de la tabla tipo titulacion';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.tipo_titulacion.descripcion IS 'DescripciÛn opcional del par·metro';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.tipo_titulacion.codigo_abreviacion IS 'CÛdigo de abreviaciÛn, sigla, acrÛnimo u otra representaciÛn corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.tipo_titulacion.activo IS 'Campo que indica si el par·metro est· activo';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.tipo_titulacion.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicaciÛn ni por BD. AllÌ se almacena permitiendo realizar subitems mediante la precisiÛn.';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.tipo_titulacion.fecha_creacion IS 'Fecha de creacion de una tipo titulacion';
-- ddl-end --
COMMENT ON COLUMN proyecto_academico.tipo_titulacion.fecha_modificacion IS 'Fecha de modificacion de  tipo titulacion';
-- ddl-end --

-- object: fk_tipo_titulacion_titulacion | type: CONSTRAINT --
-- ALTER TABLE proyecto_academico.titulacion DROP CONSTRAINT IF EXISTS fk_tipo_titulacion_titulacion CASCADE;
ALTER TABLE proyecto_academico.titulacion ADD CONSTRAINT fk_tipo_titulacion_titulacion FOREIGN KEY (tipo_titulacion_id)
REFERENCES proyecto_academico.tipo_titulacion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_proyecto_academico_institucion_proyecto_academico_rol_4414 | type: CONSTRAINT --
-- ALTER TABLE proyecto_academico.proyecto_academico_rol_persona_dependecia DROP CONSTRAINT IF EXISTS fk_proyecto_academico_institucion_proyecto_academico_rol_4414 CASCADE;
ALTER TABLE proyecto_academico.proyecto_academico_rol_persona_dependecia ADD CONSTRAINT fk_proyecto_academico_institucion_proyecto_academico_rol_4414 FOREIGN KEY (proyecto_academico_institucion_id)
REFERENCES proyecto_academico.proyecto_academico_institucion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_proyecto_academico_institucion_titulacion | type: CONSTRAINT --
-- ALTER TABLE proyecto_academico.titulacion DROP CONSTRAINT IF EXISTS fk_proyecto_academico_institucion_titulacion CASCADE;
ALTER TABLE proyecto_academico.titulacion ADD CONSTRAINT fk_proyecto_academico_institucion_titulacion FOREIGN KEY (proyecto_academico_institucion_id)
REFERENCES proyecto_academico.proyecto_academico_institucion (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- Permisos de usuario
GRANT USAGE ON SCHEMA proyecto_academico TO test;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA proyecto_academico TO test;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA proyecto_academico TO test;

