package gesa

import "time"

// Timestamp は esa API のレスポンスに含まれる時刻情報を表す型
type Timestamp int64

// Time はタムスタンプを *time.Time に変換して返す
// ts が nil の場合は nil を返す
func (ts *Timestamp) Time() *time.Time {
	if ts == nil {
		return nil
	}

	t := time.Unix(int64(*ts), 0)
	return &t
}

// SafeTimestamp は *ts を int64 に変換して返す
// ts が nil の場合は 0 を返す
func (ts *Timestamp) SafeTimestamp() int64 {
	if ts == nil {
		return 0
	}

	return int64(*ts)
}

// RateLimitInformation は RateLimit の情報を保持する構造体
type RateLimitInformation struct {
	Limit     int        // 15 分間にリクエストできる回数の上限
	Remaining int        // 制限までに実行可能なリクエスト回数
	Reset     *Timestamp // 制限がリセットされる時間のタイムスタンプ
}

// PageNumber は esa API におけるページネーションのパラメータ、
// およびレスポンスを表す型
type PageNumber int

// NewPageNumber は *PageNumber を生成して返す
// n が 0 未満の場合は PageNumber(0) を返す
func NewPageNumber(n int) *PageNumber {
	if n < 0 {
		n = 0
	}
	p := PageNumber(n)
	return &p
}

// IsNull は pn が nil かどうかを判定する
func (pn *PageNumber) IsNull() bool {
	return pn == nil || *pn == 0
}

// SafeInt は pn を int に変換する
// pn が nil の場合は 0 を返す
func (pn *PageNumber) SafeInt() int {
	if pn == nil {
		return 0
	}

	return int(*pn)
}
