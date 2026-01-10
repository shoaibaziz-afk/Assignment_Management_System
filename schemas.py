"""
Schemas define:
- What data is allowed to come FROM the user
- Input validation for API requests
"""

from pydantic import BaseModel

# Used for signup and login
class UserCreate(BaseModel):
    email: str
    password: str
    role: str   # "student" or "professor"

# Used when professor creates an assignment
class AssignmentCreate(BaseModel):
    title: str
    description: str
    constraints: dict   # Grading rules

