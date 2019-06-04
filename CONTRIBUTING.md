# Contributing to Cartelize

Welcome classmates! if you want to contribute to this application, pick a current issue from the issues list, and follow the steps below.

## Recommended workflow

1. Make a [fork](#getting-started-with-github) of the project on GitHub.
2. Clone the fork on your development computer.
3. Create a [branch](#branch-name-guidelines) where you will work.
4. Test and review your work.
5. Commit your changes.
6. Sync your fork on GitHub and the repo on your computer with the latest changes from the poltrack repository (*see How to [sync](#how-to-sync-your-fork) your fork*). This will minimize the effort needed to merge your changes in the poltrack repo.
7. Merge upstream changes if necessary and make sure all tests still pass.
8. Push your work to the remote fork created in Step 2.
9. Create a Pull Request to initiate the integration of your changes in the [Main project](https://github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize).

## Getting Started with GitHub

Please check out one of the getting started guides about GitHub fork / pull requests workflow:

* [Forking project](https://guides.github.com/activities/forking/)
* [Fork a repo](https://help.github.com/en/articles/fork-a-repo)
* [Forking](https://gist.github.com/Chaser324/ce0505fbed06b947d962)

## How to sync your fork

Your fork of the repo can fall behind as more work is done in the original repository. It is always good idea to update your work before starting to work on new issue. The fork can be updated by navigating to your for directory and running the following command:

```
git checkout master --force && git fetch upstream && git merge upstream/master && git push
```

This command assumes you're using unix or unix like environment (macOS, cygwin, WSL, ...). If not you might need to execute commands one by one instead of chaining them with &&.