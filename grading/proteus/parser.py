"""
Parses Proteus .DSN file into Python objects
"""

import re
from grading.proteus.schemas import Component, CircuitDesign

def parse_dsn(file_path: str) -> CircuitDesign:
    # Read the .DSN file as text
    with open(file_path, "r", errors="ignore") as f:
        content = f.read()

    components = []
    nets = set()

    # Extract components using regex
    comp_matches = re.findall(r'Ref=(\w+).*?Type=(\w+)', content)
    for ref, ctype in comp_matches:
        components.append(Component(ref=ref, type=ctype))

    # Extract net names
    net_matches = re.findall(r'Net=(\w+)', content)
    nets.update(net_matches)

    return CircuitDesign(
        components=components,
        nets=list(nets)
    )

