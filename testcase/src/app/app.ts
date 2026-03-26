import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { DataService } from './data';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class AppComponent implements OnInit {
  posts: any[] = [];
  users: any[] = []; 
  newComments: { [key: number]: string } = {};

  constructor(private dataService: DataService) {}

  ngOnInit() {
    this.loadUsers(); 
    this.loadPosts();
  }

  loadUsers() {
    this.dataService.getUsers().subscribe({
      next: (res: any) => {
        this.users = res;
        console.log('Users loaded:', this.users);
      },
      error: (err) => console.error(err)
    });
  }

  loadPosts() {
    this.dataService.getPosts().subscribe({
      next: (res: any) => {
        this.posts = res;
      },
      error: (err) => console.error(err)
    });
  }

  // ฟังก์ชันหาชื่อ Name โดยใช้ user_id เป็นตัวเชื่อม
  getUserName(idFromPost: number): string {
    if (!this.users || this.users.length === 0) return 'Loading...';
    
    // ค้นหาในตาราง users โดยเทียบ user_id ให้ตรงกัน
    const user = this.users.find(u => u.user_id === idFromPost);
    
    // ถ้าเจอให้คืนค่าคอลัมน์ Name ถ้าไม่เจอให้โชว์ User ID เดิม
    return user ? user.Name : 'User ' + idFromPost; 
  }

  sendComment(postId: number) {
    const message = this.newComments[postId];
    if (!message || message.trim() === '') return;

    const payload = {
      post_id: postId,
      user_id: 1, // ID ของคนที่กำลังพิมพ์คอมเมนต์ (สมมติว่าเป็นเรา)
      message: message
    };

    this.dataService.addComment(payload).subscribe({
      next: () => {
        this.newComments[postId] = '';
        this.loadPosts();
      },
      error: (err) => console.error(err)
    });
  }
}