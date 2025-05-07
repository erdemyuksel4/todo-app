# 📝 Todo App API

This application is a **RESTful API** that allows users to create and manage their to-do lists.  
It is built using the **Gin Framework** and uses **JWT (JSON Web Token)** for authentication.

I also have same project that you can examine: https://github.com/erdemyuksel4/GoTodoApp 

---

## 🚀 Getting Started

Follow the steps below to run the project on your local environment.

### Requirements

- Go 1.18 or higher
- Gin Web Framework
- JWT library
- Go module management (`go mod`)

### Installation

```bash
# Clone the repository
git clone https://github.com/erdemyuksel4/todoapp.git
cd todoapp

# Install dependencies
go mod tidy

# (Optional) Set environment variables
# You can create a .env file and set a value like PORT=9090

# Start the server
go run main.go

By default, the server will run at localhost:8080.
```
## API USAGE

#### API Endpoints

```http
GET /api/todos
GET /api/todobyid/:id
GET /api/todolists
PATCH /api/complete/:id
PATCH /api/changemessage
POST /api/addtodo
DELETE /api/deletetodo
DELETE /api/deletelist
POST /api/addlist
```

#### All Todos

```http
  GET /api/todos
```
You can see all messages that you have. If you are admin you can see all todos
#### Get a TodoStep By Step Id
```http
  GET /api/todobyid/:id
```
| Parametre | Tip     | 
| :-------- | :------- | 
| `id`      | `string` | 

#### Get a TodoStep By Step Id
```http
  GET /api/todolists
```
If you are not admin you can see only yours

#### Complete a TodoStep On Your List
```http
  PATCH /api/complete/:id
```
| Parametre | Tip     | 
| :-------- | :------- | 
| `id`      | `string` | 

#### You can Change The Todo Step Message
```http
  PATCH /api/changemessage
```
| Parametre | Tip     | 
| :-------- | :------- | 
| `todoId`      | `int` | 
| `message`      | `string` | 

#### Add New Todo Step
```http
  POST /api/addtodo
```
| Parametre | Tip     | 
| :-------- | :------- | 
| `listId`      | `int` | 
| `message`      | `string` | 

#### Delete Todo Step
```http
  DELETE /api/deletetodo
```
| Parametre | Tip     | 
| :-------- | :------- | 
| `todoId`      | `int` | 

#### Delete Todo List
```http
  DELETE /api/deletelist
```
| Parametre | Tip     | 
| :-------- | :------- | 
| `id`      | `int` | 

#### Add New List
```http
  POST /api/addlist
```
| Parametre | Tip     | 
| :-------- | :------- | 
| `title`      | `string` |

Add new list and use this for new steps 

  ###

There is a out.exe file that you can run  

```bash
  ./out
```

  
