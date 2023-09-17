CREATE TABLE public."group" (
    Id SERIAL NOT NULL,
    MaxCurrent FLOAT NOT NULL,
    CONSTRAINT pk_Group PRIMARY KEY (Id)
);

CREATE TABLE public.ChargePointStatus (
    ChargePointId SERIAL NOT NULL,
    "Priority" int8 NOT NULL,
    GroupId int8 NOT NULL,
    CONSTRAINT fk_ChargePointStatus_Group FOREIGN KEY (GroupId) REFERENCES public."group"(Id),
    CONSTRAINT pk_ChargePointStatus PRIMARY KEY (ChargePointId)
);

CREATE TABLE public.ChargePointConnector (
    ChargePointConnectorId SERIAL NOT NULL,
    "Status" varchar(32) NOT NULL,
    ChargePointId int8 NOT NULL,
    CONSTRAINT fk_ChargePointConnector_ChargePointStatus FOREIGN KEY (ChargePointId) REFERENCES public.ChargePointStatus(ChargePointId),
    CONSTRAINT pk_ChargePointConnector PRIMARY KEY (ChargePointConnectorId)
);
