# Backend Case Project
<img width="595" alt="Screenshot 2024-07-14 at 09 56 27" src="https://github.com/user-attachments/assets/53d952d6-7589-44a7-b7ac-9225c831a73c">
<img width="463" alt="Screenshot 2024-07-14 at 09 56 56" src="https://github.com/user-attachments/assets/4a0d8673-502c-4273-9828-c3cca4e0f6c8">


> Projenin database tarafındaki fotograflari

---

### Table of Contents
Bölümlerinizin başlıkları, referans almak için kullanılacaktır.
- [Description](#description)
- [How To Use](#how-to-use)

---

## Description
Projeye baslamadan once, hangi golang structure yapisini kullanmam gerektigine karar verdim.
Springboot projelerinden alisik oldugum icin bu yapiyi kullandim.

<img width="256" alt="Screenshot 2024-07-14 at 10 04 08" src="https://github.com/user-attachments/assets/12c5fec7-ac15-4752-bf53-22663f1bc01f">


#### Technologies

- Golang
- GORM Library
  
  ( ORM araci olarak GORM kutuphanesini kullandim. Olusturdugum modellerin veritabani islemlerini basitlestirdi. Yazmam gereken SQL sorgu yazimini ortadan kaldirdi.
  SQL tarafinda yazmam gereken sorgulari size mail yoluyla iletecegim.)
- Gorilla Mux
  
  ( Gorilla Mux, Go'da web uygulamaları oluşturmaya yönelik güçlü bir HTTP yönlendirici ve URL eşleştiricidir. )
  
- MySQL

  (Relational Database araclarindan en yaygin veri tabanı yönetim sistemidir. ) 

[Back To The Top](#read-me-template)

---

## How To Use

#### Installation
- Projeyi bilgisayarınıza indirdikten sonra, gerekli kütüphaneleri ve eklentileri indirmek için terminale aşağıdaki komutları yazınız.

  ```html
    go mod init backend_case
    go get -u github.gorilla/mux
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/mysql    

- Database bağlantısı için 'backend_case' adında bir şema oluşturmanız yeterli oluyor.
Sonrasında config/app.go dosyasına giderek kullanıcı adı ve parola yazdıktan sonra database bağlantısı kuruyorsunuz.

- Projeyi aşağıdaki komutu kullanarak çalıştırabilirsiniz. localhost:8000 serverında ayağa kalkıyor.

  ```html
   go run cmd/main.go



## API Reference

#### Create Team

```http
  POST http://localhost:8000/team
```
<img width="718" alt="Screenshot 2024-07-14 at 11 25 53" src="https://github.com/user-attachments/assets/5be3e51a-c4ca-45f6-a308-249e35a735b9">

#### Get All Teams

```http
  GET http://localhost:8000/teams
```

#### Play All League

```http
  POST http://localhost:8000/play-league
```

#### Get Matches By WeekId

```http
  GET http://localhost:8000/matches/{weekId}
```

#### Get All Matches 

```http
  GET http://localhost:8000/matches
```

## Mesaj
- Aldığım "invalid memory address or nil pointer dereference golang sql" hatasından dolayı, models package'inden league.go modeli ve bir kaç endpointi silmek zorunda kaldım.



[Back To The Top](#read-me-template)

---


