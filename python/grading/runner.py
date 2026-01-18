"""
CLI entrypoint for grading.
Called by Go backend.
"""

import argparse
import json
import importlib
from .engine_registry import ENGINE_REGISTRY


def load_engine(engine_name: str):
    """
    Dynamically load grading engine class.
    """
    if engine_name not in ENGINE_REGISTRY:
        raise ValueError(f"Unknown grading engine: {engine_name}")

    module_path, class_name = ENGINE_REGISTRY[engine_name].rsplit(".", 1)
    module = importlib.import_module(module_path)
    return getattr(module, class_name)()


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--engine", required=True)
    parser.add_argument("--rules", required=True)
    parser.add_argument("--submission", required=True)

    args = parser.parse_args()

    # Load grading rules
    with open(args.rules, "r") as f:
        rules = json.load(f)

    # Load engine
    engine = load_engine(args.engine)

    # Grading pipeline
    parsed_data = engine.parse(args.submission)
    validation = engine.validate(parsed_data, rules)
    result = engine.grade(validation)

    # Output JSON (Go backend will capture this)
    print(json.dumps(result, indent=2))


if __name__ == "__main__":
    main()
