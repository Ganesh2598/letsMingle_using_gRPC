package main

import (
	"log"
	"golang.org/x/net/context"
	"./proto"
	"../server/model"
	"../server/database"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"time"
	"net"
)

type server struct {}

type Post struct {
	Description string
	ImageUrl string
	Email string
}

func main()  {
	database.Connection()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal(err)
	}
	srv :=  grpc.NewServer()
	proto.RegisterAddAppServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		log.Fatal(e)
	}
}

//Registeration
func(s *server) Register(ctx context.Context, request *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error)  {
	var userData model.User
	var available []model.User
	userData.Username = request.GetUsername()
	userData.Email = request.GetEmail()
	userData.Password = request.GetPassword()
	userData.ImageUrl = request.GetImageUrl()

	database.Db.Where("Email=?", userData.Email).Find(&available)

	if (len(available) > 0) {
		return &proto.RegisterUserResponse{Result: "User Already exist"}, nil
	} else {
		database.Db.Create(&userData)
		return &proto.RegisterUserResponse{Result: "Success"}, nil
	}
}

//Login functionality
func(s *server) Login(ctx context.Context, request *proto.LoginUserRequest) (*proto.LoginUserResponse, error) {
	var userData []model.User

	database.Db.Where("Email=?", request.GetEmail()).Find(&userData)

	if (len(userData) <= 0) {
		return &proto.LoginUserResponse{
			Result: "Failed",
		}, nil
	} else {
		log.Println(request.GetPassword())
		if (request.GetPassword() == userData[0].Password) {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user": userData[0].Username,
				"email": userData[0].Email,
				"exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
			})
			tokenString, err := token.SignedString([]byte("secret"))
			if err != nil {
				log.Fatal(err)
			}
			return &proto.LoginUserResponse{
				Id: 0,
				Username: userData[0].Username,
				Email: userData[0].Email,
				ImageUrl: userData[0].ImageUrl,
				Token: tokenString,
				Result: "Logged in",
			}, nil
		} else {
			return &proto.LoginUserResponse{
				Result: "Failed",
			}, nil
		}
	}
}

//Uploading Post Functionality
func(s *server) UploadPost(ctx context.Context, request *proto.PostRequest) (*proto.PostResponse, error)  {
	var postData model.Post
	postData.Description = request.GetDescription()
	postData.ImageUrl = request.GetImageUrl()
	postData.Email = request.GetEmail()
	postData.Option = request.GetOption()
	postData.Username = request.GetUsername()
	log.Println(postData)

	database.Db.Create(&postData)

	return &proto.PostResponse{Result: "Success"}, nil
}

//Uploading Comments Functionality
func(s *server) UploadComment(ctx context.Context, request *proto.CommentRequest) (*proto.CommentResponse, error) {
	var commentData model.Comment
	commentData.Comment = request.GetComment()
	commentData.Email = request.GetEmail()
	commentData.Name = request.GetName()
	commentData.PostId = request.GetPostid()
	log.Println(commentData)

	database.Db.Create(&commentData)

	return &proto.CommentResponse{Result: "Success"}, nil
}

//Make a Friend
func(s *server) MakeFriend(ctx context.Context, request *proto.FriendRequest) (*proto.FriendResponse, error) {
	var friendData model.Friend
	friendData.FriendMail = request.GetFriendMail()
	friendData.FriendName = request.GetFriendName()
	friendData.FriendImage = request.GetFriendImage()
	friendData.Mymail = request.GetMymail()
	friendData.MyimageUrl = request.GetMyimageUrl()
	friendData.Myusername = request.GetMyusername()
	friendData.Status = request.GetStatus()

	database.Db.Create(&friendData)

	return &proto.FriendResponse{Result: "success"}, nil
}

//get user by mail id
func(s *server) GetUser(ctx context.Context, request *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	var usermail = request.GetEmail()
	var userData model.User

	database.Db.Where("Email=?", usermail).Find(&userData)

	return &proto.GetUserResponse{
		Username: userData.Username,
		Email: userData.Email,
		ImageUrl: userData.ImageUrl,
	}, nil
}

