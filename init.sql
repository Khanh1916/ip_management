-- khoi tao co so du lieu
CREATE TABLE IF NOT EXISTS ips (
    id INT AUTO_INCREMENT PRIMARY KEY,
    ip_address VARCHAR(45) UNIQUE NOT NULL,
    allocated BOOLEAN DEFAULT FALSE
);

-- Tao chi muc toi uu hoa tim kiem IP chua duoc cap phat
CREATE INDEX idx_allocated ON ips(allocated);
CREATE INDEX idx_ip_address ON ips(ip_address);

