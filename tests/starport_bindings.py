"""
Common functions for tests
"""

import sys
import subprocess


def starport_bindings(args):
    """
    Get the address of an account in the blockchain
    """

    command = "./starport_bindings.sh"
    for arg in args:
        command += " "
        command += arg

    process = subprocess.Popen(
        command.split(), stdout=subprocess.PIPE
    )
    output, error = process.communicate()

    if error is not None:
        print("Error:", error)

    return output.strip().decode("utf-8")

def main():
    """
    Run starport bindings standalone
    """

    assert len(sys.argv) >= 1, "Expected at least one arg"
    args = list(sys.argv[1:])
    starport_bindings(args)

if "__name__" == "__main__":
    main()
