#!/usr/bin/env python3

import sys
import os

def print_usage():
    print("Usage: python payment_processor.py [option]")

def print_help():
    print("Options:")
    print("  -h, --help         Show this help message and exit")
    print("  -v, --version      Show version number and exit")
    print("  -f, --file         Specify the input file")
    print("  -o, --output       Specify the output file")

def main():
    if len(sys.argv) == 1:
        print_usage()
        sys.exit(1)

    if sys.argv[1] in ["-h", "--help"]:
        print_help()
        sys.exit(0)

    if sys.argv[1] == "-v" or sys.argv[1] == "--version":
        print("Payment Processor version 1.0")
        sys.exit(0)

    if sys.argv[1] == "-f" or sys.argv[1] == "--file":
        if len(sys.argv) < 3:
            print("Error: input file not specified")
            print_usage()
            sys.exit(1)

        if not os.path.isfile(sys.argv[2]):
            print("Error: input file does not exist")
            print_usage()
            sys.exit(1)

    if sys.argv[1] == "-o" or sys.argv[1] == "--output":
        if len(sys.argv) < 3:
            print("Error: output file not specified")
            print_usage()
            sys.exit(1)

    else:
        print("Error: unknown option")
        print_usage()
        sys.exit(1)

if __name__ == "__main__":
    main()