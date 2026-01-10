"""
Final Proteus grading logic
"""

from grading.base import BaseGrader, GradingResult
from grading.proteus.parser import parse_dsn
from grading.proteus.rules import validate_circuit

class ProteusGrader(BaseGrader):

    def __init__(self, constraints):
        self.constraints = constraints

    def parse(self, file_path):
        return parse_dsn(file_path)

    def grade(self, parsed_data):
        violations = validate_circuit(parsed_data, self.constraints)

        # Basic scoring logic
        score = max(0, 100 - 20 * len(violations))

        # Flags for suspicious/poor designs
        flags = []
        if score < 40:
            flags.append("Severe design issues")

        return GradingResult(score, violations, flags)

