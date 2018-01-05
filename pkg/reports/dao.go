package reports

import (
	"time"
	"../confs"
)

type RTimeType struct {
	Id   int
	Name string
}

func (rtt *RTimeType) Html()string {
	switch rtt.Id {
	case RTT_ID_Real:
	case RTT_ID_Hour:
	case RTT_ID_Day:
	case RTT_ID_Week:
	case RTT_ID_TenDay:
	case RTT_ID_Month:
	case RTT_ID_Quarter:
	case RTT_ID_Year:
	default:
	//default is any

	}

	return ""
}

const (
	RTT_ID_Any     = iota
	RTT_ID_Real
	RTT_ID_Hour
	RTT_ID_Day
	RTT_ID_Week
	RTT_ID_TenDay
	RTT_ID_Month
	RTT_ID_Quarter
	RTT_ID_Year
)

var (
	RTT_Any     = &RTimeType{RTT_ID_Any, confs.GetString("report", "rTimeType", "Any", confs.Locale())}
	RTT_Real    = &RTimeType{RTT_ID_Real, confs.GetString("report", "rTimeType", "Real", confs.Locale())}
	RTT_Hour    = &RTimeType{RTT_ID_Hour, confs.GetString("report", "rTimeType", "Hour", confs.Locale())}
	RTT_Day     = &RTimeType{RTT_ID_Day, confs.GetString("report", "rTimeType", "Day", confs.Locale())}
	RTT_Week    = &RTimeType{RTT_ID_Week, confs.GetString("report", "rTimeType", "Week", confs.Locale())}
	RTT_TenDay  = &RTimeType{RTT_ID_TenDay, confs.GetString("report", "rTimeType", "TenDay", confs.Locale())}
	RTT_Month   = &RTimeType{RTT_ID_Month, confs.GetString("report", "rTimeType", "Month", confs.Locale())}
	RTT_Quarter = &RTimeType{RTT_ID_Quarter, confs.GetString("report", "rTimeType", "Quarter", confs.Locale())}
	RTT_Year    = &RTimeType{RTT_ID_Year, confs.GetString("report", "rTimeType", "Year", confs.Locale())}
)

type ReportInfo struct {
	Id              int
	Name            string
	Profession      string
	Department      string
	Creator         string
	CreatedTime     time.Time
	LastUpdatedTime time.Time
	VisitedCount    int
	TimeType        *RTimeType
}

func GetReportInfo(reportId int) ReportInfo {
	// from db get report by id
	return ReportInfo{Id: 1, Name: "我的报表", TimeType: RTT_Any}

}
