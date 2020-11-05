import { Component, OnInit } from '@angular/core';
import { Router } from "@angular/router";
import { AuthService } from "../auth.service";
import M from "materialize-css";

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  constructor(private _router : Router, public check:AuthService) { }
  ngOnInit(): void {
  }

  logout = () => {
    localStorage.clear()
    M.toast({html : "Successfully Logged out", classes : "rounded red-button"})
    this._router.navigate(["/login"])
  }

}
