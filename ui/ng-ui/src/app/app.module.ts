import { APP_BASE_HREF } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { NgModule, APP_INITIALIZER } from '@angular/core';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { environment } from '@env';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { TranslateModule } from '@ngx-translate/core';
import { HttpClientInMemoryWebApiModule } from 'angular-in-memory-web-api';
import json from 'highlight.js/lib/languages/json';
import scss from 'highlight.js/lib/languages/scss';
import typescript from 'highlight.js/lib/languages/typescript';
import xml from 'highlight.js/lib/languages/xml';
import { InlineSVGModule } from 'ng-inline-svg';
import { ClipboardModule } from 'ngx-clipboard';

import { HighlightModule, HIGHLIGHT_OPTIONS } from 'ngx-highlightjs';

import { FakeAPIService } from './_helpers/fake/fake-api.service';
import { SplashScreenModule } from './_metronic/partials/layout/splash-screen/splash-screen.module';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { GraphQLModule } from './graphql.module';
import { AuthService } from './modules/auth/_services/auth.service';
// Highlight JS

function appInitializer(authService: AuthService) {
  return () => {
    return new Promise((resolve) => {
      authService.getCurrentUserFromContext().subscribe({
        next: (user) => {
          // Successfully loaded user, continue initialization
          resolve(true);
        },
        error: () => {
          // Don't redirect on initial load, just allow app to initialize
          // This prevents redirect loops in Angular 16
          console.warn('Auth initialization failed, continuing without auth');
          resolve(true);
        },
        complete: () => {} // Completion is handled in next/error callbacks
      });
    });
  };
}

/**
 * Import specific languages to avoid importing everything
 * The following will lazy load highlight.js core script (~9.6KB) + the selected languages bundle (each lang. ~1kb)
 */
export function getHighlightLanguages() {
  return [
    { name: 'typescript', func: typescript },
    { name: 'scss', func: scss },
    { name: 'xml', func: xml },
    { name: 'json', func: json }
  ];
}

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    SplashScreenModule,
    TranslateModule.forRoot(),
    HttpClientModule,
    HighlightModule,
    ClipboardModule,
    environment.isMockEnabled
      ? HttpClientInMemoryWebApiModule.forRoot(FakeAPIService, {
          passThruUnknownUrl: true,
          dataEncapsulation: false
        })
      : [],
    AppRoutingModule,
    InlineSVGModule.forRoot(),
    NgbModule,
    GraphQLModule,
    MatSnackBarModule
  ],
  exports: [],
  providers: [
    {
      provide: APP_INITIALIZER,
      useFactory: appInitializer,
      multi: true,
      deps: [AuthService]
    },
    {
      provide: HIGHLIGHT_OPTIONS,
      useValue: {
        coreLibraryLoader: () => import('highlight.js/lib/core'),
        languages: getHighlightLanguages
      }
    },
    {
      provide: APP_BASE_HREF,
      useValue: '/'
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule {}
