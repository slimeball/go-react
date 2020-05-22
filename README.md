[front-end 前端地址](https://github.com/slimeball/go-react/tree/react-frontend)  
[back-end 後端地址](https://github.com/slimeball/go-react/tree/go-backend)    


> skills
backend：go, gorm, gin, mysql
frontend: react, redux, weui 

-------------------------
Go Api support 

Sign up
Sign in
cookie session sign in authenticate

add book
book list
update book
delete book

book detail

add cart
delete cart
clear cart 

create order 

after 389 episodes basics Go learning, today will begin a Full-stack web project with Go, React, Mysql, Docker. Wish be good.

Day1. very hard to figure. 1, run Go with gin framework connect Mysql in docker, react install all dependencies and make signin page. 2, react signin page, send request receive response from Go Api. 3, Go read mysql find user exists or not give response.

Day2. 1, Go SignUp Api. 2, react SignUp page. request validate and response data.Develop process: main.go initialize route gin gorm etc, route, controller define func pass to route, model define struct same as table name, operate data.

D3. Go add new book Api.react packge ajax func with axios,ajax reqst put in independent file, book list page, add book page.Go use gorm operate database should use Table func give specify table name,if not it will auto use struct name+s be table name to operate



> 應用技術棧
後端：go, gorm, gin, mysql
前端: react, redux, weui 

Go接口支持

註冊
登錄
cookie session登錄控制

添加圖書
圖書列表
更新修改圖書
刪除圖書

圖書詳情

添加購物車
刪除購物車
清空購物車

創建訂單

在學習了389集go基礎視頻後，開始一個基於go，react，mysql的全棧模擬項目

第一天
初始化go環境，使用gin框架，gorm，在docker運行mysql並使用workbench連接。react初始化，引入axios，redux。
登錄功能，查詢是否存在用戶

第二天
go提供註冊api，react註冊頁面，請求並驗證數據。
整理了go開發流程，main.go入口文件初始化使用的依賴，路由，gorm等，路由定義後引入controller的函數作為回調，controller處理過程中引入model操作數據庫。

第三天
go添加書籍接口，react添加axio異步請求封裝包，異步函數獨立包，書籍列表頁，添加書籍頁。
go藉助gorm操作數據庫時，需要指定表名，不然就會在結構體名添加s以作為表名操作。

第四天
補充go mod, gin路由中間件，登錄cookie認證知識

第五天
go的接口，書列表，添加書，刪除書。
react添加書，刪除書頁面

第六天
調整go項目結構，書籍更新接口，react更新添加頁面合併

第七天
go分頁，默認頁碼為1，顯示條數為5，總頁碼通過數據庫計算。
前端請求接口時，不足10條則全部顯示，超過5條則顯示頁碼及總頁碼，
前端可以傳希望顯示的條數，go通過前端傳入的條數計算總頁碼，不傳則默認5條
dao層 data access object

第八天
登錄session記錄




---------------------------------
* 信息記錄

服務端agent
後端服務組件開發

技術棧
kafka  數據解耦
zookeeper
ES和Kibana使用
ercd

> 日志系统
问题排查，数据分析

- elk方案
问题：运维成本高，每增加一个日志收藏，都要手动配置，监控缺失，无法准确获取logstash状态，无法做定制化开发及维护

- 
log agent 日誌收集客戶端，收集服務器的日誌
lafka 高吞吐量分佈式隊列
ES - elasticsearch 開源搜索引擎
hadoop 分佈式計算框架，大量數據進行分部署，小量數據用mysql足夠
storm 實時計算框架，計算數據流

