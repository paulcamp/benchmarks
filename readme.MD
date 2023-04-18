# Introduction

Welcome to the benchmarks repo.  
This is a collection of folders, each containing a separate readme.MD to state the intent.


# Usage

Open folder in vscode or your choice of golang IDE.

`go mod tidy`   
`go build ./...`

Then run all benchmarks using the IDE.


# Githooks

githooks are present to safeguard non-building commits.  
Read more about [githooks](https://www.atlassian.com/git/tutorials/git-hooks)


to use the githooks, run these commands inside the repo folder: 
```
rm -rf .git/hooks/*
cd .git/hooks
cp -s ../../gittemplates/hooks/pre-commit ./
cp -s ../../gittemplates/hooks/pre-push ./
```

this creates a symbolic link to the repo copies of the hooks and uses them locally.
