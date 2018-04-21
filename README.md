# MIT-Distributed-Systems
This is a solution about 6.824 labs.

### lab1 MapReduce
> MapReduce 是一种编程模型，用于处理和生成大数据集。

#### lab1.1 Map/Reduce input and output 
* do_map函数实现
do_map执行一个map任务。从给定的文件中读取内容，通过用户自定义的map函数从文件内容中生成一系列中间键值对。通过对键进行哈希求余后生成R个文件对应R个Reduce任务。
* do_reduce函数实现
do_reduce执行一个reduce任务。读取M个map处理结果文件，并将得到的键值对按键排序后，对同一个键的值进行聚合，将<key, list<value>>格式的参数传给用户自定义的Reduce函数进行处理，最后将处理结束后的结果存储到输出文件。

#### lab1.2 Single-worker word count
* mapF函数实现
mapF对文件内容按照单词分词，并按照<word, 1>格式生成KeyValue类型。
* reduceF函数实现
reduceF函数遍历value值将同一键的值相加得到一个单词出现的总次数。

#### lab1.3 distributing MapReduce tasks
schedule调度任务给空闲的worker。
获取空闲worker，并分配任务给它，使任务并发进行。等待任务都执行结束后，schedule函数返回。

#### lab1.4 Handling worker failures
处理worker错误，对于出现错误的任务重新分配一个worker再次执行该任务

#### lab1.5 Inverted index generation
先从文档中分词生成倒排表；reduce阶段对文档进行去重排序，并按照相应的格式生成结果。

实验一结果
```bash
bash ./test-mr.sh

==> Part I
ok      mapreduce   3.544s

==> Part II
Passed test

==> Part III
ok      mapreduce   3.884s

==> Part IV
ok      mapreduce   4.233s

==> Part V (inverted index)
Passed test

```
