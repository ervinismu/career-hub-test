# career-hub-test
1. 

- tools and Lang :  
1. Golang = karena golang adalah bahasa yang mumpuni, cepat dan bagus untuk dibuat sebagai service backend  
2. JWT(Json Web Token) = karena digunakkan sebagai pengaman untuk backend RestApi, agar data lebih aman, pengaksesan data backend dilakukkan menggunakkan JWT untuk authentication
3. VsCode = karena ini termasuk tools editor yang cepat dan ringkas, ada banyak plugin yang bisa digunakkan untuk membantu develop  
4. CI/CD (Travis.ci, Circle.ci, Gitlab.ci) = karena agar memudahkan kita untuk melakukkan deploy produk, agar automation oleh tools dan mengurangi kesalahan Human Error  
5. AWS(Amazon Web Service) = karena mudah dalam penggunaannya  
6. Gitlab = karena untuk memanajemen kode kita saat dikerjakan secara team agar lebih mudah

2.Untuk menanggani pengiriman data pada backend dari mobile di RestApi, kita menggunakkan JWT (Json Web Token) untuk setiap pengaksesan url Api yang menggandung data sensitif, setiap ada yang meng akses harus menyertakan Token agar kita dapat memastikan bahwa yang mengakses untuk mengedit/mengambil data adalah user/perangkat yang valid atau sudah terdaftar pada sistem kita.

3. How to run apps ?
- Install Go-lang
- go run main.go  

# Testing in Url
Testing your url with Postman or curl

| Name                | Url                   | Method   |
| --------------------|:---------------------:|:--------:|
| Register            | `<your_url>/register` |   **POST**   |
| Login               | `<your_url>/login`    |   **POST**   |
| Logout              | `<your_url>/logout`   |   **GET**    |
| Get All Notes        | `<your_url>/notes`     |   **GET**    |
| Create Notes         | `<your_url>/notes`     |   **POST**   |
| Delete Notes         | `<your_url>/notes`     |   **DELETE** |
| Update Notes         | `<your_url>/notes`     |   **PUT** |
