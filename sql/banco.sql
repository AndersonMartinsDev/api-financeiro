CREATE DATABASE IF NOT EXISTS finance;

USE finance;

CREATE TABLE usuario(
    id bigInt auto_increment primary key,
    avatar BLOB, 
    nome varchar(200) not null,
    username varchar(200) not null,
    senha varchar(100) not null,
    email varchar(200)
)ENGINE = INNODB;

CREATE TABLE associacao_carteira_usuario(
    usuario_id bigInt not NULL,
    FOREIGN KEY (usuario_id)
    REFERENCES usuario(id)
	ON DELETE CASCADE,
    carteira_id VARCHAR(100) not NULL PRIMARY key
)ENGINE = INNODB;

CREATE TABLE envelopes(
    id bigInt auto_increment primary key,
    titulo varchar(12) not null unique,
    valor double not null,
    observacao varchar(200),
    carteira VARCHAR(100),
     FOREIGN KEY(carteira)
     REFERENCES associacao_carteira_usuario(carteira_id)
) ENGINE=INNODB;

CREATE TABLE despesas(
    id bigInt auto_increment primary key,
    titulo varchar(30) not null,
    valor double not null,
    quitada bool not null, 
    tipo varchar(20) not null,
    dia_vencimento bigInt,
    data_cadastro timestamp default current_timestamp(),
    observacao varchar(255),
    envelope_id bigInt,
     FOREIGN KEY(envelope_id)
     REFERENCES envelopes(id),
    carteira VARCHAR(100),
     FOREIGN KEY(carteira)
     REFERENCES associacao_carteira_usuario(carteira_id)
)ENGINE = INNODB;

CREATE TABLE pagamentos(
    id bigInt auto_increment primary key,
    valor double not null,
    data_pagamento timestamp,
    data_vencimento timestamp,
    forma_pagamento varchar(20),
    FOREIGN KEY(usuario_id)
    REFERENCES usuario(id),
    despesa_id bigInt not null,
    FOREIGN KEY (despesa_id)
    REFERENCES despesas(id)
)ENGINE = INNODB;

-- finance.v_despesa source
CREATE OR REPLACE
algorithm = undefined view `v_despesa`
AS
  SELECT DISTINCT `des`.`id`
                  AS `id`,
                  `des`.`titulo`
                     AS `titulo`,
                  `des`.`tipo`
                     AS `tipo`,
                  `des`.`data_cadastro`
                     AS `data_cadastro`,
                  `pgto`.`data_vencimento`
                     AS `data_vencimento`,
                  IF(`des`.`tipo` <> 'PARCELADA', Round(`des`.valor
                  , 2), `pgto`.valor)  AS valor,
                  IF(( `des`.`tipo` = 'UNICA' ),
                  'Essa conta não possui frequência',
                  Concat('Vence dia ', IF(( `des`.`tipo` <> 'PARCELADA' ),
                                       `des`.`dia_vencimento`,
                                       Dayofmonth(`pgto`.`data_vencimento`))))
                     AS `condicao`,
                  IF(( `des`.`tipo` <> 'PARCELADA' ), 'à vista', Concat(
                  (SELECT Count(`pg`.`id`)
                   FROM   `pagamentos` `pg`
                   WHERE  ( ( `pg`.`despesa_id` = `des`.`id` )
                            AND ( `pg`.`data_pagamento` IS NOT NULL ) )), '/',
                  (SELECT Count(`pg`.`id`)
                   FROM
                  `pagamentos` `pg`
                  WHERE  (
                  `pg`.`despesa_id` = `des`.`id` ))))
                     AS `pagamento`,
                   IF(( Date_format(`pgto`.`data_pagamento`, '%m/%y') =
                    Date_format(Now(), '%m/%y')
                       ), true, false) AS
                  `quitada`,
                  `des`.`carteira`
                     AS `carteira`
  FROM   (`despesas` `des`
          LEFT JOIN `pagamentos` `pgto`
                 ON (( `pgto`.`despesa_id` = `des`.`id` )))
  GROUP  BY `des`.`id`,
            `pgto`.`id`; 