def validate_circuit(design, constraints):
    violations = []

    required = constraints.get("required_components", [])
    present = [c.type for c in design.components]

    for comp in required:
        if comp not in present:
            violations.append(f"Missing component: {comp}")

    if len(design.nets) < constraints.get("min_nets", 1):
        violations.append("Insufficient nets")

    return violations
