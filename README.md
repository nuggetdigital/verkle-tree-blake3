# verkle-tree-blake3

[![ci](https://github.com/nuggetdigital/verkle-tree-blake3/workflows/ci/badge.svg)](https://github.com/nuggetdigital/verkle-tree-blake3/actions/workflows/ci.yml) ![](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)

A ***very very* experimental** implementation of Verkle 🌳🌲

This is a simple adaption of [`go-verkle`](https://github.com/gballet/go-verkle) that replaces Golang's [`crypto/sha256`](https://golang.org/pkg/crypto/sha256/) (SHA2) with [`BLAKE3-256`](https://github.com/BLAKE3-team/BLAKE3), for faster leaf hashing, maintaining an adequate 128-bit collision resistance.

## readables

+ [math.mit.edu paper](https://math.mit.edu/research/highschool/primes/materials/2018/Kuszmaul.pdf)

+ [math.mit.edu slides](https://math.mit.edu/research/highschool/primes/materials/2019/conf/12-5-Kuszmaul.pdf)

+ [vitalik.ca blogpost](https://vitalik.ca/general/2021/06/18/verkle.html)

+ [notes.ethereum.org blogpost](https://notes.ethereum.org/nrQqhVpQRi6acQckwm1Ryg)

## test

```
$ go test .
```

## bench

```
$ go test . -bench Bench
```

<!-- ## Performance measurements

This table measures the time it takes to calculate the root commitment of the current state of an Ethereum network:

|Network|Node size|Parallel?|Storage?|DB?|BLS library|Time|# accounts|#slots|
|-------|---------|---------|--------|---|-----------|----|----------|------|
|Mainnet|1024|No|No|No|Herumi|3h30m24.663s|114215117|0|
|Mainnet|1024|No|Yes|No|Herumi|16h36m7.043s|114215117|400223042|
|Mainnet|1024|Yes|Yes|No|Herumi|10h1m34.056s|114215117|400223042|
|Mainnet|1024|Yes|Yes|Yes|Herumi|12h47m22.309s|114215117|400223042|
|Mainnet|256|Yes|Yes|No|Herumi|18h45m21.182s|114215117|400223042|
|Mainnet|256|Yes|Yes|Yes|Herumi|9h8m24.923s|114215117|400223042|
|Mainnet|256|Yes|Yes|Yes|Kilic|16h23m11.616s|114215117|400223042|
|Görli|1024|No|Yes|No|Herumi|~30min|1104810|35900044|

## Size measurements

Values with experimental encoding

|Network|Node size|Verkle tree size in DB|
|-------|---------|----------------------|
|Mainnet|1024|68G|
|Mainnet|256|32G|
|Görli|1024|3.6G|


## Insertion/update benchmarks

|Initial tree size|Action|Width|Average time per insert|
|-----------------|------|-----|-----------------------|
|1M leaves|Insert 10K leaves|1024|25ms|
|1M leaves|Update 10K leaves|1024|7ms|
|1M leaves|Insert 10K leaves|256|1.5ms|

## Proof generation/verification benchmarks

|Initial tree size|Action|Width|Average time|
|-----------------|------|-----|------------|
|10K leaves|Proof for 1 leaf|1024|0.93s|
|10K leaves|Verify proof|1024|4ms|
 -->
