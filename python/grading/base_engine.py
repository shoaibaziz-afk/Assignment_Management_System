"""
Abstract base class for all grading engines.

Every grading engine MUST:
- Parse the submission
- Validate it against rules
- Produce a deterministic grading result
"""

from abc import ABC, abstractmethod


class GradingEngine(ABC):
    @abstractmethod
    def parse(self, submission_path: str) -> dict:
        """
        Convert raw submission files into a structured internal representation.
        Example output:
        {
            "components": [...],
            "nets": [...],
            "parameters": {...}
        }
        """
        pass

    @abstractmethod
    def validate(self, parsed_data: dict, rules: dict) -> dict:
        """
        Validate parsed data against grading rules.
        Returns rule-by-rule validation status.
        """
        pass

    @abstractmethod
    def grade(self, validation_report: dict) -> dict:
        """
        Compute final score, breakdown, and feedback.
        """
        pass
