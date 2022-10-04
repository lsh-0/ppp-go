# PPP, Pretty Perfect Publishing

An experiment in transitioning a microservices architecture to a monolith, preserving the good qualities and suppressing the bad.

This experiment uses Go.

Problems that can't be (reasonably) overcome:

* values can't be null
* members require capitalisation for export
* complete absence of dynamism leads to elaborate hacks
    - see `pprint`
* development feedback loop is kind of long

## usage

    go build
    ./elife
