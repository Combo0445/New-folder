import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DataService {
  // URL ของ Go Backend
  private apiUrl = 'http://localhost:3000'; 

  constructor(private http: HttpClient) { }

  // 1. ฟังก์ชันดึงข้อมูล Posts ทั้งหมด (รวมคอมเมนต์ที่ Preload มาจาก Go)
  getPosts(): Observable<any> {
    return this.http.get(`${this.apiUrl}/posts`);
  }

  // 2. ฟังก์ชันสำหรับ "ส่งคอมเมนต์ใหม่" ไปที่ Go Backend (เพิ่มอันนี้เข้าไปครับ)
  postComment(commentData: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/comments`, commentData);
  }

  // 3. แถม: ถ้าอยากดึง User ทั้งหมด
  getUsers(): Observable<any> {
    return this.http.get(`${this.apiUrl}/users`);
  }
}