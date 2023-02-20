CREATE DATABASE IF NOT EXISTS finance;

USE finance;

CREATE TABLE envelopes(
    id bigInt auto_increment primary key,
    titulo varchar(12) not null unique,
    valor double not null,
    observacao varchar(200)
) ENGINE=INNODB;


-- resolver problema do CASCADE
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
     REFERENCES envelopes(id)
)ENGINE = INNODB;

CREATE TABLE pagamentos(
    id bigInt auto_increment primary key,
    valor double not null,
    data_pagamento timestamp ,
    data_vencimento timestamp,
    forma_pagamento varchar(20),
    -- usuario_id bigInt not null, 
    -- FOREIGN KEY(usuario_id)
    -- REFERENCES usuario(id)
    -- ON NO ACTION
    despesa_id bigInt not NULL,
    FOREIGN KEY (despesa_id)
    REFERENCES despesas(id)
)ENGINE = INNODB;

CREATE VIEW v_despesa AS 
SELECT 
DISTINCT(des.id),
des.titulo AS titulo,
des.tipo AS tipo,
FORMAT(des.valor,2) AS valor,
IF(des.tipo='UNICA','Essa conta não possui frequência',CONCAT('Vence dia ',IF(des.tipo <> 'PARCELADA', des.dia_vencimento,DAY(pgto.data_vencimento)))) AS condicao,
IF(des.tipo <> 'PARCELADA','à vista',CONCAT((SELECT COUNT(pg.id) FROM pagamentos pg WHERE pg.despesa_id = des.id AND pg.data_pagamento IS NOT NULL),'/' ,(SELECT COUNT(pg.id) FROM pagamentos pg WHERE pg.despesa_id = des.id))) AS pagamento,
IF(des.tipo <> 'PARCELADA', des.quitada, IF(DATE_FORMAT(pgto.data_pagamento,'%m/%y')=DATE_FORMAT(NOW(),'%m/%y'),TRUE,FALSE)) AS quitada 
FROM despesas des 
LEFT JOIN pagamentos pgto ON pgto.despesa_id = des.id
GROUP BY 
des.id,
pgto.id
