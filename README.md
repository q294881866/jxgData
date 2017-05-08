# jxgData
数据模拟生产，按时间，可中断并配置
# 用户使用
 * 编译后运行可执行文件，并运行
   * 只有一个参数：数据生成保存目录
   * 如果参数为空：默认保存目录：可执行文件目录+/data
   * 中断恢复
     * 在数据保存目录的上级目录，自动创建 abort 文件
     * abort 文件保存启动以来生成的分钟数，线程安全
     * 如：100分钟，表明生成了100分钟内的数据
     * 中断数据部分丢失：数据批量刷导致内存数据丢失；先增加abort分钟数后生成文件导致数据不一定生成完
     * 中断恢复，在数据保存目录的上级目录读取 abort 文件，接着生成
 * 执行逻辑
   * 5000台机器，每台机器一秒钟生成一条数据，每分钟生成一个文件
   * 根据每个文件的大小（固定），磁盘写入速率，将5000个写文件均匀分布
   * 这里随机sleep 12ms 将写文件均匀
   * 文件重命名，预防句柄泄露
     * windows可以启动批量命名任务，提供性能
     * linux只能文件写完关闭后，重命名
 * 用户拓展，task/task.go
   * 指定BeginTime，设置起始时间
   * 指定ThreadNumber，协程并发数量，默认300，根据操作系统和cpu
   * 修改WorkTime 方法，制订数据生成的阶段，比如：Am8:00--Pm9:00
   * 用户简单修改程序，通过输入参数制订数据起始时间、结束时间（推荐分钟数）、并行文件生成数量（并发量）
   
