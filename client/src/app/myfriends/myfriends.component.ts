import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import M from "materialize-css"

@Component({
  selector: 'app-myfriends',
  templateUrl: './myfriends.component.html',
  styleUrls: ['./myfriends.component.css']
})
export class MyfriendsComponent implements OnInit {

  constructor(private _router : Router) { }
  friendsData = null
  userData = JSON.parse(localStorage.getItem("user"))

  deleteHandler = (friendmail, mymail) =>{
    fetch(`http://localhost:8080/deleteFriend?friendmail=${friendmail}&mymail=${mymail}`,{
                method : "delete"
        }).then(res => res.json())
        .then(data =>{
          if (data.result === "Success") {
            M.toast({html : "Removed from friend list", classes : "rounded green-button"})
            this._router.navigate(["/addfriends"])
          }
            
        })
        .catch(err =>{
            console.log(err)
        })
      }

  ngOnInit(): void {
    fetch(`http://localhost:8080/userFriend?mail=${this.userData.email}`,{
                method : "get"
        }).then(res => res.json())
        .then(data =>{
          console.log(data)
            if (Object.keys(data).length != 0) {
              this.friendsData = data.friends
            }
            
        })
        .catch(err =>{
            console.log(err)
        })
  }

}
