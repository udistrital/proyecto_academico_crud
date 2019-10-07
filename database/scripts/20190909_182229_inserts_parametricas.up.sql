-- Inserts tipo_titulación
insert into proyecto_academico.tipo_titulacion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (1, 'Titulación SNIES', 'Titulación registrada en el SNIES', 'TSNIES', true, 1, now(), now());
insert into proyecto_academico.tipo_titulacion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (2, 'Titulación Hombre', 'Titulación del proyecto para el hombre', 'THOMBRE', true, 2, now(), now()); 
insert into proyecto_academico.tipo_titulacion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (3, 'Titulación Mujer', 'Titulación del proyecto para el mujer', 'TMUJER', true, 3, now(), now()); 

-- Inserts tipo_registro
insert into proyecto_academico.tipo_registro (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (1, 'Registro de acreditación', 'Registor de acreditación', 'REGACR', true, 1, now(), now());
insert into proyecto_academico.tipo_registro (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (2, 'Renovación', 'Renovación', 'REGREN', true, 2, now(), now()); 

-- Inserts enfasis
insert into proyecto_academico.enfasis (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (1, 'Enfasis 1', 'Enfasis 1', 'ENF1', true, 1, now(), now());
insert into proyecto_academico.enfasis (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (2, 'Enfasis 2', 'Enfasis 2', 'ENF2', true, 2, now(), now()); 

-- Inserts nivel_formacion
insert into proyecto_academico.nivel_formacion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (1, 'Pregrado', 'Pregrado', 'PRE', true, 1, now(), now());
insert into proyecto_academico.nivel_formacion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (2, 'Posgrado', 'Posgrado', 'POS', true, 2, now(), now()); 

-- Inserts metodologia
insert into proyecto_academico.metodologia (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (1, 'Presencial', 'Presencial', 'PRE', true, 1, now(), now());
insert into proyecto_academico.metodologia (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion) 
	values (2, 'Virtual', 'Virtual', 'VIR', true, 2, now(), now()); 

