create table Log
(
    id            INTEGER GENERATED BY DEFAULT ON NULL AS IDENTITY,
    user_id       INTEGER,
    phone         VARCHAR2(20),
    action_id     INTEGER,
    action_title  VARCHAR2(255),
    action_type   VARCHAR2(50),
    message       VARCHAR2(1000),
    sender        VARCHAR2(100),
    status        VARCHAR2(50),
    language      VARCHAR2(10),
    full_response VARCHAR2(1000),
    created       TIMESTAMP,
    updated       TIMESTAMP,
    message_id    VARCHAR2(100),
    CONSTRAINT logs_pk PRIMARY KEY (id)
);

ALTER TABLE Log
    ADD statusDelive NUMBER(2,0);
ALTER TABLE Log
    ADD cost NUMBER(10,4);
ALTER TABLE Log
    ADD CONSTRAINT unique_message_phone_sender UNIQUE (message_id, phone, sender);