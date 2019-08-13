#!.env/bin/python3

'''A file with helper functions.'''

import sys


def print_error(*err_msg):
    ''' Print an error. '''

    err_out = "Error: "

    for err in err_msg:
        err_out += (err + '\n')

    print(err_out)


def are_args_valid(args):
    ''' Check number of cli arguments. '''

    if len(args) != 3:
        print_error("Missing arguments.",
                    "Syntax: create_project <projectname> <language>")
        sys.exit(1)

    if not str(sys.argv[1]).strip():
        print_error("Error: Projectname needs to be specified.")
        sys.exit(2)
    if not str(sys.argv[2]).strip():
        print_error("Error: Language needs to be specified.")
        sys.exit(3)