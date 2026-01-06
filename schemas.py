from pydantic import BaseModel

class UserCreate(BaseModel):
    email: str
    password: str
    role: str

class AssignmentCreate(BaseModel):
    title: str
    description: str
    constraints: dict
