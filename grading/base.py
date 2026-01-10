"""
This file is responsible for:
- Creating the database connection
- Providing a session object for DB operations
"""

from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker, declarative_base

# SQLite database for MVP (simple and file-based)
DATABASE_URL = "sqlite:///./mvp.db"

# Create database engine
engine = create_engine(
    DATABASE_URL,
    connect_args={"check_same_thread": False}  # Needed for SQLite
)

# SessionLocal is used whenever we talk to the database
SessionLocal = sessionmaker(bind=engine)

# Base class for all database models
Base = declarative_base()
