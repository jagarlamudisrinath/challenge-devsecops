# DevSecOps Challenge

You mission, [should you choose to accept it](https://www.youtube.com/watch?v=0TiqXFssKMY), is to step in the shoes of our new junior
DevSecOps recruit.

We have a some code waiting for you to review, of a simple app designed to manage a database of our secret agents, "Users" we call them.
We hope to use this app as a gateway so they can authenticate to our systems in the future.

You are now free to have a thorough look at the code and inspect the work done so far.

The goals are as follow:
- For any issue you spot, propose a fix. But the best would be to implement a fitting solution.
- Identify and implement any missing bits and best practices generally associated with your trade (being a DevSecOps, that is).

## Requirements

Special attention will be given to:

- Requirement compliance.
- Overall code formatting, readability and elegance.
- The quality of the README.md and WORK_SUMMARY.md files

### Must Haves

- Good documentation
- Everything "as-code" (infrastructure, pipeline, **everything**)
- A **README.md** file is expected, to detail the chosen solution, how to run it and a short section about how to develop in your repository (git workflow, process, etc..).
- A **WORK_SUMMARY.md** file is expected, to detail how you embrace that problem and sequentially, what was your solving approach

If you don't want to use GitHub, that's okay - you can use any hosted solution as long as it stays **private**. Please send an archive of your work, including the `.git` folder, to [hiring@loftorbital.com](mailto:hiring@loftorbital.com).

## The `challenge` app

### Minimum Requirements

This app that needs reviewing is written in `Golang`, to build and run it, you will need:

- docker and docker-compose
- golang (1.16)

### Build and Run

To build the Docker image:

```shell
make build
```

To run the stack with its Postgres SQL instance:

```shell
docker-compose up -d
```


## Questions?

Just [reach out](mailto:hiring@loftorbital.com)!

---

This README file will self-destruct in five seconds. Good luck!

![img](./img/mission-impossible.jpg)
