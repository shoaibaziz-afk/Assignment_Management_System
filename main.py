from fastapi import FastAPI
from database import Base, engine
from routers import auth, professor, student

Base.metadata.create_all(bind=engine)

app = FastAPI(title="Anti-AI Assignment MVP")

app.include_router(auth.router)
app.include_router(professor.router)
app.include_router(student.router)
