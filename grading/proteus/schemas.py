from dataclasses import dataclass
from typing import List

@dataclass
class Component:
    ref: str
    type: str

@dataclass
class CircuitDesign:
    components: List[Component]
    nets: List[str]
