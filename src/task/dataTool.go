package task

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	//	"sync/atomic"
	"time"
)

// 随机数生成器
var randCount int64 = 0

// num 保存多少小数+1
func randomNum(low int32, high int32, num int) float32 {
	randCount += 1
	r := rand.New(rand.NewSource(randCount))
	gap := high - low
	res := r.Float32()*float32(gap) + float32(low)
	pow10_n := math.Pow10(num)
	res = float32(math.Trunc((float64(res)+0.5/pow10_n)*pow10_n) / pow10_n)
	return res
}


// 开机时间
var startTime time.Time

// 一天多长时间
var DayTime = time.Date(2017, 1, 2, 1, 1, 1, 1, time.Local).UnixNano() - time.Date(2017, 1, 1, 1, 1, 1, 1, time.Local).UnixNano()
var HourTime = time.Date(2017, 1, 1, 2, 1, 1, 1, time.Local).UnixNano() - time.Date(2017, 1, 1, 1, 1, 1, 1, time.Local).UnixNano()

type Data struct {
	// 时间戳
	ts string
	// == 按天计算
	// 每天开机时间
	startHour string
	//停机时间
	stopHour string
	//运行时间
	runTime string
	//待机时间
	restTime string
	//效率
	radioWork string
	//X轴电机平滑电流
	ampereX string
	//Y轴电机平滑电流
	ampereY string
	//Z轴电机平滑电流
	ampereZ string
	//B轴电机平滑电流
	ampereB string
	//C轴电机平滑电流
	ampereC string

	//X轴电机转速
	speedX string
	//Y轴电机转速
	speedY string
	//Z轴电机转速
	speedZ string
	//B轴电机转速
	speedB string
	//C轴电机转速
	speedC string
}

// 生产当前时间戳字符串
func Timestamp(now time.Time) string {
	// 时间戳
	nowString := now.String()
	index := strings.LastIndex(nowString, "+")
	return string([]rune(nowString)[:index-1])
}

// 数据生产器
func DataProducter(now time.Time) Data {
	data := Data{}
	// 时间戳
	data.ts = Timestamp(now)

	// 0:00开机
	startTime = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	start := float64(now.UnixNano()-startTime.UnixNano()) / float64(HourTime)

	// 换算成小时
	data.startHour = fmt.Sprintf("%.3f", start)
	stop := float64(24) - start
	data.stopHour = fmt.Sprintf("%.3f", stop)

	// 机器运行信息 5000台机器
	RestTime += float64(0.00000001) * float64(randomNum(0, 4, 9))
	//	fmt.Println(RestTime)
	run := start - RestTime
	if run < float64(0) {
		run = float64(0)
		RestTime = float64(0)
	}
	data.restTime = fmt.Sprintf("%.3f", RestTime)
	data.runTime = fmt.Sprintf("%.3f", run)
	data.radioWork = fmt.Sprintf("%.3f", run/start)
	if strings.EqualFold("NaN", data.radioWork) {
		data.radioWork = "0"
	}
	//电流
	data.ampereX = fmt.Sprintf("%.3f", randomNum(0, 30, 3))
	data.ampereY = fmt.Sprintf("%.3f", randomNum(0, 20, 3))
	data.ampereZ = fmt.Sprintf("%.3f", randomNum(0, 17, 3))
	data.ampereB = fmt.Sprintf("%.3f", randomNum(0, 2, 3))
	data.ampereC = fmt.Sprintf("%.3f", randomNum(0, 5, 3)/3)
	// 转速
	data.speedX = fmt.Sprintf("%.3f", randomNum(-1000, 1000, 3))
	data.speedY = fmt.Sprintf("%.3f", randomNum(-400, 400, 3))
	data.speedZ = fmt.Sprintf("%.3f", randomNum(-700, 400, 3))
	data.speedB = fmt.Sprintf("%.3f", randomNum(-5, 2, 3))
	data.speedC = fmt.Sprintf("%.3f", randomNum(-2, 2, 3))
	return data
}

func Data2String(split string, data Data) string {
	return data.ts + split + data.startHour + split + data.stopHour + split + data.runTime + split + data.restTime + split + data.radioWork + split + data.ampereX + split + data.ampereY + split + data.ampereZ + split + data.ampereB + split + data.ampereC + split + data.speedX + split + data.speedY + split + data.speedZ + split + data.speedB + split + data.speedC
}
