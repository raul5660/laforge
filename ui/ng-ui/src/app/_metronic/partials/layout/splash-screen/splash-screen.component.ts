import { Component, OnInit, ViewChild, ElementRef, OnDestroy } from '@angular/core';
import { SplashScreenService } from './splash-screen.service';

@Component({
  selector: 'app-splash-screen',
  templateUrl: './splash-screen.component.html',
  styleUrls: ['./splash-screen.component.scss'],
})
export class SplashScreenComponent implements OnInit, OnDestroy {
  @ViewChild('splashScreen', { static: true }) splashScreen: ElementRef;

  private forceHideTimeout: ReturnType<typeof setTimeout>;
  
  constructor(
    private el: ElementRef,
    private splashScreenService: SplashScreenService
  ) {}

  ngOnInit(): void {
    console.log('SplashScreenComponent: Initializing');
    this.splashScreenService.init(this.splashScreen);
    
    // Force hide splash screen after 10 seconds as a fallback
    // This ensures the app becomes usable even if there are WebSocket or auth issues
    this.forceHideTimeout = setTimeout(() => {
      console.log('SplashScreenComponent: Force hiding splash screen after timeout');
      this.splashScreenService.forceHide();
      document.body.classList.add('page-loaded'); // Ensure body content is displayed
    }, 10000);
  }
  
  ngOnDestroy(): void {
    // Clean up timeout if component is destroyed
    if (this.forceHideTimeout) {
      clearTimeout(this.forceHideTimeout);
    }
  }
}
