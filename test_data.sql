DO $$ DECLARE
    r RECORD;
BEGIN
    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
        EXECUTE 'TRUNCATE TABLE ' || quote_ident(r.tablename) || ' RESTART IDENTITY CASCADE';
    END LOOP;
END $$;

INSERT INTO departments(name, note)
VALUES ('Mammo center', 'Only during weekdays'),
       ('CT surgery');

INSERT INTO consult_types(consult_type, duration_minutes)
VALUES ('Mammography', 10),
       ('Breast ultrasound', 15),
       ('Breast MRI', 30),
       ('Abdominal CT surgery', 20),
       ('High Resolution CT', 30),
       ('Native x-ray', 5);

INSERT INTO day_templates(name, department_id)
VALUES ('Monday', 1),
       ('Thursday', 1),
       ('Monday', 2),
       ('Tuesday', 2);

INSERT INTO day_template_consults(consult_type_id, day_template_id, time, note)
VALUES (1, 1, '08:00', null),
       (1, 1, '08:10', null),
       (1, 1, '08:20', null),
       (1, 1, '08:30', null),
       (1, 1, '08:40', null),
       (1, 1, '08:50', null),
       (1, 1, '09:00', null),
       (1, 1, '09:10', null),
       (1, 1, '09:20', null),
       (2, 1, '09:30', null),
       (2, 1, '09:45', null),
       (2, 1, '10:00', null),
       (2, 1, '10:15', null),
       (2, 1, '10:30', null),
       (2, 1, '10:45', null),
       (2, 1, '11:00', null),
       (2, 1, '10:45', null),
       (3, 1, '11:00', null),
       (3, 1, '11:30', null),
       (null, 1, '12:00', 'Lunch break'),
       (3, 1, '13:00', null),
       (3, 1, '13:30', null),

       (2, 2, '10:00', null),
       (2, 2, '10:15', null),
       (2, 2, '10:30', null),
       (2, 2, '10:45', null),
       (2, 2, '11:00', null),
       (2, 2, '10:45', null),
       (3, 2, '11:00', null),
       (3, 2, '11:30', null),
       (null, 2, '12:00', 'Lunch break'),
       (1, 2, '13:00', null),
       (1, 2, '13:10', null),
       (1, 2, '13:20', null),
       (1, 2, '13:30', null),
       (1, 2, '13:40', null),
       (1, 2, '13:50', null),

       (4, 3, '08:00', null),
       (4, 3, '08:20', null),
       (4, 3, '08:40', null),
       (5, 3, '09:00', null),
       (5, 3, '09:30', null),
       (5, 3, '10:00', null),
       (5, 3, '10:30', null),
       (5, 3, '11:00', null),
       (5, 3, '11:30', null),
       (null, 3, '12:00', 'Lunch break'),
       (6, 3, '13:00', null),
       (6, 3, '13:05', null),
       (6, 3, '13:10', null),
       (6, 3, '13:15', null),
       (6, 3, '13:20', null),
       (6, 3, '13:25', null),

       (4, 4, '08:00', null),
       (4, 4, '08:20', null),
       (4, 4, '08:40', null),
       (5, 4, '09:00', null),
       (5, 4, '09:30', null),
       (5, 4, '10:00', null),
       (5, 4, '10:30', null),
       (5, 4, '11:00', null),
       (5, 4, '11:30', null),
       (null, 4, '12:00', 'Lunch break'),
       (6, 4, '13:00', null),
       (6, 4, '13:05', null),
       (6, 4, '13:10', null),
       (6, 4, '13:15', null),
       (6, 4, '13:20', null),
       (6, 4, '13:25', null);
