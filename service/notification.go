package service

import (
	constant "github.com/timakin/airshooter/constant"
	db "github.com/timakin/airshooter/datasource"
	m "github.com/timakin/airshooter/model"
	"time"
)

func EnqueueNotification(notification *m.Notification) (result *m.Notification, err error) {
	createdAt := time.Now().Unix()
	expiry := time.Now().Unix() + constant.NotificationExpiryDuration

	notification.CreatedAt = &createdAt
	notification.Expiry = &expiry
	if result, err = db.InsertNotification(notification); err != nil {
		return nil, err
	}

	return result, nil
}

func GetNotification(id *int64) (result *m.Notification, err error) {
	if result, err = db.SelectNotification(id); err != nil {
		return nil, err
	}
	return result, nil
}

func GetNotifications() (results []*m.Notification, err error) {
	if results, err = db.SelectNotifications(); err != nil {
		return nil, err
	}

	return results, nil
}
