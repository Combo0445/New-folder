import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class DataService {
  private apiUrl = 'http://localhost:3000'; 

  constructor(private http: HttpClient) { }

  // 1. ดึงข้อมูล Posts ทั้งหมด
  getPosts(): Observable<any> {
    return this.http.get(`${this.apiUrl}/posts`);
  }

  // 2. ส่งคอมเมนต์ใหม่ไปยัง Backend (เพิ่มตัวนี้!)
  addComment(commentData: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/comments`, commentData);
  }

  // 3. สร้างโพสต์ใหม่ (ถ้าในอนาคตอยากทำปุ่มเพิ่มรูปสัตว์เลี้ยง)
  addPost(postData: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/posts`, postData);
  }

  getUsers(): Observable<any> {
    return this.http.get(`${this.apiUrl}/users`);
  }
}