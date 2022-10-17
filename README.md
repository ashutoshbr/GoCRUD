<h1 align="center">GoCRUD 🐀</h1>
<h4 align="center">A minimal project to perform CRUD operations on mongoDB.</h4>

## Available routes 🚀
```bash
GET       localhost:8000/
GET       localhost:8000/read
POST      localhost:8000/create
PUT       localhost:8000/update
DELETE    localhost:8000/delete
```
## Dependencies 🔧
- MongoDB drivers

## Execution ⚙️
1. Install the required dependencies.
```bash
go mod dowload
```
2. Run locally.
```bash
# Runs at port 8000 by default
go build .\cmd\ && .\cmd.exe
```
3. Use `curl` or `postman` or any other client to perform CRUD operations at the available routes.