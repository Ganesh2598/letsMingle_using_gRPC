import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from "./login/login.component";
import { RegisterComponent } from "./register/register.component";
import { ProfileComponent } from "./profile/profile.component";
import { AddpostComponent } from "./addpost/addpost.component";
import { AddfriendsComponent } from "./addfriends/addfriends.component";
import { UserProfileComponent } from "./user-profile/user-profile.component";
import { MyfriendsComponent } from "./myfriends/myfriends.component";
import { HomeComponent } from "./home/home.component";
import { AuthGuard } from "./auth.guard";


const routes: Routes = [
  {
    path : "",
    component : LoginComponent
  },
  {
    path : "login",
    component : LoginComponent
  },
  {
    path : "register",
    component : RegisterComponent
  },
  {
    path : "profile",
    component : ProfileComponent,
    canActivate : [AuthGuard]
    
  },
  {
    path : "userprofile/:mail",
    component : UserProfileComponent,
    canActivate : [AuthGuard]
  },
  {
    path : "myfriends",
    component : MyfriendsComponent,
    canActivate : [AuthGuard]
  },
  {
    path : "addpost",
    component : AddpostComponent,
    canActivate : [AuthGuard]
  },
  {
    path : "addfriends",
    component : AddfriendsComponent,
    canActivate : [AuthGuard]
  },
  {
    path : "home",
    component : HomeComponent,
    canActivate : [AuthGuard]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
