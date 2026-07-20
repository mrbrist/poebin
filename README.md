# POEBIN
![example](https://github.com/mrbrist/poebin/blob/main/example.png "example image")

A Path of Exile build sharing tool

Written in Go + React (Typescript)

## Motivation
This is one of several projects on my journey to learning the Go + Typescript tech stack, this project was chosen to test different storage solutions for large text files.

The solution that I landed on was Cloudfare R2 which works perfectly for this project as it allows for single file storage and is cheaper than alternatives like AWS.

# How it Works
This project functions as a pastebin site where the user pastes in a large string of text (exported from Path of Building) and the result is displayed in a nice way so the user can share builds with friends.

There are a few other sites that do the same thing which allows me to measure the success of this project by looking at how others have achieved the same result.

## Quick Start
1. Clone the project
2. Setup the .env files
3. Run `make dev -j` in the root folder 

## Usage
1. Go to the website on the port that is printed into the terminal when you start the web app
2. Paste the build code from Path of Building

## Contributing
If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

