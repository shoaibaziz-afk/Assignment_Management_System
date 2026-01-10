"""
Professor endpoints
"""

from fastapi import APIRouter, Depends
from database import SessionLocal
from models import Assignment
from schemas import AssignmentCreate
from dependencies import professor_only
import json

router = APIRouter(prefix="/professor")

@router.post("/assignments")
def create_assignment(data: AssignmentCreate, user=Depends(professor_only)):
    db = SessionLocal()
    assignment = Assignment(
        title=data.title,
        description=data.description,
        constraints=json.dumps(data.constraints),
        professor_id=1
    )
    db.add(assignment)
    db.commit()
    return {"message": "Assignment created"}

