# Dự án triển khai hệ thống quản lý địa chỉ IP, cung cấp các API thêm, xóa, truy xuất và cấp phát địa chỉ IP tự động tránh xung đột IP, kiểm tra IPv4 và v6.


- Ngôn ngữ: Golang.
- Sử dụng Gin framework triển khai REST API. 
- Dùng MySQL để lưu trữ dữ liệu.

CÁC CHỨC NĂNG CHÍNH: 
+ Thêm địa chỉ IP
+ Hiển thị danh sách IP
+ Xóa IP
+ Xác thực IPv4 và IPv6
(*) Mức advance: Cấp phát động địa chỉ IP và tối ưu hiệu suất tìm địa chỉ IP

Phiên bản hệ thống: Golang v1.18+, Docker và Docker Compose, MySQL 8.0+
(*)LƯU Ý: + Các phiên bản được để cập phải trùng hoặc tương thích, tránh xung đột khi chạy container.
	  + Trên máy khách chạy Docker Container phải để trống cổng 8080 cho Go app và 3306 cho MySQL. Nếu cả 2 cổng đã chạy dịch vụ khác anh có thể khắc phục bằng lệnh sau: 
		-> docker run -d -p 9090:8080 --name go_app ip_management-app (app trong container vẫn chạy 8080 nhưng bên ngoài chạy 9090)
		-> docker run -d -p 3307:3306 --name mysql_container mysql:8.0 (tương tự như lệnh trên)

Hướng dẫn sử dụng:

B1: Giải nén file ip_management.zip

B2: Sử dụng Powershell (Windows), di chuyển vào vị trí thư mục ip_management

B3: chạy lệnh: docker-compose up -d
- Tạo container MySQL và API
- Tạo cơ sở dữ liệu, table, index,...

B4: kiểm tra container bằng câu lệnh: docker ps
- Nếu cột IMAGE có 2 mục go_app và mysql:x.x và cả 2 có STATUS là Up (Healthy) là API đã khởi chạy

B5: kiểm tra API:
Với WINDOWS
- Thêm địa chỉ IP: Invoke-RestMethod -Uri "http://localhost:8080/ip" -Method POST -Body '{"ip_address":"192.168.1.10"}' -ContentType "application/json"
- Lấy danh sách IP: Invoke-RestMethod -Uri "http://localhost:8080/ips" -Method GET
- Xóa địa chỉ IP: Invoke-RestMethod -Uri "http://localhost:8080/ip/192.168.1.10" -Method DELETE
- Cấp phát động địa chỉ IP: Invoke-RestMethod -Uri "http://localhost:8080/ip/allocate" -Method GET
- Validate địa chỉ IP:
+ IPv4: Invoke-RestMethod -Uri "http://localhost:8080/validate-ip" -Method POST -Body '{"ip_address":"192.168.1.10"}' -ContentType "application/json"
+ IPv6: Invoke-RestMethod -Uri "http://localhost:8080/validate-ip" -Method POST -Body '{"ip_address":"2001:db8::ff00:42:8329"}' -ContentType "application/json"
Với LINUX
- Thêm địa chỉ IP: curl -X POST "http://localhost:8080/ip" -H "Content-Type: application/json" -d '{"ip_address":"192.168.1.10"}'
- Lấy danh sách IP: curl -X GET "http://localhost:8080/ips"
- Xóa địa chỉ IP: curl -X DELETE "http://localhost:8080/ip/192.168.1.10"
- Cấp phát động địa chỉ IP: curl -X GET "http://localhost:8080/ip/allocate"
- Validate địa chỉ IP:
+ IPv4: curl -X POST "http://localhost:8080/validate-ip" -H "Content-Type: application/json" -d '{"ip_address":"192.168.1.10"}'
+ IPv6: curl -X POST "http://localhost:8080/validate-ip" -H "Content-Type: application/json" -d '{"ip_address":"2001:db8::ff00:42:8329"}'
(*)LƯU Ý: Trong các câu lệnh Thêm, Xóa, Validate có các địa chỉ IP minh họa, anh có thể thay thế bằng địa chỉ IP mình muốn.

B6: Dừng và xóa container:
+ Dừng: docker-compose down
+ Xóa: docker-compose down -v