//get []userpost by mail id
func(s *server) GetUserPost(ctx context.Context, request *proto.GetPostRequest) (*proto.GetPostResponse, error) {
	var usermail = request.GetEmail()
	var postData []model.Post
	var postss []model.Post

	database.Db.Where("Email=?", usermail).Find(&postData)
	database.Db.Preload("Comments").Find(&postss)

	log.Println(postss)

	posts := []*proto.PostRequest{}

	for _,post := range postData {
		data := new(proto.PostRequest)
		log.Println(post.Email)
		data.Email = post.Email
		data.Description = post.Description 
		data.ImageUrl = post.ImageUrl
		data.Option = post.Option
		data.Username = post.Username
		posts = append(posts, data)
	}

	log.Println(posts)

	return &proto.GetPostResponse{
		Postdata: posts,
	} , nil

}

//get User friends
func(s *server) GetUserFriend(ctx context.Context, request *proto.GetUserFriendRequest) (*proto.GetUserFriendResponse, error)  {
	var friends []model.Friend
	var mail = request.GetEmail()
	database.Db.Where("Mymail=? and Status=?",mail,"Accepted").Find(&friends)

	friendsData := []*proto.FriendRequest{}
	for _, friend := range friends {
		data := new(proto.FriendRequest)
		data.FriendMail = friend.FriendMail
		data.FriendImage = friend.FriendImage
		data.FriendName = friend.FriendName
		data.Mymail = friend.Mymail
		friendsData = append(friendsData, data)
	}

	return &proto.GetUserFriendResponse{
		Friends: friendsData,
	}, nil

}

//get all people,s in letsmingle
func(s *server) AllPeople(ctx context.Context, request *proto.Empty) (*proto.AllPerson, error)  {
	var users []model.User
	database.Db.Find(&users)

	usersData := []*proto.People{}
	for _, user := range users {
		data := new(proto.People)
		data.Username = user.Username
		data.Email = user.Email
		data.ImageUrl = user.ImageUrl
		usersData = append(usersData, data)
	}

	return &proto.AllPerson{
		Users: usersData,
	}, nil

}

//Request from Users
func(s *server) RequestFromUser(ctx context.Context, request *proto.RequestFromUsers) (*proto.RequestFromUserResponse, error)  {
	var requests []model.Friend
	var mail = request.GetEmail()

	database.Db.Where("friend_mail=? and Status=?", mail, "Pending").Find(&requests)

	friends := []*proto.FriendRequest{}
	for _, friend := range requests {
		data := new(proto.FriendRequest)
		data.Mymail = friend.Mymail
		data.MyimageUrl = friend.MyimageUrl
		data.Myusername = friend.Myusername
		friends = append(friends, data)
	}

	return &proto.RequestFromUserResponse{
		Requests: friends,
	}, nil

}

//Accept Friend Request
func(s *server) AcceptFriendRequest(ctx context.Context, request *proto.AcceptRequest) (*proto.AcceptResponse, error) {
	var friendMail = request.GetFriendMail()
	var mymail = request.GetMymail()

	database.Db.Model(&model.Friend{}).Where("friend_mail=? and mymail=?", friendMail, mymail).Update("Status", "Accepted")
	return &proto.AcceptResponse{Result: "Success"}, nil
}

//Delete a friend (or) unfollow him
func(s *server) DeleteFriend(ctx context.Context, request *proto.DeleteFriendRequest) (*proto.DeleteFriendResponse, error) {
	var friendMail = request.GetFriendMail()
	var mymail = request.GetMymail()

	database.Db.Where("friend_mail=? and mymail=?", friendMail, mymail).Unscoped().Delete(&model.Friend{})

	return &proto.DeleteFriendResponse{Result: "Success"}, nil
}

//All posts
func(s *server) AllPost(ctx context.Context, request *proto.Empty) (*proto.MultiPostReturing, error) {
	var posts []model.Post
	database.Db.Preload("Comments").Find(&posts)

	lists := []*proto.SinglePostReturing{}

	for _, post := range posts {
		data := new(proto.SinglePostReturing)
		data.Postid = uint32(post.ID)
		data.Description = post.Description
		data.ImageUrl = post.ImageUrl
		data.Email = post.Email
		data.Username = post.Username
		data.Comments = []*proto.CommentRequest{}
		for _, comment := range post.Comments {
			data1 := new(proto.CommentRequest)
			data1.Comment = comment.Comment
			data1.Name = comment.Name
			data1.Email = comment.Email
			data1.Postid = comment.PostId
			data.Comments = append(data.Comments, data1)
		}
		lists = append(lists, data)
	}

	return &proto.MultiPostReturing{Posts: lists}, nil
}