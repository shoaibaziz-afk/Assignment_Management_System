from fastapi import APIRouter, Depends
from database import SessionLocal
from models import Assignment
from schemas import AssignmentCreate
from dependencies import get_current_user

router = APIRouter(prefix="/professor")

@router.post("/assignments")
def create_assignment(data: AssignmentCreate, user=Depends(get_current_user)):
    db = SessionLocal()
    a = Assignment(
        title=data.title,
        description=data.description,
        constraints=str(data.constraints),
        professor_id=1
    )
    db.add(a)
    db.commit()
    return {"message": "Assignment created"}
