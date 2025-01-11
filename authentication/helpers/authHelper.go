package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) error {
	userType := c.GetString("user_type")
	if userType != role {
		return errors.New("user type mismatch")
	}
	return nil
}

func MatchUserTypeToUid(c *gin.Context, userId string) error {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	if userType == "USER" && uid != userId {
		return errors.New("user type is USER but UID is not equal to userId")
	}
	return CheckUserType(c, userType)
}
