package settings

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"blog/utils"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email"`
}

// 登录响应
type LoginResponse struct {
	Token        string           `json:"token"`
	RefreshToken string           `json:"refresh_token"`
	UserModel    models.UserModel `json:"user"`
}

// Login 用户登录
func (Settingsapi) Login(c *gin.Context) {
	var req LoginRequest
	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		res.ResultFail("参数错误", c)
		return
	}
	var user models.UserModel
	result := global.DB.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			res.ResultFail("用户名不存在！", c)
			return
		} else {
			res.ResultFail("数据库查询失败", c)
			return
		}

	}
	if CheckPassword(user.Password, req.Password) {
		res.ResultFail("密码错误", c)
		return
	}
	// 生成 token
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Username)
	response := LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
		UserModel:    user,
	}
	c.Set("user", user)
	c.Set("user_id", user.ID)
	res.ResultOk(response, "登录成功", c)
}

// Register 用户注册
func (Settingsapi) Register(c *gin.Context) {
	var req LoginRequest
	// 绑定 JSON 到结构体
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		res.ResultFail("参数错误", c)
		return
	}
	var existingUser models.UserModel
	result := global.DB.Where("username = ?", req.Username).First(&existingUser)
	if result.Error == nil {
		res.ResultFail("用户名已经存在", c)
		return
	} else {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			res.ResultFail("数据库查询失败", c)
			return
		}
	}
	pass, err := HashPassword(req.Password)
	if err != nil {
		res.ResultFail("密码加密失败", c)
		return
	}
	// 创建新用户
	newUser := models.UserModel{
		Username: req.Username,
		Password: pass,
		Email:    req.Email,
	}
	result = global.DB.Create(&newUser)
	if result.Error != nil {
		res.ResultFail("用户创建失败", c)
		return
	}
	res.ResultOk(newUser, "用户注册成功", c)
}

// HashPassword 加密密码
func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPassword 验证密码
func CheckPassword(hashedPassword, password string) bool {
	// 基本验证：哈希值应该以 $2a$、$2b$ 或 $2y$ 开头
	if len(hashedPassword) < 60 || !strings.HasPrefix(hashedPassword, "$2a$") {
		fmt.Printf("无效的哈希值格式: %s\n", hashedPassword)
		return true
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Printf("密码验证失败: %v\n", err)
		return true
	}
	return false
}
