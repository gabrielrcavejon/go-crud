CREATE DATABASE IF NOT EXISTS gocrud;

USE gocrud;

CREATE TABLE usuario (
  idusuario INT AUTO_INCREMENT PRIMARY KEY,
  nome VARCHAR(150) NOT NULL,
  email VARCHAR(200) NOT NULL UNIQUE,
  senha VARCHAR(250) NOT NULL,
  criadoEm timestamp DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE produto (
  idproduto INT AUTO_INCREMENT PRIMARY KEY,
  nome VARCHAR(200) NOT NULL,
  descricao TEXT,
  idusuario INT NOT NULL,
  criadoEm timestamp DEFAULT current_timestamp(),
  CONSTRAINT fk_produto_usuario
    FOREIGN KEY (idusuario)
    REFERENCES usuario(idusuario)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
