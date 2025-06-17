import { Injectable } from '@angular/core';
import {
  Router,
  CanActivate,
  ActivatedRouteSnapshot,
  RouterStateSnapshot,
} from '@angular/router';
import { AuthService } from './auth.service';

@Injectable({ providedIn: 'root' })
export class AuthGuard implements CanActivate {
  constructor(private authService: AuthService, private router: Router) {}

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    console.log('AuthGuard: canActivate check for route', state.url);
    const currentUser = this.authService.currentUserValue;
    
    if (currentUser) {
      // logged in so return true
      console.log('AuthGuard: User authenticated, allowing navigation');
      return true;
    }

    console.log('AuthGuard: No authenticated user found');
    
    // If we're already on a login-related page, don't trigger another logout
    // This prevents potential redirect loops
    if (state.url.includes('/auth/login')) {
      console.log('AuthGuard: Already on login page, not redirecting');
      return true;
    }
    
    // not logged in so redirect to login page with the return url
    console.log('AuthGuard: Redirecting to login page');
    this.authService.logout().subscribe(() => {
      console.log('AuthGuard: Logout successful, navigating to login page');
      this.router.navigate(['auth', 'login']); 
    });
    return false;
  }
}
