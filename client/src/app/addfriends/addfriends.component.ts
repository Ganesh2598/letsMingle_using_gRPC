import { Component, OnInit } from '@angular/core';
import M from "materialize-css";
import { Router } from "@angular/router" 

@Component({
  selector: 'app-addfriends',
  templateUrl: './addfriends.component.html',
  styleUrls: ['./addfriends.component.css']
})
export class AddfriendsComponent implements OnInit {

  constructor(private _router : Router) { }
  friendsData = null;
  userData = JSON.parse(localStorage.getItem("user"))
  requestData = null;

  acceptHandler = (mymail) => {
    fetch(`http://localhost:8080/acceptFriendRequest?friendmail=${this.userData.email}&mymail=${mymail}`)
    .then(res=>res.json())
    .then(data=>{
      console.log(data)
      M.toast({html : "Added as Friend", classes : "rounded green-button"})
    })
  }

  addHandler = (username, imageUrl, email)=> {
    fetch("http://localhost:8080/makeFriend",{
                    method : "post",
                    body : JSON.stringify({
                        friendMail: email,
                        friendImage: imageUrl,
                        friendName: username,
                        mymail: this.userData.email,
                        myimageUrl: this.userData.imageUrl,
                        myusername: this.userData.username,
                        status: "Pending"
                    })
            }).then(res => res.json())
            .then(data =>{
              console.log(data)
                M.toast({html : "Request Send", classes : "rounded green-button"})
                this._router.navigate(["/myfriends"])
            })
            .catch(err =>{
                console.log(err)
            })
  }

  ngOnInit(): void {
    fetch("http://localhost:8080/allUser",{
                    method : "get"
            }).then(res => res.json())
            .then(data =>{
                console.log(data)
                this.friendsData = data.users
                fetch(`http://localhost:8080/getRequest?mail=${this.userData.email}`)
                .then(res => res.json())
                .then(data => {
                  console.log(Object.keys(data).length)
                  if (Object.keys(data).length != 0) {
                    this.requestData = data.requests
                  }
                  console.log(this.requestData)
                  
                })
            })
            .catch(err =>{
                console.log(err)
            })
  }

}
