package store

type FeedStore interface {
	// Save 保存订阅信息
	Save(feedUrl string) error

	// GetAll 获取订阅信息
	GetAll() ([]string, error)

	// Delete 删除订阅信息
	Delete(feedUrl string) error
}
