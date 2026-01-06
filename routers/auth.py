from fastapi import APIRouter
from database import SessionLocal
from models import User
from schemas import UserCreate
from auth import hash_password, create_token

router = APIRouter(prefix="/auth")

@router.post("/signup")
def signup(user: UserCreate):
    db = SessionLocal()
    db.add(User(
        email=user.email,
        password=hash_password(user.password),
        role=user.role
    ))
    db.commit()
    return {"message": "User created"}

@router.post("/login")
def login(user: UserCreate):
    token = create_token({"email": user.email, "role": user.role})
    return {"access_token": token}

