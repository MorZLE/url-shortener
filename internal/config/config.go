package config

import (
	"flag"
)

func NewConfig() *Config {
	cnf := &Config{}
	return ParseFlags(cnf)
}

type Config struct {
	FlagAddrReq       string
	FlagAddrShortener string
}

func ParseFlags(p *Config) *Config {

	flag.StringVar(&p.FlagAddrReq, "a", ":8080", "address and port to run server")
	flag.StringVar(&p.FlagAddrShortener, "b", "http://127.0.0.1:8080/", "address shortURLer")

	flag.Parse()

	//if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
	//	//	p.FlagAddr = envRunAddr
	//	//}
	//	//
	//	//if ReportInterval := os.Getenv("REPORT_INTERVAL "); ReportInterval != "" {
	//	//	ReportInterval, err := strconv.Atoi(ReportInterval)
	//	//	if err != nil {
	//	//		log.Fatalln("Ошибка преобразования строки в число:", err)
	//	//
	//	//	}
	//	//	p.FlagReportInterval = ReportInterval
	//	//}
	//	//
	//	//if PollInterval := os.Getenv("POLL_INTERVAL "); PollInterval != "" {
	//	//	PollInterval, err := strconv.Atoi(PollInterval)
	//	//	if err != nil {
	//	//		log.Fatalln("Ошибка преобразования строки в число:", err)
	//	//	}
	//	//	p.FlagPollInterval = PollInterval
	//	//}
	return p
}
