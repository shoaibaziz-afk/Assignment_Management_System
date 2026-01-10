"""
Defines electrical rules for grading
"""

def validate_circuit(design, constraints):
    violations = []

    required_components = constraints.get("required_components", [])
    present_components = [c.type for c in design.components]

    # Check required components
    for comp in required_components:
        if comp not in present_components:
            violations.append(f"Missing component: {comp}")

    # Check minimum number of nets
    if len(design.nets) < constraints.get("min_nets", 1):
        violations.append("Insufficient net connections")

    return violations

