"""
Proteus Analog Circuit Grading Engine

MVP capabilities:
- Parse components from Proteus .dsn files
- Recognize real Proteus device declarations
- Validate presence of required components
- Deterministic scoring
"""

import zipfile
from ..base_engine import GradingEngine


# Logical component categories → Proteus keywords
COMPONENT_KEYWORDS = {
    "VDC": ["VDC", "DC", "BATTERY", "POWER"],
    "RES": ["RES", "RESISTOR", "R"],
    "LED": ["LED"]
}


class ProteusAnalogEngine(GradingEngine):
    def parse(self, submission_path: str) -> dict:
        """
        Parse Proteus project ZIP and extract component declarations.
        """

        if not zipfile.is_zipfile(submission_path):
            raise ValueError("Submission must be a ZIP file")

        components = []

        with zipfile.ZipFile(submission_path, "r") as zip_ref:
            for file_name in zip_ref.namelist():
                if file_name.lower().endswith(".dsn"):
                    content = zip_ref.read(file_name).decode(errors="ignore")

                    for line in content.splitlines():
                        line_upper = line.upper()

                        # Detect real Proteus component lines
                        if (
                            "DEVICE" in line_upper
                            or "$COMPONENT" in line_upper
                            or "$DEVICE" in line_upper
                        ):
                            components.append(line_upper.strip())

        return {
            "components": components,
            "component_count": len(components),
        }

    def validate(self, parsed_data: dict, rules: dict) -> dict:
        """
        Validate parsed circuit against grading rules.
        """

        results = {
            "rules": {},
            "violations": []
        }

        required = rules.get("required_components", {})

        for component, required_count in required.items():
            found = 0

            for comp_line in parsed_data["components"]:
                for keyword in COMPONENT_KEYWORDS.get(component, []):
                    if keyword in comp_line:
                        found += 1
                        break

            if found >= required_count:
                results["rules"][component] = True
            else:
                results["rules"][component] = False
                results["violations"].append(
                    f"Missing {component}: required {required_count}, found {found}"
                )

        return results

    def grade(self, validation_report: dict) -> dict:
        """
        Deterministic grading logic.
        """

        total_rules = len(validation_report["rules"])
        passed_rules = sum(1 for v in validation_report["rules"].values() if v)

        score = int((passed_rules / max(total_rules, 1)) * 100)

        return {
            "score": score,
            "passed_rules": passed_rules,
            "total_rules": total_rules,
            "violations": validation_report["violations"],
            "status": "PASS" if score >= 50 else "FAIL"
        }
