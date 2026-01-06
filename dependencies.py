from fastapi import Header, HTTPException
from auth import decode_token

def get_current_user(authorization: str = Header(...)):
    try:
        token = authorization.split(" ")[1]
        return decode_token(token)
    except:
        raise HTTPException(status_code=401, detail="Invalid token")

def professor_only(user=Depends(get_current_user)):
    if user["role"] != "professor":
        raise HTTPException(status_code=403)
    return user

def student_only(user=Depends(get_current_user)):
    if user["role"] != "student":
        raise HTTPException(status_code=403)
    return user
