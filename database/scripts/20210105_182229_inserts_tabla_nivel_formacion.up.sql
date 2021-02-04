-- Inserts nivel_formacion
insert into proyecto_academico.nivel_formacion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion, nivel_formacion_padre_id)
values (3, 'tecnico', 'tecnico', 'TEC', true, 3, now(), now(),1);
insert into proyecto_academico.nivel_formacion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion, nivel_formacion_padre_id)
values (4, 'tecnologo', 'tecnologo', 'TECNO', true, 4, now(), now(),1);
insert into proyecto_academico.nivel_formacion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion, nivel_formacion_padre_id)
values (5, 'profesional', 'profesional', 'PRO', true, 5, now(), now(),1);
insert into proyecto_academico.nivel_formacion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion, nivel_formacion_padre_id)
values (6, 'especializacion', 'especializacion', 'ESP', true, 6, now(), now(),2);
insert into proyecto_academico.nivel_formacion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion, nivel_formacion_padre_id)
values (7, 'maestria', 'maestria', 'MAS', true, 7, now(), now(),2);
insert into proyecto_academico.nivel_formacion (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden, fecha_creacion, fecha_modificacion, nivel_formacion_padre_id)
values (8, 'doctorado', 'doctorado', 'DOC', true, 8, now(), now(),2);