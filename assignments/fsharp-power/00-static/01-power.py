#!/usr/bin/env python3

import re, sys, os.path

# exit 1 for continue
# exit 2 for break

POWER = re.compile(r"let\s+power\s+.*?=", re.DOTALL)
POWER_WITH_ARGS = re.compile(r"let\s+power(\s+\w+){2}\s+=", re.DOTALL)

def main():
  path = os.path.join(sys.argv[1], "1G.fsx")
  contents = open(path).read()

  matches = POWER.findall(contents)
  len_matches = len(matches)

  if len(matches) == 0:
    print("Vi kan ikke finde funktionen power.")
    print("Har du måske glemt at erklære den?")
    sys.exit(2)

  if len(matches) > 1:
    print("De lader til at times42 forekommer mere end en gang!")
    print("Slet det ene forekomst og prøv igen.")
    sys.exit(2)

  # Der lader til at være et times42 funktion..
  power_decl = matches[0]

  match = POWER_WITH_ARGS.match(power_decl)
  if not match:
    print("power skulle have netop 2 argumenter: m og n.")
    print("Resultatet skulle være m opløftet i n'te potens.")
    sys.exit(2)

if __name__ == "__main__":
  main()
