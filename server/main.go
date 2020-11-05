package main

import (
	"log"
	"google.golang.org/grpc"
	"github.com/gin-gonic/gin"
	"../grpcServer/proto"
	"./model"
	"github.com/gin-contrib/cors"
)

type LoginData struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main()  {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	Client := proto.NewAddAppServiceClient(conn)

	g := gin.Default()

	g.Use(cors.Default())

	//register User
	g.POST("/signup", func (c *gin.Context)  {
		var userData model.User
	
		if err := c.ShouldBindJSON(&userData); err != nil {
			log.Fatal(err)
		}
	
		var registerRequest = &proto.RegisterUserRequest{
			Username: userData.Username,
			Email: userData.Email,
			Password: userData.Password,
			ImageUrl: userData.ImageUrl,
		}
	
		registerResponse, e := Client.Register(c, registerRequest)
		if e != nil {
			log.Fatal(e)
		} 
		c.JSON(200, registerResponse)
	})

	//Login User

	g.POST("/signin", func (c *gin.Context)  {
		var userData LoginData
	
		if err := c.ShouldBindJSON(&userData); err != nil {
			log.Fatal(err)
		}
	
		log.Println(userData.Password)

		var loginRequest = &proto.LoginUserRequest{
			Email: userData.Email,
			Password: userData.Password,
		}
	
		loginResponse, e := Client.Login(c, loginRequest)
		if e != nil {
			log.Fatal(e)
		} 
		c.JSON(200, loginResponse)
	})

	//Uploading Posts
	g.POST("/uploadPost", func (c *gin.Context) {
		var postData model.Post

		if err := c.ShouldBindJSON(&postData); err != nil {
			log.Fatal(err)
		}
		var request = &proto.PostRequest{
			Description: postData.Description,
			ImageUrl: postData.ImageUrl,
			Email: postData.Email,
			Option: postData.Option,
			Username: postData.Username,
		}
		response, e := Client.UploadPost(c, request)
		if e != nil {
			log.Fatal(e)
		}
		c.JSON(200, response)
	})

	//Upload comments
	g.POST("/uploadComment", func (c *gin.Context) {
		var commentData model.Comment

		if err := c.ShouldBindJSON(&commentData); err != nil {
			log.Fatal(err)
		}

		var request = &proto.CommentRequest{
			Comment: commentData.Comment,
			Email: commentData.Email,
			Name: commentData.Name,
			Postid: commentData.PostId,
		}
		response, e := Client.UploadComment(c, request)
		if e != nil {
			log.Fatal(e)
		}
		c.JSON(200, response)

	})

	//Make Friend
	g.POST("/makeFriend", func (c *gin.Context) {
		var friendData model.Friend

		if err := c.ShouldBindJSON(&friendData); err != nil {
			log.Fatal(err)
		}

		var request = &proto.FriendRequest{
			FriendMail: friendData.FriendMail,
			FriendName: friendData.FriendName,
			FriendImage: friendData.FriendImage,
			Mymail: friendData.Mymail,
			MyimageUrl: friendData.MyimageUrl,
			Myusername: friendData.Myusername,
			Status: friendData.Status,
		}

		response, e := Client.MakeFriend(c, request)
		if e != nil {
			log.Fatal(e)
		}

		c.JSON(200, response)

	})

	//get []posts by mail id
	g.GET("/userPost", func (c *gin.Context) {
		var mail = c.Request.URL.Query().Get("mail")

		var request = &proto.GetPostRequest{Email: mail}

		response, err := Client.GetUserPost(c, request)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, response)

	})

	//get User by mail id
	g.GET("/getUser", func (c *gin.Context) {
		var mail = c.Request.URL.Query().Get("mail")
		var request = &proto.GetUserRequest{Email: mail}

		response, err := Client.GetUser(c, request)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, response)
	})

	g.GET("/userFriend", func (c *gin.Context) {
		var mail = c.Request.URL.Query().Get("mail")
		var request = &proto.GetUserFriendRequest{Email: mail}

		response, err := Client.GetUserFriend(c, request)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, response)
	})

	g.GET("/allUser", func (c *gin.Context) {
		var request = &proto.Empty{}

		response, err := Client.AllPeople(c, request)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, response)
	})

	g.GET("/getRequest", func (c *gin.Context) {
		var mail = c.Request.URL.Query().Get("mail")
		var request = &proto.RequestFromUsers{Email: mail}

		response, err := Client.RequestFromUser(c, request)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, response)
	})

	//Accept Friend Request
	g.GET("/acceptFriendRequest", func (c *gin.Context) {
		var friendMail = c.Request.URL.Query().Get("friendmail")
		var mymail = c.Request.URL.Query().Get("mymail")

		var request = &proto.AcceptRequest{
			FriendMail: friendMail,
			Mymail: mymail,
		}

		response, err := Client.AcceptFriendRequest(c, request)
		if err != nil {
			log.Fatal(err)
		} 
		c.JSON(200, response)
	})

	//Delete a friend
	g.DELETE("/deleteFriend", func (c *gin.Context) {
		var friendMail = c.Request.URL.Query().Get("friendmail")
		var mymail = c.Request.URL.Query().Get("mymail")

		var request = &proto.DeleteFriendRequest{
			FriendMail: friendMail,
			Mymail: mymail,
		}

		response, err := Client.DeleteFriend(c, request)
		if err != nil {
			log.Fatal(err)
		} 
		c.JSON(200, response)
	})

	//Get All Posts
	g.GET("/allPosts", func (c *gin.Context) {
		var request = &proto.Empty{}

		response, err := Client.AllPost(c, request)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, response)
	})


	if e := g.Run(":8080"); err != nil {
		log.Fatal(e)
	}

}