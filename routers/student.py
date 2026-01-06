from fastapi import APIRouter, UploadFile, Depends
from database import SessionLocal
from models import Submission, Assignment
from grading.proteus.grader import ProteusGrader
from dependencies import get_current_user
import json, os

router = APIRouter(prefix="/student")

@router.post("/submit")
def submit(assignment_id: int, file: UploadFile, user=Depends(get_current_user)):
    os.makedirs("uploads", exist_ok=True)
    path = f"uploads/{file.filename}"

    with open(path, "wb") as f:
        f.write(file.file.read())

    db = SessionLocal()
    assignment = db.query(Assignment).filter_by(id=assignment_id).first()
    constraints = json.loads(assignment.constraints)

    grader = ProteusGrader(constraints)
    parsed = grader.parse(path)
    result = grader.grade(parsed)

    sub = Submission(
        assignment_id=assignment_id,
        student_id=1,
        file_path=path,
        score=result.score,
        violations=str(result.violations),
        flags=str(result.flags)
    )
    db.add(sub)
    db.commit()

    return {
        "score": result.score,
        "violations": result.violations,
        "flags": result.flags
    }
