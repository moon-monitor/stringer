package vobj

// NoticeType 通知类型
//
//go:generate stringer -type=NoticeType -linecomment -output=notice_type.string.go
type NoticeType int8

const (
	NoticeTypeUnknown NoticeType = 1 << iota // 未知
	NoticeTypeSystem                         // 系统
	NoticeTypeUser                           // 用户
	NoticeTypeAdmin                          // 管理员
	NoticeTypeAll                            // 所有
)
