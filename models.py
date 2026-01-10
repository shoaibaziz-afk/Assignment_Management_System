"""
This file defines:
- What tables exist in the database
- What columns each table has
"""

from sqlalchemy import Column, Integer, String, ForeignKey, DateTime, Text
from datetime import datetime
from database import Base

# ---------------- USER TABLE ----------------
class User(Base):
    __tablename__ = "users"

    id = Column(Integer, primary_key=True)
    email = Column(String, unique=True, index=True)
    password = Column(String)        # Stored as hashed password
    role = Column(String)            # "student" or "professor"


# ---------------- ASSIGNMENT TABLE ----------------
class Assignment(Base):
    __tablename__ = "assignments"

    id = Column(Integer, primary_key=True)
    title = Column(String)
    description = Column(Text)

    # Constraints are stored as JSON converted to string
    constraints = Column(Text)

    professor_id = Column(Integer, ForeignKey("users.id"))
    created_at = Column(DateTime, default=datetime.utcnow)


# ---------------- SUBMISSION TABLE ----------------
class Submission(Base):
    __tablename__ = "submissions"

    id = Column(Integer, primary_key=True)
    assignment_id = Column(Integer)
    student_id = Column(Integer)

    file_path = Column(String)       # Where the uploaded .DSN file is stored
    score = Column(Integer)
    violations = Column(Text)        # Rule violations
    flags = Column(Text)             # Suspicious activity flags

    submitted_at = Column(DateTime, default=datetime.utcnow)

