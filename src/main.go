package main

import (
	"github.com/phachon/go-logger"
)

func main()  {
	logger := go_logger.NewLogger()

	logger.Detach("console")

	// 命令行输出配置
	consoleConfig := &go_logger.ConsoleConfig{
		Color: true, // 命令行输出字符串是否显示颜色
		JsonFormat: true, // 命令行输出字符串是否格式化
		Format: "", // 如果输出的不是 json 字符串，JsonFormat: false, 自定义输出的格式
	}
	// 添加 console 为 logger 的一个输出
	logger.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)

	// 文件输出配置
	fileConfig := &go_logger.FileConfig {
		Filename : "./test.log", // 日志输出文件名，不自动存在
		// 如果要将单独的日志分离为文件，请配置LealFrimeNem参数。
		LevelFileName : map[int]string {
			logger.LoggerLevel("error"): "./error.log",    // Error 级别日志被写入 error .log 文件
			logger.LoggerLevel("info"): "./info.log",      // Info 级别日志被写入到 info.log 文件中
			logger.LoggerLevel("debug"): "./debug.log",    // Debug 级别日志被写入到 debug.log 文件中
		},
		MaxSize : 1024 * 1024,  // 文件最大值（KB），默认值0不限
		MaxLine : 100000, // 文件最大行数，默认 0 不限制
		DateSlice : "d",  // 文件根据日期切分， 支持 "Y" (年), "m" (月), "d" (日), "H" (时), 默认 "no"， 不切分
		JsonFormat: true, // 写入文件的数据是否 json 格式化
		Format: "", // 如果写入文件的数据不 json 格式化，自定义日志格式
	}
	// 添加 file 为 logger 的一个输出
	logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)


	logger.Info("this is a info log!")
	logger.Errorf("this is a error %s log!", "format")
}
