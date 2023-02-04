# Curso Alura : Go Crie uma Aplicação Web. 


Instalar o PostGree e criar um banco de dados chamado alura_loja. 

usar os scripts para criar a tabela de produtos e popular:

```
CREATE TABLE produtos (
             id serial primary key,
             nome varchar,
             descricao varchar,
             preco decimal,
             quantidade integer)


INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES
('Camiseta', 'Preta', 19, 10),
('Fone', 'Muito bom', 99, 5)

```
