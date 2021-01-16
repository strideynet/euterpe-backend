# go-monorepo-microservices

An attempt to hash out and develop my own personal preferences for developing Go microservices. There may be some weird
stuff here on branches, where I try out different structures in order to try and come to a reasoned preference. 


## Mono, mono, mono

Monorepo. Monomod. Monobinary.

### Monorepo

As many will be familiar with, this is a monorepo. All code related to the project is within one Git repository.

My reasoning:

- Easier to develop, and review, and release, changes that have occurred across multiples components.
- The counterpoints can be worked around.

Counterpoints:

- Sometimes more difficult to use common CI solutions with.
- Heavier git tree and more files. More files, more changes means more time spent downloading the repo.

### Monomod

A single go.mod file.

For this project, I am assuming that a relatively small team is working on it. Whilst there are arguable advantages
to splitting out the shared libraries to their own independently versioned and released packages for larger teams and
more established organisations, this seems to me to slow down initial development by small teams significantly.

It is simple enough to split these libraries out to their own independently versioned packages in future, and hence
makes more sense to wait for it to become clear this is needed.

### Monobinary (and *monodockerimage*)

A single compiled binary and docker image.

Again, whilst the project remains small and I don't intend on releasing each service individually, a single binary and
image is easier to work with. It also seems again that the work to split this out at a later date is not that huge, so
it is better to wait for it to be clear that this is needed.

## Other shit

I've included far more than usual README.MDs in packages to explain logic/reasoning.

Haven't written a makefile or scripts yet, so here's some dumped commands I've been using.

```shell script
mockery --all --keeptree 
```

```shell script
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative service.emoji/proto/v1/emoji.proto
```
