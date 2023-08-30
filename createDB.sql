create table Log
(
    id            INTEGER,
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
)