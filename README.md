## Teste Sipub Tech

Configurações:

A solução foi desenvolvida através do Sistema Operacional Ubuntu 20.04.5 LTS; <br>
A IDE utilizada foi o VS Code, versão 1.74.2, sendo adicionadas extensões para o Docker e o Go; <br>
Foi utilizado o Docker versão 20.10.22 e o Go versão 1.19.4.<br>

Para levantar o container do Postgres, utilizou-se o comando: <br>
``` docker run -d --name api-todo  -p 5432:5432 -e POSTGRES_PASSWORD=1234 postgres:13.5 ```<br>

Para conectar com o container:<br>
```docker exec -it api-todo psql -U postgres```<br>

Para criar o database:<br>
```create database api_todo;```<br>

Para criar um usuário:<br>
```create user user_todo;```<br>

Para definir uma senha para o usuário criado:<br>
```alter user user_todo with encrypted password ‘1122’;```<br>

Para dar permissão para o usuário utilizar o database:<br>
```grant all privileges on database api_todo to user_todo;```<br>

Para conectar com o database:<br>
```\c api_todo;```<br>
 
Para a criação da tabela de tarefas (todos):<br>
```create table todos (id serial primary key, title varchar, description text, done bool default FALSE);```<br>

Para dar permissão para o usuário gravar dados no banco:<br>
```grant all privileges on table todos to user_todo;```<br>

Para realizar a importação das bibliotecas utilizadas no projeto:<br>
```go mod tidy```<br>

Para rodar o servidor:<br>
```go run main.go```<br>

O arquivo config.toml possui as configurações da API e do banco de dados Postgres<br>
<br>
 ### Requisitos

A solução foi implementada através de um CRUD, estruturado em 4 packages:<br>
<br>
DB: realiza a conexão com o banco de dados;<br>
CONFIGS: recebe as configurações do servidor;<br>
HANDLERS: recebe as chamadas das API e trata as chamadas; e<br>
MODELS: realiza transações com o banco de dados.<br>
<br>
 
O arquivo API-todo.postman_collection.json contém as requisições necessárias para atender aos parâmetros propostos.
<br>
<br>
- Para criar tarefas: método POST + JSON
```
{
    "title": "tarefa 1",
    "description": "descrição 1",
    "done": true
}
```
- Para editar tarefas: método PUT + ID + JSON (o usuário pode escolher qual item deseja editar, inclusive a mudança de status entre true ou false)
- Para deletar tarefas: método DELETE + ID
- Para listar todas as tarefas: método GET_ALL
- Para aplicar filtro de status: método GET_STATUS + STATUS (true ou false)
