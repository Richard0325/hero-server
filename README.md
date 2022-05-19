# Hahow專案

## 我們該如何跑起這個 server?
### 安裝docker並開啟docker engine
參考[docker官網](https://docs.docker.com/get-docker/)
### 啟動server
```
docker-compose build
```
接下來
```
docker-compose up -d
```
執行完後便在container中執行，預設port為8080
```
/yourserver:8080/heroes
/yourserver:8080/heroes/{heroId}
```
## 專案的架構，API server 的架構邏輯?
### 資料夾及檔案結構如下：
```
|__cmd              # 主程式
|  |__server              # 最終執行的主程式
|     |__main.go
|  |__test                # 開發過程中測試區
|     |__main.go
|__controler        # 為各個route的handler
|  |__controller_test.go  # 單元測試
|  |__controller.go
|__data             # 通用的資料結構皆在此定義
|  |__data.go
|  |__error.go
|__model            # service需要取得資料的dao
|  |__dao.go
|  |__hahow_test.go       # 單元測試
|  |__hahow.go            # 呼叫Hahow API的函式
|  |__mock.go             # 呼叫hahow API的模擬版（為測試用）
|__service          # handler所用到的服務(service)
|  |__service_test.go     # 單元測試
|  |__service.go        
|__tools            # 通用的函式庫
|  |__tools.go
|__APITEST.md       # API測試紀錄
|__docker-compose.yml   
|__Dockerfile           
|__go.mod           
|__go.sum
|__README.md
```
### API邏輯如下：
```
/GetHeroes
    |-> 有Name及Password
    |   '-> (A)檢查Auth
    |       |-> 檢查成功
    |       |   '-> (C)取得heroes及各自的profile
    |       |       |-> 成功
    |       |       |   '-> 回傳 200, heroes及profiles
    |       |       |
    |       |       '-> 失敗
    |       |           '-> 回傳 500
    |       '-> 失敗或error
    |           '-> (B)取得heroes
    |               |-> 成功
    |               |   '-> 回傳 200, heroes
    |               |
    |               '-> 失敗
    |                   '-> 回傳 500
    '-> 缺少Name或Password
        '-> (B)取得heroes
            |-> 成功
            |   '-> 回傳 200
            |
            '-> 失敗
                '-> 回傳 500

/GetHeroes/:heroId
    |-> 有Name及Password
    |   '-> (A)檢查Auth
    |       |-> 檢查成功
    |       |   '-> (E)取得指定hero及其profile
    |       |       |-> heroId正確
    |       |       |   |-> 成功
    |       |       |   |   '-> 回傳 200, hero及其profile
    |       |       |   |
    |       |       |   '-> 失敗
    |       |       |       '-> 回傳 500
    |       |       |
    |       |       '-> heroId不正確
    |       |           '-> 回傳 404
    |       '-> 失敗或error
    |           '-> (D)取得指定hero
    |               |-> heroId正確
    |               |   |-> 成功
    |               |   |   '-> 回傳 200, hero及其profile
    |               |   |
    |               |   '-> 失敗
    |               |       '-> 回傳 500
    |               |
    |               '-> heroId不正確
    |                   '-> 回傳 404
    '-> 缺少Name或Password
        '-> (D)取得指定hero
            |-> heroId正確
            |   |-> 成功
            |   |   '-> 回傳 200, hero及其profile
            |   |
            |   '-> 失敗
            |       '-> 回傳 500
            |
            '-> heroId不正確
                '-> 回傳 404

(A)檢查Auth：呼叫Hahow Authenticate(*)，回傳是否Authorized
(B)取得heroes：呼叫Hahow List Heroes(*)，回傳heroes
(C)取得heroes及各自的profile：執行(B)，歷遍每個hero呼叫Hahow Profile of Hero取得profile並回傳
(D)取得指定hero：呼叫Hahow Single Hero(*), 回傳指定hero或告知未找到。
(E)取得指定hero及其profile：執行(D)，若找到指定hero便呼叫Hahow Profile of Hero填上profile並回傳
＊：遇到backend error就重複呼叫直到取得結果，若大於十次未果，則回傳錯誤
```
## 你對於所有使用到的第三方 library 的理解，以及他們的功能簡介?
- [gin](https://github.com/gin-gonic/gin)\
為 golang 的 API 框架，可以很方便的架起 API server，使對應的 url 導入到對應的 handler，並且讓對應的 handler 可以很簡便的取得 http request 中 query string、header 或 body 以及回傳 response。
- [logrus](https://github.com/sirupsen/logrus)\
為 golang 原生 log 系統的強化版，我主要用其可以將 log 區分等級，方便在開發時在適當的時機過濾並且取得想要的 log 以方便 debug。
## 你在程式碼中寫註解的原則，遇到什麼狀況會寫註解?
原則上我希望盡量透過命名及結構來增加可讀性，註解越少越好，以下情形是比較常會寫註解的情況：
1. 有點長的判斷式，不好看出目的
2. 程式不夠直觀看出意圖，通常是包了很多層的函式
3. 用了非常多行在完成一個小目標，說明目標
4. 假如函式的命名不夠給足意義，陳述函式的用途

## 在這份專案中你遇到的困難、問題，以及解決的方法?
<br>Q: 需求上並未定義萬一使用者Name和Password錯誤要如何應對？要回傳除profile以外的資料嗎？要提示嗎？<br/>
<br>A: 詢問專案內成員後，成員希望我去思考各種作法的利弊，最後決定只要 Name 或 Password 不正確，皆會給一模一樣的回應，不會暗示任何關於key或value有關的訊息，假設profile是給特定的使用者才可以得到的資料，那特定使用者應該會知道要正確輸入name及password才會去得到隱藏資料，而打錯也不會有任何提示，因為在這安全性不嚴謹的情況下太清楚地透露都會是危險的，所以我覺得只有不給提示是一個比較好的做法。<br/>
<br>Q: 如何處理hahow 提供的 API 在回傳 200 時，不定期會出現 error 導致無法取得原來的資料？ <br/>
<br>A: 應用上每當我需要這個API並呼叫它時，我發現其並不會一直壞掉，因此只要判斷出是 error 時會再次呼叫直到拿到想要的資料為止，但萬一一直是 error 超過10次便放棄, 回傳 server error。<br/>
<br>Q: 呈上個問題，如何解決API 這個不穩定的行為亦造成再上一層用到這個API的程式 (service) 測試困難?<br/>
<br>A: 多寫了一個 mock 的物件來模擬hahow API會出現的回傳，模擬當其一直出錯和模擬其一直保持正常來測試上層（service），並將 dao 改成了 interface，使其可以彈性的調整為呼叫 Hahow server 或 mock 以便於掌控行為已測試各種情況的應對行為。<br/>

## 後記，有什麼可能可以做更好?
1. git的使用上應該可以讓每個commit完成的事項更單一，避免跨足太多feature
2. API在過程中先取得基本資料再確認是否Authorized可能更有效率
3. 可能可以從先思考如何測試再完備功能會讓unit test好寫點
4. 為了應對 Hahow API error 導致 API 反應的時間比較慢，可以設定 timeout 以避開等太久的問題