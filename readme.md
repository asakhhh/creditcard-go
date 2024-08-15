[36m[1mcreditcard[0m[0m is a tool for handling credit card number(s).
[1musages:[0m
./creditcard option [flags] [args]
./creditcard --help		- outputs help message for the program.
./creditcard option --help	- outputs help message for the option.

[1mOptions:
[0m[34mvalidate[0m - option for checking the validity of one or more credit cards. Evaluation is based on the Luhn algorithm.
[1musages:[0m
./creditcard validate [number]...	- prints OK or INCORRECT for entered card number(s).
./creditcard validate --stdin		- card numbers are read from standard input.
./creditcard validate --help


[34mgenerate[0m - option for generating valid credit card numbers based on a given template. Prints all combinations by substituting digits for all asterisks[31m (maximum 4)[0m.
[1musages:[0m
./creditcard generate <template>	 - outputs all possible combinations that satisfy the given template.
./creditcard generate --pick <template>	 - randomly chooses one number that satisfies the given template.
./creditcard generate --help


[34minformation[0m - option for getting info about the card number's validity, brand and issuer.
[1musages:[0m
./creditcard information [--brands=FILE] [--issuers=FILE] [numbers]...
./creditcard information --help

[1mFlags:[0m
[1m--brands=FILE[0m	- reads the list of brands from a file.
[1m--issuers=FILE[0m	- reads the list of issuers from a file.
[1m--help[0m		- prints help message for this option.


[34missue[0m - option for generating a random card number for a given brand and issuer.
[1musages:[0m
./creditcard issue [--brands=FILE] [--issuers=FILE] [--brand=] [--issuer=]
./creditcard issue --help

[1mFlags:[0m
[1m--brands=FILE[0m	- reads the list of brands from a file.
[1m--issuers=FILE[0m	- reads the list of issuers from a file.
[1m--brand=[0m	- specifies the brand of the card.
[1m--issuer=[0m	- specifies the issuer of the card.
[1m--help[0m		- prints help message for this option.

