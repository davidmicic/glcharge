INSERT INTO public."group" (MaxCurrent) VALUES
    (200.0),
    (300.0);

INSERT INTO public.ChargePointStatus (ChargePointId, "Priority", GroupId) VALUES
    (1, 2, 1),
    (2, 1, 1),
    (3, 3, 2),
    (4, 2, 2),
    (5, 1, 2);

INSERT INTO public.ChargePointConnector (ChargePointConnectorId, "Status", ChargePointId) VALUES
    (1, 'Available', 1),
    (2, 'Available', 1),
    (3, 'Charging', 1),
    (4, 'Available', 1),
    (5, 'Charging', 2),
    (6, 'Available', 2),
    (7, 'Charging', 2),
    (8, 'Available', 3),
    (9, 'Available', 3),
    (10, 'Available', 4),
    (11, 'Charging', 4),
    (12, 'Available', 4),
    (13, 'Charging', 5),
    (14, 'Available', 5),
    (15, 'Charging', 5);

