package main

type Platform interface {
	GetPlatformName() string
	CalculateTrafficScore(c Content) float64
}

type DiliDili struct{}

func (d DiliDili) GetPlatformName() string { return "d站" }
func (d DiliDili) CalculateTrafficScore(c Content) float64 {
	return float64(c.playCount)*0.3 + float64(c.likeCount)*2 + float64(c.commentCount)*4 + float64(c.saveCount)*6
}

type TikTook struct{}

func (t TikTook) GetPlatformName() string { return "跳音" }
func (t TikTook) CalculateTrafficScore(c Content) float64 {
	return float64(c.playCount)*0.5 + float64(c.likeCount)*3 + float64(c.commentCount)*2 + float64(c.shareCount)*5
}

type Xls struct{}

func (b Xls) GetPlatformName() string { return "小蓝书" }
func (b Xls) CalculateTrafficScore(c Content) float64 {
	return float64(c.playCount)*0.3 + float64(c.likeCount)*2 + float64(c.commentCount)*3 + float64(c.saveCount)*8
}

type SlowFoot struct{}

func (s SlowFoot) GetPlatformName() string { return "慢脚" }
func (s SlowFoot) CalculateTrafficScore(c Content) float64 {
	return float64(c.playCount)*0.4 + float64(c.likeCount)*2 + float64(c.commentCount)*5 + float64(c.shareCount)*3
}

func GetPlatform(name string) Platform {
	switch name {
	case "d站":
		return DiliDili{}
	case "跳音":
		return TikTook{}
	case "小蓝书":
		return Xls{}
	case "慢脚":
		return SlowFoot{}
	default:
		return nil
	}
}
