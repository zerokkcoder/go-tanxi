package captcha

import "go-tanxi/pkg/session"

// SessionStore 实现 base64Captcha.Store interface
type SessionStore struct {
	KeyPrefix string
}

// Set 实现 base64Captcha.Store interface 的 Set 方法
func (s *SessionStore) Set(key string, value string) error {
	session.Put(s.KeyPrefix+key, value)
	return nil
}

// Get 实现 base64Captcha.Store interface 的 Get 方法
func (s *SessionStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := session.Get(key)
	if clear {
		session.Forget(key)
	}
	return val.(string)
}

// Verify 实现 base64Captcha.Store interface 的 Verify 方法
func (s *SessionStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
