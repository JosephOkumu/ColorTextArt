
# ColorTextArt

This project takes input text, an optional color flag and letters to be colored from command-line arguments, and prints out its colorized ASCII-Art graphical representation in the terminal as provided for with the color flag and the letters to be colored. If the options are not provided, it prints the graphical art in the terminal's default color.



## Documentation

This part contains instructions on how to make use of our program.

### Usage

To use this program, you need to have the latest version of Go installed in your system.
After cloning the projects repository, navigate to the ascii-art-color directory and execute the following command in your terminal.
```bash
go run . [OPTION] [STRING]

EX: go run . --color=<color> <letters to be colored> "something"
```
Provide the color flag and enter the color you want as an English named color, RGB value or HEX value to display the colorized art in the terminal.

For example:


```bash
go run main.go --color=red hel "hello guys" # to colorize h, e and l's in "hello guys"
go run . --color="RGB(255, 0, 0)" "hello guys" # to colorize everything in "hello guys"
go run main.go "hello guys" # for no color
go run . --color="#00ff00" u "hello guys" # to colorize u in "hello guys"
```


## Authors

- [@josotieno](https://learn.zone01kisumu.ke/git/josotieno/)
- [@jwambugu](https://learn.zone01kisumu.ke/git/jwambugu)
- [@kada](https://learn.zone01kisumu.ke/git/kada)


## License

[MIT](https://choosealicense.com/licenses/mit/)


## Credits

[Zone01Kisumu](https://zone01kisumu.ke)
