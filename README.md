# MIT-Distributed-Systems
This is a solution about 6.824 labs.

### lab1 MapReduce
> MapReduce 是一种编程模型，用于处理和生成大数据集。

#### lab1.1 Map/Reduce input and output 
* do_map函数实现
do_map执行一个map任务。从给定的文件中读取内容，通过用户自定义的map函数从文件内容中生成一系列中间键值对。通过对键进行哈希求余后生成R个文件对应R个Reduce任务。
* do_reduce函数实现
do_reduce执行一个reduce任务。读取M个map处理结果文件，并将得到的键值对按键排序后，对同一个键的值进行聚合，将<key, list<value>>格式的参数传给用户自定义的Reduce函数进行处理，最后将处理结束后的结果存储到输出文件。