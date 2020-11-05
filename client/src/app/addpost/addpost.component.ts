import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import M from "materialize-css";

@Component({
  selector: 'app-addpost',
  templateUrl: './addpost.component.html',
  styleUrls: ['./addpost.component.css']
})
export class AddpostComponent implements OnInit {

  constructor(private _router : Router) { }

  postData = {
    description : "",
    imageUrl : "",
    option : "",
    email: ""
  }

  userData = JSON.parse(localStorage.getItem("user"))

  ngOnInit(): void {
  }

  uploadFile = (file) => {
    this.postData.imageUrl = file
  }

  optionChange = (option) => {
    this.postData.option = option
  }

  onSubmit = () => {
    const formData = new FormData();
    formData.append("file",this.postData.imageUrl);
    formData.append("upload_preset","letsmingle");
    formData.append("cloud_name","dvezzidsw");
    fetch("https://api.cloudinary.com/v1_1/dvezzidsw/image/upload",{
        method : "post",
        body : formData,
    })
        .then(response =>response.json())
        .then(data =>{
            if (data.error){
                M.toast({html : data.error.message, classes : "rounded red-button"})
            }else{
                this.postData.imageUrl = data.url
                if (this.postData.option) {
                  console.log("hello")
                  fetch("http://localhost:8080/uploadPost",{
                      method : "post",
                      body : JSON.stringify({
                          description : this.postData.description,
                          imageUrl : this.postData.imageUrl,
                          option : "public",
                          email: this.userData.email,
                          username: this.userData.username
                      })
                  }).then(response =>response.json())
                  .then(data =>{
                      if (data){
                            M.toast({html : "Posted!!!", classes : "rounded green-button"})
                            this._router.navigate(["/home"])
                      }else{
                            M.toast({html : "Something Went Wrong", classes : "rounded red-button"})
                      }
                  })
              }else{
                  fetch("http://localhost:8080/uploadPost",{
                          body : JSON.stringify({
                              description : this.postData.description,
                              imageUrl : this.postData.imageUrl,
                              option : "private",
                              email: this.userData.email,
                              username: this.userData.username
                          })
                      }).then(response =>response.json())
                      .then(data =>{
                          if (data){
                              M.toast({html : "Posted!!!", classes : "rounded green-button"})
                              this._router.navigate(["/home"])
                          }else{
                              M.toast({html : "Something went wrong", classes : "rounded red-button"})
                          }
                      })
                    }
              }
            
        })
    
      }
}
