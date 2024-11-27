package redis

/*
一、基于redis实现接口的幂等性验证
	1、前端执行创建或更新操作时，先从后端获取唯一标识符（此时，后端会将生成的唯一标识符存入redis中）
	2、执行创建或更新请求时，将获取到的唯一标识符跟随请求数据传入到后端，后端从redis中查询该唯一标识符，若有，则执行创建或更新操作，并从redis删除该唯一标识符
	3、当有重复请求执行时，后端查询redis获取不到该唯一标识符，则说明该请求是重复的，此时，后端返回错误信息并拒绝执行请求；
*/

/*
缓存一致性解决方案：

一、更新数据库后更新缓存 （写写并发时，会导致数据不一致）
	data=100，
	1、更新数据库：
	线程A -> 线程B 更新数据data-1，data: 100->99->98
	2、更新缓存
	线程B -> 线程A data: 100->98->99
    3、最总导致 database = 98，cache = 99

二、更新数据库前更新缓存 （写写并发时,会导致数据不一致）
	data =100
	1、更新缓存
	先线程A -> 后线程B， data: 100 -> 99 -> 98
	2、更新数据库
	先线程B -> 后线程A  data: 100 -> 98 -> 99
	同一！高并发场景下，由于无法保证更新database和cache的顺序一致，可能存在导致数据不一致问题

所以通常情况下，“更新缓存再更新数据库” 是应该避免使用的一种手段。

三、更新数据库前删除缓存（读写并发时，会导致数据不一致）
	data = 100
	1、先删除缓存 2、再更新数据库

	线程A执行写操作，先删除缓存，

	此时线程B读，由于缓存为空，则读到database=100

	线程A再更新database=99

	线程B再更新cache=100

	最后导致database=99,cache=100

	解决方法：延时删除，线程A更新完数据库后，停留时间N，然后删除缓存，此时缓存为空，数据库数据为最新值
	注意：需要衡量时间N的大小，N太小，导致删除操作在线程B更新cache操纵之前，导致无法删除脏数据，N太大，导致在N期间读取的缓存都是脏数据

四、更新数据库后删除缓存（最佳方案）（缓存失效的前提下，读写并发会导致数据不一致）
	data = 100
	1、先更新数据库 2、再删除缓存

	线程A先写，database = 99，（注：在线程A写期间，再执行读操作，此时database = 99，cache = 100，实际读取到的都是不一致数据）

	线程A再删除缓存

	线程B读，缓存为空，则去读取database = 99，无误！

	注意：
	若缓存失效，

	线程1执行读，无缓存，则读到database = 100

	此时线程2执行写，database 100 -> 99，

	然后线程1更新cache = 100

	最终导致database = 99,cache = 100,数据不一致！

	并发量很大的情况下，缓存失效的可能性不是很大，所以这种情况出现的概率极小。
*/
