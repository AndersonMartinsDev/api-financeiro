CREATE DATABASE IF NOT EXISTS finance;

USE finance;

CREATE TABLE envelope(
    id bigInt auto_increment primary key,
    titulo varchar(12) not null unique,
    valor double not null,
    observacao varchar(200)
) ENGINE=INNODB;


CREATE TABLE recorrencia(
    id bigInt auto_increment primary key,
    meses int not null,
    dia_vencimento int not null
)ENGINE = INNODB;

CREATE TABLE despesas(
    id bigInt auto_increment primary key,
    titulo varchar(22) not null,
    valor double not null,
    quitada bool not null, 
    fixa bool not null,
    dia_vencimento bigInt,
    data_cadastro timestamp default current_timestamp(),
    envelope_id bigInt not null,
     FOREIGN KEY(envelope_id)
     REFERENCES envelope(id)
    ON DELETE CASCADE,
    recorrencia_id bigInt,
    FOREIGN KEY(recorrencia_id)
    REFERENCES recorrencia(id)
    ON DELETE CASCADE
)ENGINE = INNODB;
