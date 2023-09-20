INSERT INTO public."group" (MaxCurrent) VALUES
    (200.0),
    (300.0);

INSERT INTO public.ChargePointStatus ("Priority", GroupId) VALUES
    (1, 1),
    (0, 1),
    (2, 2),
    (1, 2),
    (0, 2);

INSERT INTO public.ChargePointConnector ("Status", ChargePointId) VALUES
    ('Available', 1),
    ('Available', 1),
    ('Charging', 1),
    ('Available', 1),
    ('Charging', 2),
    ('Available', 2),
    ('Charging', 2),
    ('Available', 3),
    ('Available', 3),
    ('Available', 4),
    ('Charging', 4),
    ('Available', 4),
    ('Charging', 5),
    ('Available', 5),
    ('Charging', 5);

