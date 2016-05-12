package pili

import (
	"fmt"
	"net/http"
	"time"
)

// Config 用于配置 Client.
type Config struct {
	AccessKey string
	SecretKey string
}

// Client 代表一个 pili 用户的客户端.
type Client struct {
	*rpc
	mac *mac
}

// New 初始化 Client.
func New(accessKey, secretKey string, tr http.RoundTripper) *Client {
	mac := &mac{accessKey, []byte(secretKey)}
	rpc := newRPC(newTransport(mac, tr))
	return &Client{
		rpc: rpc,
		mac: mac,
	}
}

// RTMPPublishURL 生成 RTMP 推流地址.
// expireAfterSeconds 表示 URL 在多久之后失效.
func (c *Client) RTMPPublishURL(domain, hub, streamKey string, expireAfterSeconds int64) string {
	expire := time.Now().Unix() + expireAfterSeconds
	path := fmt.Sprintf("/%s/%s?e=%d", hub, streamKey, expire)
	token := c.mac.Sign([]byte(path))
	return fmt.Sprintf("rtmp://%s%s&token=%s", domain, path, token)
}

// RTMPPlayURL 生成 RTMP 直播地址.
func RTMPPlayURL(domain, hub, streamKey string) string {
	return fmt.Sprintf("rtmp://%s/%v/%v", domain, hub, streamKey)
}

// HLSPlayURL 生成 HLS 直播地址.
func HLSPlayURL(domain, hub, streamKey string) string {
	return fmt.Sprintf("http://%s/%s/%s.m3u8", domain, hub, streamKey)
}

// HDLPlayURL 生成 HDL 直播地址.
func HDLPlayURL(domain, hub, streamKey string) string {
	return fmt.Sprintf("http://%s/%s/%s.flv", domain, hub, streamKey)
}

// SnapshotPlayURL 生成截图直播地址.
func SnapshotPlayURL(domain, hub, streamKey string) string {
	return fmt.Sprintf("http://%s/%s/%s.jpg", domain, hub, streamKey)
}
