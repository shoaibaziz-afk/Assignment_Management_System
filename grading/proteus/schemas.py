"""
Defines Python representation of a circuit
"""

from dataclasses import dataclass
from typing import List

@dataclass
class Component:
    ref: str     # Reference ID (U1, R1, etc.)
    type: str    # AND, OR, XOR, etc.

@dataclass
class CircuitDesign:
    components: List[Component]
    nets: List[str]

