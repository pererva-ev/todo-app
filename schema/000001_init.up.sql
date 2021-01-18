
CREATE TABLE task (
    id int(20) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(50) DEFAULT NULL,
    description varchar(50) DEFAULT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE project (
    id int(20) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(50) DEFAULT NULL,
    description varchar(50) DEFAULT NULL,
    PRIMARY KEY (id)
);


CREATE TABLE comment (
    id int(20) unsigned NOT NULL AUTO_INCREMENT,
    text varchar(50) DEFAULT NULL,
    PRIMARY KEY (id)
);


CREATE TABLE column (
    id int(20) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(50) DEFAULT NULL,
    status varchar(50) DEFAULT NULL,
    PRIMARY KEY (id)
);