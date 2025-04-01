package handlers

import (
	"RestuarantBackend/interfaces"
	"encoding/csv"
	"encoding/json"
	"net/http"

	custom "RestuarantBackend/custom"
	dto "RestuarantBackend/models/dto"
	models "RestuarantBackend/models/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service interfaces.UserInterface
}

// Constructor for denpendency injection
func NewUserController(service interfaces.UserInterface) *UserController {
	if service == nil {
		panic("NewUserController service is nil")
	}
	return &UserController{service: service}
}

// User Sign Up
func (u *UserController) Register(c *gin.Context) {
	var request dto.SignupRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": custom.ErrBadRequest})
		return
	}
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": custom.ErrInternalServer})
		return
	}
	_, err = u.service.Register(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Success! Please wait..."})
}

// User Login
func (u *UserController) Login(c *gin.Context) {
	var request *dto.LoginRequest
	// gan gia tri tu request vao request
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": custom.ErrBadRequest})
		return
	}
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": custom.ErrInternalServer})
		return
	}
	result, err := u.service.Login(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Login Success", "Data": result})
}

// Token Login
func (u *UserController) LoginToken(c *gin.Context) {
	var request *dto.LoginRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": custom.ErrBadRequest})
		return
	}
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": custom.ErrInternalServer})
		return
	}
	result, err := u.service.TokenLogin(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save Token into Cookie
	c.SetCookie("token", result, 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{"Message": "Login Success", "Data": result})
}

// Googgle Login
func (u *UserController) LoginGoogle(c *gin.Context) {
	request, err := c.Cookie("ggrequest")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot read request from cookies"})
		return
	}
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Service doesn't work!"})
		return
	}
	var data *models.LoginGoogleRequest
	err = json.Unmarshal([]byte(request), &data)
	result, err := u.service.LoginGoogle(data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Something wrong when login!"})
		return
	}
	c.SetCookie("token", result, 3600, "/", "localhost", false, false)
	c.JSON(http.StatusOK, gin.H{"Message": "Login Success", "Data": result})
	c.SetCookie("ggrequest", "", -1, "/", "localhost", false, true)
}

// Udpdate User Information
func (u *UserController) Update(c *gin.Context) {
	var request *dto.UserUpdateRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Service is not initialized"})
		return
	}
	result, err := u.service.Update(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Update Success", "Data": result})
}

// Get all User
func (u *UserController) GetAllUSerPagingList(c *gin.Context) {
	var request *models.PagingRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in Step 1 in GetAllUserPagingList"})
		return
	}

	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error in GetAllUserPagingList"})
		return
	}
	result, err := u.service.PagingListAllUser(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error in Step 2 in GetAllUserPaging List"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": result})
}

// Export User Into CSV File
func (u *UserController) ExportUserCSVFile(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Diposition", "attachment; filename =user.csv")

	// Create a CSV Writer
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	// Write Header
	writer.Write([]string{"ID", "PhoneNumber", "Email", "FullName", "CreatedAt", "UpdatedAt", "DeletedAt", "Role", "Point"})

	// Get User Data
	users, err := u.service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Cannot get user data in Export Function"})
		return
	}
	for _, user := range users {
		writer.Write([]string{strconv.Itoa(user.Id), user.PhoneNumber, user.Email, user.FullName, user.CreatedAt, user.UpdatedAt, user.DeletedAt.String, strconv.Itoa(user.Role), strconv.Itoa(user.Point)})
	}
}

// Block or UnBlock User
func (u *UserController) BlockOrUnblockUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
	}
	if u.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Service does not active !"})
	}
	result, err := u.service.BlockOrUnBlockUser(&userId)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, err)
	}
	c.JSON(http.StatusOK, gin.H{"Message": result})
}
