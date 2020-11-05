import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit {

  constructor() { }

  userData = JSON.parse(localStorage.getItem("user"));
  postData = null
  friendData = null

  ngOnInit(): void {

    fetch(`http://localhost:8080/userPost?mail=${this.userData.email}`,{
                    //headers : { "Content-Type" : "application/json" ,"Authorization" : "Bearer "+localStorage.getItem("token")},
                    method : "get"
            }).then(res => res.json())
            .then(data =>{
                //console.log(data)
                this.postData = data.postdata
                /*fetch(`http://localhost:5000/userfriends/${this.userData.userid}`,{
                    headers : { "Content-Type" : "application/json" ,"Authorization" : "Bearer "+localStorage.getItem("token")},
                    credentials : "include",
                    method : "get",
                }).then(res =>res.json())
                .then(friends =>{
                    //console.log(friends)
                    this.friendData = friends
                })*/
            })
            .catch(err =>{
                console.log(err)
            })

  }

}
