# Benchmarks for the paper "Multiscale event detection using convolutional quadtrees and adaptive geogrids"

In order to evaluate convolutional quadtree (https://github.com/visheratin/conv-tree) against the standard quadtree we performed a series of experiments, where we built ConvTrees and quadtrees for three datasets containing Instagram posts in the New York City area and calculated metrics that define characteristics of trees. We used following metrics: depth of tree branch, number of points in the leaf, total number of leafs ($N$) and number of jumps from leafs to their spatial neighbors. It is worth mentioning that the last metric is especially important because in cases of applying tree structures to spatial data processing it is often required to perform search for data in spatially close areas.

To reproduce the results of benchmarks presented in the paper, you need to clone this repository, build it using `go build` command and execute the generated file.

Output of the execution will be the following:
```go
Checking small file.
Total number of points - 7825
ConvTree stats
Tree depth:
Min, max - 2 9
Mean - 6.031161473087819
Standard deviaion - 1.3936606461346635
Number of points in leaf:
Min, max - 0 505
Mean - 11.08356940509915
Standard deviaion - 27.228532943373324
Number of leafs - 706
Number of transitions to neighbour leafs:
Min, max - 2 13
Total - 17266
Mean - 3.735612289052358
Standard deviation - 2.09072095019556
-----
QuadTree stats
Tree depth:
Min, max - 2 8
Mean - 6.524038461538462
Standard deviation - 1.480527124543579
Number of points in leaf:
Min, max - 0 492
Mean - 9.405048076923077
Standard deviation - 24.791658926829722
Number of leafs - 832
Number of transitions to neighbour leafs:
Min, max - 2 16
Total - 25194
Mean - 4.206010016694491
Standard deviation - 2.918718089262513

=====================

Checking medium file.
Total number of points - 79195
ConvTree stats
Tree depth:
Min, max - 3 12
Mean - 7.447932387564141
Standard deviaion - 1.6077698886192822
Number of points in leaf:
Min, max - 0 8491
Mean - 23.904316329610626
Standard deviaion - 167.6879062348485
Number of leafs - 3313
Number of transitions to neighbour leafs:
Min, max - 2 15
Total - 83338
Mean - 3.8046932067202337
Standard deviation - 2.1996259797777933
-----
QuadTree stats
Tree depth:
Min, max - 3 8
Mean - 7.0470321931589535
Standard deviation - 1.10757312802899
Number of points in leaf:
Min, max - 0 8473
Mean - 19.918259557344065
Standard deviation - 154.13079019660415
Number of leafs - 3976
Number of transitions to neighbour leafs:
Min, max - 2 16
Total - 123518
Mean - 4.237324185248713
Standard deviation - 2.7929339397826918

=====================

Checking large file.
Total number of points - 279773
ConvTree stats
Tree depth:
Min, max - 3 12
Mean - 7.706101403609281
Standard deviaion - 1.456179384052614
Number of points in leaf:
Min, max - 0 42610
Mean - 40.07061014036093
Standard deviaion - 555.2718925777671
Number of leafs - 6982
Number of transitions to neighbour leafs:
Min, max - 2 16
Total - 178496
Mean - 3.8533742066404733
Standard deviation - 2.294822452980898
-----
QuadTree stats
Tree depth:
Min, max - 3 8
Mean - 7.296025474702206
Standard deviation - 0.8724354834937058
Number of points in leaf:
Min, max - 0 42594
Mean - 32.995990093171365
Standard deviation - 502.8797504278614
Number of leafs - 8479
Number of transitions to neighbour leafs:
Min, max - 2 16
Total - 267148
Mean - 4.26972254187444
Standard deviation - 2.7793369182313135

```
