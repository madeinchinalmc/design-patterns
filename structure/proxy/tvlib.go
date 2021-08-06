package proxy

import "fmt"

// 远程服务接口
type ThirdPartyTVLib interface {
	ListVideo()
	GetVideoInfo(id int) string
	DownloadVideo(id int)
}

// 第三方TV
type ThirdPartyTVClass struct {
}
var videos = map[int]string{
	1:"亮剑",
	2:"觉醒年代",
	3:"建军大业",
}

func (t *ThirdPartyTVClass)ListVideo() {
	fmt.Printf("发送一个获取视频列表的api请求 %v", videos)
}

func (t *ThirdPartyTVClass)GetVideoInfo(id int) string{
	v,ok := videos[id]
	if ok {
		return v
	}
	return ""
}


func (t *ThirdPartyTVClass)DownloadVideo(id int){
	v,ok := videos[id]
	if ok {
		fmt.Printf("正在下载电影%v",v)
		return
	}
	fmt.Printf("未找到电影")
}

//本地TV缓存代理
type CachedTVClass struct {
	service ThirdPartyTVClass
	NeedReset bool
	videoCache map[int]string
}

func NewCachedTVClass(thirdService ThirdPartyTVClass)CachedTVClass{
	cache := make(map[int]string)
	cache[1] = "亮剑"
	return CachedTVClass{
		videoCache:cache,
		NeedReset: false,
		service: thirdService,
	}
}

func (t *CachedTVClass)ListVideo() {
	if t.videoCache == nil || t.NeedReset{
		fmt.Println("需要向第三方发送远程请求")
		t.service.ListVideo()
		return
	}
	fmt.Printf(" %v", t.videoCache)
}

func (t *CachedTVClass)GetVideoInfo(id int) string{
	v,ok :=  t.videoCache[id]
	if ok {
		fmt.Println("在本地缓存中找到")
		return v
	}
	return t.GetVideoInfo(id)
}


func (t *CachedTVClass)DownloadVideo(id int){
	v,ok :=  t.videoCache[id]
	if ok {
		fmt.Printf("正在从本地缓存下载电影%v",v)
		return
	}
	t.DownloadVideo(id)
}

// 客户端使用
type TVManager struct {
	Service CachedTVClass
}

func NewTVManager(service CachedTVClass)TVManager{
	return TVManager{Service: service}
}

func(t *TVManager) RenderVideoPage(id int){
	info := t.Service.GetVideoInfo(id)
	fmt.Printf("RenderVideoPage %v",info)
}

func(t *TVManager) renderListPanel(){
	t.Service.ListVideo()
}

type Application struct{

}

func(a *Application) ApplicationRun(){
	aTVService := ThirdPartyTVClass{}
	aTVProxy :=  NewCachedTVClass(aTVService)
	manager :=  NewTVManager(aTVProxy)
	manager.renderListPanel()
	manager.RenderVideoPage(1)
}