from fastapi import FastAPI
from fastapi.responses import JSONResponse

app = FastAPI()

@app.get("/ping")
def ping():
    return JSONResponse(content={"status": "ok", "service": "2"})

@app.get("/hello")
def hello():
    return JSONResponse(content={"message": "Hello from Service 2"})

