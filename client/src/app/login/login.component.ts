import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import M from "materialize-css";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(private _router: Router) { }

  loginData = {
    email : "",
    password : ""
  }

  ngOnInit(): void {
  }

  onSubmit = () => {
    fetch("http://localhost:8080/signin",{
            method : "post",
            headers : { "Content-Type" : "application/json"},
            body : JSON.stringify({
                email : this.loginData.email,
                password : this.loginData.password
            })
        }).then(response => response.json())
        .then(data =>{
            if(Object.keys(data).length > 1){
                const user = {
                  username: data.username,
                  email: data.email,
                  imageUrl: data.imageUrl
                }
                localStorage.setItem("token",data.token)
                localStorage.setItem("user",JSON.stringify(user))
                M.toast({html : `Successfully Logged in`, classes : "rounded green-button"})
                this._router.navigate(["/home"])
            }else{
               M.toast({html : "Email or Password is Wrong", classes : "rounded red-button"})
            }
        })
  }

}
