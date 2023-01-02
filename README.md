# gh-stats

Exploring Golang whilist building a tiny command line util.

## Why?

I have this funny idea that it's much harder to actually remove/refactor code rather than producing new code. A couple of references: [1](https://news.ycombinator.com/item?id=10979240) [2](https://blog.codinghorror.com/the-best-code-is-no-code-at-all/)

So I wanted to have a little tool that would help me see how many lines of code I've produced ~~into this cursed world~~ against how many had I removed.

### Next steps

* Can I do the same calculation, but across the whole org?
* Improve command line experience, provide reasonable help.
* Is there a better way to access a token?

### What I've learned

* Github API is quite powerful.
   * Although there are some limitations. Initially I expected that I'd be able to easily fetch something like "all commits made by me in this org", but those were naive expectations as I realise now how much work would that require on the GitHub API side.
* Golang type system appears to be very powerful. It was quite easy to navigate types around and what APIs are available to me.
* [go-github](https://github.com/google/go-github) is very nicely crafted. 
   * There are many great [examples](https://github.com/google/go-github/tree/master/example) in the repo that helped me to get started.
* The concept of [context in Golang](https://pkg.go.dev/context).
* Pointers and derefencing pointers. It was sweet to get to work with those again.
* Basic array functionality is not present by default. So if you need to filter an array, you'd either need to compose a custom function, or use something third-party.
   * Stack Overflow example: [link](https://stackoverflow.com/questions/37562873/most-idiomatic-way-to-select-elements-from-an-array-in-golang)
   * Generally curious why it's this way. Naive guess is that it might somehow optimise compiled executable, as you'd really only include the code you would use in your package. Will ask my golang colleagues about that.
* Built-in [flag](https://pkg.go.dev/flag) package is superb!
* It's funny how much of code you produce while you try things out, but when you are ready to publish you trim everything and the source code looks so simple compared to the time/effort invested. Duh!
