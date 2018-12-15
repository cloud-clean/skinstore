PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- 表：auth_code
CREATE TABLE auth_code (account VARCHAR (20) PRIMARY KEY, code VARCHAR (50), expire DATETIME);
INSERT INTO auth_code (account, code, expire) VALUES ('test', 'aaaa', '2018-11-16 12:00:00');
INSERT INTO auth_code (account, code, expire) VALUES ('tt', 'bbbbb', '2018-11-16 18:30:00');

-- 表：auth_token
CREATE TABLE auth_token (account VARCHAR (50) PRIMARY KEY, access_token VARCHAR (150), flash_token VARCHAR (150), expire DATETIME);

-- 表：lot_status
CREATE TABLE lot_status (pos VARCHAR (20) PRIMARY KEY, status VARCHAR (5));
INSERT INTO lot_status (pos, status) VALUES ('lot3', '1');
INSERT INTO lot_status (pos, status) VALUES ('lot1', '0');

-- 表：lot_user
CREATE TABLE lot_user (account VARCHAR (50) PRIMARY KEY, password VARCHAR (50), "group" VARCHAR (12));
INSERT INTO lot_user (account, password, "group") VALUES ('cloud', 'cloud', '1');
INSERT INTO lot_user (account, password, "group") VALUES ('test', 'teset', '');

-- 表：project
CREATE TABLE project (id INTEGER PRIMARY KEY AUTOINCREMENT,name VARCHAR (50) NOT NULL,description VARCHAR (500),original_price INTEGER,type VARCHAR (5) NOT NULL,img_url VARCHAR (200) NOT NULL,cur_price INTEGER NOT NULL,status CHAR (1) DEFAULT (0));
INSERT INTO project (id, name, description, original_price, type, img_url, cur_price, status) VALUES (1, 'test01', 'is for test', 1500, 'f', 'http://img', 1000, '1');
INSERT INTO project (id, name, description, original_price, type, img_url, cur_price, status) VALUES (2, 'test02', 'test ofr test', 2000, 'c', 'img', 1500, '1');
INSERT INTO project (id, name, description, original_price, type, img_url, cur_price, status) VALUES (3, 'two', 'this a good product', 300, 'skin', 'http://www.img.com', 200, '1');

-- 表：reservation
CREATE TABLE reservation (id INTEGER PRIMARY KEY AUTOINCREMENT, user_id INTEGER, name VARCHAR (50) NOT NULL, re_time DATETIME NOT NULL, status CHAR (1) DEFAULT (0), mobile VARCHAR (12) NOT NULL, create_tm DATETIME DEFAULT (datetime('now', 'localtime')), project_id INTEGER NOT NULL);
INSERT INTO reservation (id, user_id, name, re_time, status, mobile, create_tm, project_id) VALUES (1, 22, 'cloud', '2018-05-11 13:41:14+08:00', '0', '13066852501', '2018-05-12 12:12:15', 2);
INSERT INTO reservation (id, user_id, name, re_time, status, mobile, create_tm, project_id) VALUES (2, 22, 'cloud', '2018-05-12 12:25:10+08:00', '0', '13066852501', '2018-05-12 12:25:36', 3);
INSERT INTO reservation (id, user_id, name, re_time, status, mobile, create_tm, project_id) VALUES (3, 23, 'cloud', '2018-05-13 12:25:10+08:00', '0', '13066852501', '2018-05-12 15:58:29', 5);

-- 表：user
CREATE TABLE user (user_id INTEGER PRIMARY KEY ON CONFLICT ROLLBACK AUTOINCREMENT, open_id VARCHAR (32) UNIQUE NOT NULL, nick_name VARCHAR (50), mobile VARCHAR (12), create_tm DATETIME DEFAULT ((datetime('now', 'localtime'))) COLLATE RTRIM, img_url VARCHAR (200));
INSERT INTO user (user_id, open_id, nick_name, mobile, create_tm, img_url) VALUES (1, '', '', '', '2018-05-01 08:56:39', '');
INSERT INTO user (user_id, open_id, nick_name, mobile, create_tm, img_url) VALUES (2, 'omm-iv0HbOZliIKyiEg9tmKTMwk8', '清灵', '', '2018-05-01 09:15:46', 'http://thirdwx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTLSrpu0ZJm6p2t6pKG1mvWOJ84cYvRDX4sMEElFv5KHickH0BibdS1UOgnVfR6MraJgqM55JFSAXp6w/132');

-- 表：wx
CREATE TABLE wx ("key" VARCHAR (10), token VARCHAR (100), expire DATETIME);
INSERT INTO wx ("key", token, expire) VALUES ('ACCTOKEN', 'AA', '2018-05-13');
INSERT INTO wx ("key", token, expire) VALUES ('JSTOKEN', 'VVV', NULL);

COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
