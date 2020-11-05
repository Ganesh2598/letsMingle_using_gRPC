import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from "@angular/forms";
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { NavbarComponent } from './navbar/navbar.component';
import { ProfileComponent } from './profile/profile.component';
import { AddpostComponent } from './addpost/addpost.component';
import { AddfriendsComponent } from './addfriends/addfriends.component';
import { HomeComponent } from './home/home.component';
import { AuthGuard } from "./auth.guard";
import { AuthService } from "./auth.service";
import { UserProfileComponent } from './user-profile/user-profile.component';
import { MyfriendsComponent } from './myfriends/myfriends.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegisterComponent,
    NavbarComponent,
    ProfileComponent,
    AddpostComponent,
    AddfriendsComponent,
    HomeComponent,
    UserProfileComponent,
    MyfriendsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule
  ],
  providers: [AuthGuard,AuthService],
  bootstrap: [AppComponent]
})
export class AppModule { }
