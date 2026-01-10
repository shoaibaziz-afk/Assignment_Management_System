"""
Application entry point
"""

from fastapi import FastAPI
from database import Base, engine
from routers import auth, professor, student

# Create database tables
Base.metadata.create_all(bind=engine)

# Initialize FastAPI app
app = FastAPI(title="Anti-AI Engineering Assignment MVP")

# Register routers
app.include_router(auth.router)
app.include_router(professor.router)
app.include_router(student.router)

