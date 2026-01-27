# Concepts of Programming Languages - Course Materials

## Chapters

| Chapter | Topic | PDF | Slides |
|---------|-------|-----|--------|
| 1.0 | About | [PDF](1.0-About.pdf) | [Slides](1.0-About.slide) |
| 1.1 | Overview | [PDF](1.1-Overview.pdf) | [Slides](1.1-Overview.slide) |
| 1.2 | Introduction to Golang | [PDF](1.2-Introduction%20to%20Golang.pdf) | [Slides](1.2-Introduction%20to%20Golang.slide) |
| 2.0 | Go Programming - Basics and OOP | [PDF](2.0-Go%20Programming%20-%20Basics%20and%20OOP.pdf) | [Slides](2.0-Go%20Programming%20-%20Basics%20and%20OOP.slide) |
| 3.0 | Go Programming - OOP | [PDF](3.0-Go-Programming-OOP.pdf) | [Slides](3.0-Go-Programming-OOP.slide) |
| 4.0 | Go Programming - Parser | [PDF](4.0-Go-Programming-Parser.pdf) | [Slides](4.0-Go-Programming-Parser.slide) |
| 5.0 | Functional Programming | [PDF](5.0-Functional-Programming.pdf) | [Slides](5.0-Functional-Programming.slide) |
| 6.0 | Concurrent Programming | [PDF](6.0-Concurrent-Programming.pdf) | [Slides](6.0-Concurrent-Programming.slide) |
| 7.0 | Distributed Programming - Raft | [PDF](7.0-Distributed-Programming-Raft.pdf) | [Slides](7.0-Distributed-Programming-Raft.slide) |
| 8.0 | Systems Programming | [PDF](8.0-Systems-Programming.pdf) | [Slides](8.0-Systems-Programming.slide) |
| 9.0 | Enterprise Programming - Modules | [PDF](9.0-Enterprise-Programming-Modules.pdf) | [Slides](9.0-Enterprise-Programming-Modules.slide) |

## How to Use the Slides

These slide files (.slide) are intended to be viewed with the 'present' tool.
Use go install to install 'present':

```bash
go install golang.org/x/tools/cmd/present@latest
```

You can start the slideshow by running the present tool:

```bash
present
```

This will start a webserver. The current directory and all subdirs are served.

Or use the provided Makefile from the project root:

```bash
make slideshow
```

This will also start the present tool.

## Additional Resources

- [exercises/](exercises/) - Practice exercises
- [studywork/](studywork/) - Student project templates
- [talks/](talks/) - Additional presentations
