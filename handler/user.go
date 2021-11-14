package handler

import (
	"merchant-service/auth"
	"merchant-service/entity"
	"merchant-service/user"
	"merchant-service/utils/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.UserService
	authService auth.Service
}

func NewUserHandler(userService user.UserService, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// CREATE NEW USER OR REGISTER
func (h *userHandler) RegisterUserHandler(c *gin.Context) {
	var inputUser entity.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	// max len password 6
	if err := helper.ValidatePassword(inputUser.Password); err != nil {

		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"error": "error validate password length < 6"})

		c.JSON(400, responseError)
		return
	}

	newUser, err := h.userService.SaveNewUser(inputUser)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create new user", 201, "Status Created", newUser)
	c.JSON(201, response)
}

//LOGIN User
func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var inputLoginUser entity.LoginUserInput

	if err := c.ShouldBindJSON(&inputLoginUser); err != nil {
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	userData, err := h.userService.LoginUser(inputLoginUser)

	if err != nil {
		responseError := helper.APIResponse("Input data error", 401, "bad request", gin.H{"errors": err.Error()})

		c.JSON(401, responseError)
		return
	}

	token, err := h.authService.GenerateTokenUser(userData.Id)
	if err != nil {
		responseError := helper.APIResponse("Internal server error", 500, "internal server error", gin.H{"error": err.Error()})

		c.JSON(401, responseError)
		return
	}
	response := helper.APIResponse("success", 200, "success", gin.H{"token": token, "id": userData.Id})
	c.JSON(200, response)
}

// SHOW ALL USER
func (h *userHandler) ShowAllUserHandler(c *gin.Context) {
	userUser, err := h.userService.GetAllUser()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all user", 200, "status OK", userUser)
	c.JSON(200, response)
}

// FIND User BY ID
func (h *userHandler) GetUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	userUser, err := h.userService.GetUserByID(id)
	if err != nil {
		responseError := helper.APIResponse("input params error", 400, "bad request", gin.H{"errors": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success get user by ID", 200, "success", userUser)
	c.JSON(200, response)
}

// UPDATE User BY ID
func (h *userHandler) UpdateUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	var updateUserInput entity.UpdateUserInput

	if err := c.ShouldBindJSON(&updateUserInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	userData := c.MustGet("currentUser").(gin.H)
	userID := userData["user_id"]

	if id != userID {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	userUser, err := h.userService.UpdateUserByID(id, updateUserInput)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"error": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success update user by ID", http.StatusOK, "success", userUser)
	c.JSON(http.StatusOK, response)
}

// DELETE User BY ID
func (h *userHandler) DeleteUserByIDHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	userData := c.MustGet("currentUser").(gin.H)
	userID := userData["user_id"]

	if id != userID {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user ID not authorize"})

		c.JSON(401, responseError)
		return
	}

	userUser, err := h.userService.DeleteUserByID(id)

	if err != nil {
		responseError := helper.APIResponse("error bad request delete user", 400, "error", gin.H{"error": err.Error()})

		c.JSON(400, responseError)
		return
	}

	response := helper.APIResponse("success delete user by ID", 200, "success", userUser)
	c.JSON(200, response)
}

// CREATE NEW Outlet
func (h *userHandler) CreateOutletUserHandler(c *gin.Context) {

	// file, err := c.FormFile("Picture")

	userData := c.MustGet("currentUser").(gin.H)
	userID := userData["user_id"].(string)

	if len(userID) == 0 {
		responseError := helper.APIResponse("Unauthorize", 401, "error", gin.H{"error": "user not authorize / not login"})

		c.JSON(401, responseError)
		return
	}
	// path := fmt.Sprintf("images/picture-%d-%s", userData, file.Filename)

	// err = c.SaveUploadedFile(file, path)

	// userID := strconv.Itoa(userData)

	var inputOutlet entity.OutletInput

	if err := c.ShouldBindJSON(&inputOutlet); err != nil {

		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIResponse("input data required", 400, "bad request", gin.H{"errors": splitError})

		c.JSON(400, responseError)
		return
	}

	// pathOutletSave := "https://merchant.herokuapp.com/" + path

	newOutlet, err := h.userService.CreateOutletUser(inputOutlet, userID)
	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}
	response := helper.APIResponse("success create new Outlet", 201, "Status OK", newOutlet)
	c.JSON(201, response)
}

// SHOW ALL Outlet
func (h *userHandler) ShowAllOutletUserHandler(c *gin.Context) {
	outlet, err := h.userService.ShowAllOutletUser()

	if err != nil {
		responseError := helper.APIResponse("internal server error", 500, "error", gin.H{"errors": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse("success get all Outlet", 200, "status OK", outlet)
	c.JSON(200, response)
}
