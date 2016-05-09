package route

import (
    "github.com/timakin/airshooter/controller"
    "github.com/labstack/echo"
)

func AddNotificationAPI(e *echo.Group) (combined *echo.Group) {
	combined = e.Group("/notifications")
    combined.POST("/enqueue", controller.EnqueueNotification, controller.AuthClient, controller.ValidateRequest)
	combined.PUT("/publish", controller.PublishNotification, controller.ValidateRequest)
	combined.GET("/subscriptions", controller.GetNotifications, controller.ValidateRequest)
	combined.PUT("/subscriptions", controller.MarkReadNotifications, controller.ValidateRequest)
    return combined
}

