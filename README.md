# Go Learning

This project is created for learning Golang. 

## Projects

1. hello-world:
 - idea: say hello world
 - purpose: project to learning basic concept - print in go 

2. print-file:
 - idea: read and print file content
 - purpose: project to learning basic concepts - read and print in go
 - detail:
  + project read user's input about the file name
  + read the file content
  + print out the file content

  ```go
    go run main.go file.txt
  ```

3. card:
 - idea: play with card deck
 - purpose: project to learning basic concepts about go
 - detail: 
  + `initDeck` : Create a list of playing cards. Essentially an array of strings
  + `printDeck` : Log out the contents of a deck of cards
  + `shuffleDeck` : Shuffles all the cards in a deck
  + `dealDeck` : Create a 'hand' of cards
  + `saveDeck` : Save a list of cards to a file on the local machine
  + `loadDeck` : Load a list of cards from the local machine

4. odd-even:
 - idea: identify odd and even numbers
 - purpose: project to learning basic concepts about go

## Note

To enable a project to access other Go files within the same package, you may need to execute the following command on Windows OS: 

```bash
    go mod init <folder contained>
```

## Reference 

[Udemy](https://www.udemy.com/course/go-the-complete-developers-guide/)

[Github](https://github.com/StephenGrider/GoCasts/)
