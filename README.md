# ConvertDate

**Assumptions/Decisions**
* For simplicity, I made the assumption it files/data would be straight ASCII but this could be easily changed to support UTF-8 encoding, etc. 
* Since ordinals don't affect the values parsed out for the date, I chose not to validate ordinals against the preceding number.  I.e. 1nd and 4rd will be 1 and 4 respectively and not produce an error or warning.
  * I originally had support in the regex for validating them, but Go's regex does not support positive/negative look behinds
    * For those curious, here is the regex I was going to use: `((?<=1)st|(?<=2)nd|(?<=3)rd|(?<=[0456789])th)`
* I made the assumption that the formats could vary wildly where the values for dates aren't necessarily all together, for example
    * `On this 27th day of May in the year of 2023` can be a valid date (see Date Formats, below)
* I decided to log any invalid dates (i.e. Feb 31st) to the console as a warning to the user in case they wanted to investigate.
* I decided to do a CLI application, but based on a class/struct that is doing conversion.  This class/struct could easily be included in an API end point or any other application, if needed.
* I decided to support converting a string and converting file to add flexibility to the application.
* I made assumptions for one and two digit day/month formats.  I made m2/d2 (See "Date Formats" below) REQUIRE 2 digits to match, where m1/d1 will work with either.  i.e. m1/d1 will match: 1, 01, 11

**Installing Go**
* If you need to install go, you can install it from here: https://go.dev/doc/install
* If you are on a Mac and want to install via Homebrew, this might help: https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5

**Building/Testing/Running**
* Clone the repo
* Change directory into the cloned directory
* Run unit tests (optional): `go test -v ./...`
* Build: `go build .`
* Execute: `./ConvertDate inputfile inputformat outputfile outputformat`
  * Example: `./ConvertDate input.txt "ml d2or, y4" output.txt "y4-m2-d2"`
    * Convert `December 21st, 2023` to `2023-12-21`
  * Note: You can use the same file name for input and output if you want convert "in place"

**Date formats**
* You can get the valid specifiers for the date formats from the command line usage if you run the command no arguments.
* Here are the format specifiers:
  * `d1` - Non zero padded month number (i.e. 2)
  * `d2` - Zero padded day number (i.e. 02)
  * `m1` - Non zero padded month number (i.e. 1)
  * `m2` - Zero padded month number (i.e. 01)
  * `ml` - Short month name (i.e Jan)
  * `ms` - Long month name (i.e January)
  * `or` - Required ordinal string (i.e. st, nd, rd, th)
  * `oo` - Optional ordinal string (i.e. st, nd, rd, th)
  * `y2` - Year number without century (i.e. 06)
  * `y4` - Year number including century (i.e. 2006)
* Examples:
  * `y4-m2-d2` would parse/generate things like `2023-12-21`
  * `ml d1or, y4` would parse/generate things like `December 21st, 2023`
  * `this d1or day of ml in the year of y4` would parse/generate things like `this 27th day of May in the year of 2023`

**Future Enhancements**
* As a possible future enhancement, I'd have allowed the user to start the tool in an interactive mode where they could issue commands to do some of the following:
  * SetInputFormat <format> - Allow the user to change the input format at any time
  * SetOutputFormat <format> - Allow the user to change the output format at any time.
  * Convert <string> - Convert the string and display the results
  * ConvertFile <in> <out> - Convert a file on disk directly
  * Perhaps the ability to load a file into memory, perform multiple conversions then write out the results.
