CREATE TABLE Log
(
    id            SERIAL PRIMARY KEY,
    user_id       INTEGER,
    phone         VARCHAR(20),
    action_id     INTEGER,
    action_title  VARCHAR(255),
    action_type   VARCHAR(50),
    message       VARCHAR(1000),
    sender        VARCHAR(100),
    status        VARCHAR(50),
    language      VARCHAR(10),
    full_response VARCHAR(1000),
    created       TIMESTAMP,
    updated       TIMESTAMP,
    message_id    VARCHAR(100),
    statusDelive  INTEGER,
    cost          NUMERIC(10,4),
    CONSTRAINT unique_message_phone_sender UNIQUE (message_id, phone, sender)
);
