class GradingResult:
    def __init__(self, score, violations, flags):
        self.score = score
        self.violations = violations
        self.flags = flags

class BaseGrader:
    def parse(self, file_path: str):
        raise NotImplementedError

    def grade(self, parsed_data):
        raise NotImplementedError
