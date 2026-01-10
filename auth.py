"""
This file handles:
- Password hashing
- Token creation
- Token decoding
"""

from passlib.context import CryptContext
from jose import jwt
from datetime import datetime, timedelta

SECRET_KEY = "dev-secret"   # Change in production
ALGORITHM = "HS256"

# Password hashing configuration
pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")

def hash_password(password: str):
    """Hash a plaintext password"""
    return pwd_context.hash(password)

def verify_password(password, hashed):
    """Verify password against hash"""
    return pwd_context.verify(password, hashed)

def create_token(data: dict):
    """
    Create JWT token with expiration
    """
    payload = data.copy()
    payload["exp"] = datetime.utcnow() + timedelta(hours=12)
    return jwt.encode(payload, SECRET_KEY, algorithm=ALGORITHM)

def decode_token(token: str):
    """Decode JWT token"""
    return jwt.decode(token, SECRET_KEY, algorithms=[ALGORITHM])

