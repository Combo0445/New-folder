import 'zone.js';
import { bootstrapApplication } from '@angular/platform-browser';
import { appConfig } from './app/app.config';
import { AppComponent } from './app/app'; // <--- เช็คว่าแก้เป็น ./app/app หรือยัง

bootstrapApplication(AppComponent, appConfig)
  .catch((err) => console.error(err));