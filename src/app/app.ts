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
  
  // ใช้เก็บข้อความที่พิมพ์ในช่อง Input แยกตาม ID ของแต่ละโพสต์
  newComments: { [key: number]: string } = {}; 

  constructor(private dataService: DataService) {}

  ngOnInit() {
    this.loadPosts();
  }

  loadPosts() {
    this.dataService.getPosts().subscribe((res: any) => {
      this.posts = res;
      console.log("Posts loaded:", this.posts);
    });
  }

  // ฟังก์ชันส่งคอมเมนต์เมื่อกดปุ่ม
  sendComment(postId: number) {
    const message = this.newComments[postId];
    
    // ถ้าไม่ได้พิมพ์อะไรเลย หรือมีแต่ช่องว่าง จะไม่ส่งข้อมูล
    if (!message || message.trim() === '') return;

    const commentData = {
      post_id: postId,  // ส่ง ID ของโพสต์นั้นๆ ไป (เช่น โพสต์ที่ 5)
      user_id: 1,       // กำหนดให้คนคอมเมนต์เป็น User ID 1 เสมอ
      message: message
    };

    this.dataService.postComment(commentData).subscribe({
      next: (res) => {
        console.log("Comment sent successfully:", res);
        this.newComments[postId] = ''; // ล้างช่องพิมพ์ในหน้าจอ
        this.loadPosts(); // รีโหลดข้อมูลใหม่เพื่อแสดงคอมเมนต์ล่าสุด
      },
      error: (err) => {
        console.error("Error sending comment:", err);
        alert("ไม่สามารถส่งคอมเมนต์ได้ กรุณาเช็คว่ามี User ID 1 ในระบบหรือยัง");
      }
    });
  }
}