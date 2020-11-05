import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import M from "materialize-css";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  constructor(private _router : Router) { }

  userData = JSON.parse(localStorage.getItem("user"))
  postData = null;
  peopleData = null;

  

onSelect = (id) =>{
    this._router.navigate(["/userprofile",id])
} 

addComment = (comment, postid) => {
    fetch("http://localhost:8080/uploadComment",{
                    method : "post",
                    body : JSON.stringify({
                        comment: comment,
                        email: this.userData.email,
                        name: this.userData.username,
                        postid: postid
                    })
                }).then(response =>response.json())
                .then(data =>{
                    console.log(data)
                    if (data.result === "Success"){
                        location.reload()
                        M.toast({html : "Posted!!!", classes : "rounded green-button"})
                        
                    }else{
                       M.toast({html : "Something Went Wrong", classes : "rounded red-button"})
                    }
                })
}

  ngOnInit(): void {
    fetch(`http://localhost:8080/allPosts`,{
                    method : "get"
            }).then(res => res.json())
            .then(data =>{
                if (Object.keys(data).length !== 0) {
                    this.postData = data.posts
                }
                console.log(this.postData)
                
                fetch("http://localhost:8080/allUser",{
                    method : "get"
                }).then(res=> res.json())
                .then(data=>{
                    if (Object.keys(data).length !== 0) {
                        this.peopleData = data.users
                    }
                })
            })
            .catch(err =>{
                console.log(err)
            })

    
  }

}
