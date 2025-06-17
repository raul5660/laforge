import { ElementRef, Injectable } from '@angular/core';
import { animate, AnimationBuilder, style } from '@angular/animations';

@Injectable({
  providedIn: 'root'
})
export class SplashScreenService {
  // Private properties
  private el: ElementRef;
  private stopped: boolean;
  private forceHideApplied: boolean = false;

  /**
   * Service constructor
   *
   * @param animationBuilder: AnimationBuilder
   */
  constructor(private animationBuilder: AnimationBuilder) {}

  /**
   * Init
   *
   * @param element: ElementRef
   */
  init(element: ElementRef) {
    this.el = element;
  }

  /**
   * Hide
   */
  hide() {
    console.log('SplashScreenService: hide() called');
    if (this.stopped || this.forceHideApplied) {
      console.log('SplashScreenService: already hidden');
      return;
    }
    
    if (!this.el) {
      console.error('SplashScreenService: no element reference');
      return;
    }

    try {
      const player = this.animationBuilder
        .build([style({ opacity: '1' }), animate(800, style({ opacity: '0' }))])
        .create(this.el.nativeElement);

      player.onDone(() => {
        console.log('SplashScreenService: animation complete');
        try {
          if (typeof this.el.nativeElement.remove === 'function') {
            this.el.nativeElement.remove();
            console.log('SplashScreenService: element removed');
          } else {
            this.el.nativeElement.style.display = 'none !important';
            console.log('SplashScreenService: element hidden');
          }
          this.stopped = true;
        } catch (err) {
          console.error('SplashScreenService: error hiding splash screen', err);
        }
      });

      setTimeout(() => {
        console.log('SplashScreenService: playing hide animation');
        player.play();
      }, 100);
    } catch (err) {
      console.error('SplashScreenService: critical error in hide()', err);
      // Force hide splash screen on error
      if (this.el && this.el.nativeElement) {
        this.el.nativeElement.style.display = 'none !important';
        this.stopped = true;
      }
    }
  }
  
  /**
   * Force Hide
   * Emergency method to forcibly hide the splash screen when other methods fail
   */
  forceHide() {
    console.log('SplashScreenService: forceHide called - emergency hiding of splash screen');
    this.forceHideApplied = true;
    
    try {
      if (this.el && this.el.nativeElement) {
        // Try removing with standard method first
        if (typeof this.el.nativeElement.remove === 'function') {
          this.el.nativeElement.remove();
          console.log('SplashScreenService: force removed element');
        }
        // Fallback to style-based hiding
        else {
          this.el.nativeElement.style.display = 'none !important';
          console.log('SplashScreenService: force hidden with style');
        }
        
        // Mark as stopped
        this.stopped = true;
        
        // Try to use DOM methods directly as a last resort
        const splashElements = document.querySelectorAll('app-splash-screen');
        if (splashElements && splashElements.length > 0) {
          console.log(`SplashScreenService: found ${splashElements.length} splash screen elements to force remove`);
          splashElements.forEach(element => {
            if (element.parentNode) {
              element.parentNode.removeChild(element);
            }
          });
        }
      } else {
        console.error('SplashScreenService: no element reference for force hide');
      }
    } catch (err) {
      console.error('SplashScreenService: error in forceHide', err);
    }
  }
}
