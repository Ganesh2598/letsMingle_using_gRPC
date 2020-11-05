import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.css']
})
export class UserProfileComponent implements OnInit {

  constructor(private _route : ActivatedRoute) { }
  userData = null;
  postData = null;
  friendData = null;


  ngOnInit(): void {
    let mail = this._route.snapshot.paramMap.get("mail");
    fetch(`http://localhost:8080/getUser?mail=${mail}`,{
                    //headers : { "Content-Type" : "application/json" ,"Authorization" : "Bearer "+localStorage.getItem("token")},
                    method : "get"
            }).then(res => res.json())
            .then(data =>{
              console.log(data)
                this.userData = data
                    fetch(`http://localhost:8080/userPost?mail=${mail}`,{
                    //headers : { "Content-Type" : "application/json" ,"Authorization" : "Bearer "+localStorage.getItem("token")},
                    method : "get"
                }).then(res => res.json())
                .then(data =>{
                  this.postData = data.postdata
                    //console.log(data)
                })
            })
            .catch(err =>{
                console.log(err)
            })
  }

}
