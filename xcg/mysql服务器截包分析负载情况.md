有几种工具可以帮助你截取 `tcpdump` 捕获的 MySQL 数据包，并解析生成负载报告。以下是一些常见的工具：

1. **tcpdump + Wireshark：** 
   - 使用 `tcpdump` 截获 MySQL 服务器的数据包。
   - 将捕获的数据包文件（例如：`capture.pcap`）导入到 Wireshark 中。
   - 使用 Wireshark 的过滤器和分析功能查看 MySQL 协议的详细信息。

2. **Percona Toolkit：**
   - Percona Toolkit 提供了 `pt-query-digest` 工具，用于分析 MySQL 查询日志。
   - 使用 `tcpdump` 截获 MySQL 数据包。
   - 使用 `pt-query-digest` 分析 `tcpdump` 输出，生成负载报告。
   - 示例：`tcpdump -i eth0 -s 65535 -x -nn -q -tttt -c 1000 port 3306 > capture.txt`

3. **tshark (命令行版 Wireshark)：**
   - 使用 `tcpdump` 截获 MySQL 数据包。
   - 使用 `tshark` 命令行工具来解析 `tcpdump` 的输出文件。
   - 示例：`tshark -r capture.pcap -q -z "proto,colinfo,proto,mysql" > mysql_report.txt`

4. **MySQL Performance Schema：**
   - 使用 MySQL Performance Schema 收集 MySQL 服务器的性能数据。
   - 通过配置 Performance Schema，可以收集有关查询、连接、等待事件等方面的信息。
   - 使用 Performance Schema 查询相关表来获取详细信息。

选择适合你需求的工具取决于你对数据分析和报告的具体要求。Wireshark 提供了可视化的界面和功能强大的过滤和分析工具，而 Percona Toolkit 和 tshark 则提供了命令行工具，适用于自动化和脚本化的场景。 MySQL Performance Schema 则是一个 MySQL 内建的性能监测工具，可以在数据库层面进行更详细的分析。