CREATE TABLE moviesgolang.movie
(
    id           MEDIUMINT NOT NULL AUTO_INCREMENT,
    title        varchar(100) NOT NULL,
    `cast`       varchar(250) NULL,
    release_date DATE         NULL,
    genre        varchar(100) NULL,
    director     varchar(100) NULL,
    CONSTRAINT movie_PK PRIMARY KEY (id)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_0900_ai_ci;


CREATE TABLE moviesgolang.`user`
(
    id       MEDIUMINT NOT NULL AUTO_INCREMENT,
    username varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    CONSTRAINT user_PK PRIMARY KEY (id),
    CONSTRAINT user_UN UNIQUE KEY (username)
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4
    COLLATE = utf8mb4_0900_ai_ci;


CREATE TABLE moviesgolang.wish_list (
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    user_id MEDIUMINT NOT NULL,
    movie_id MEDIUMINT NOT NULL,
    comment VARCHAR(300) NULL,
    CONSTRAINT wish_list_PK PRIMARY KEY (id),
    CONSTRAINT wish_list_FK FOREIGN KEY (user_id)
        REFERENCES moviesgolang.`user` (id)
        ON DELETE CASCADE,
    CONSTRAINT wish_list_FK_1 FOREIGN KEY (movie_id)
        REFERENCES moviesgolang.movie (id)
        ON DELETE CASCADE
)  ENGINE=INNODB DEFAULT CHARSET=UTF8MB4 COLLATE = UTF8MB4_0900_AI_CI;