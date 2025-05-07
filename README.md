Todo App API
Bu, kullanıcıların yapılacaklar listelerini oluşturabileceği ve yönetebileceği bir RESTful API'dir. Uygulama, Gin framework'u kullanılarak geliştirilmiştir ve JWT (JSON Web Token) ile kimlik doğrulama sağlar.

Başlangıç
Bu adımlar, projeyi yerel ortamınızda çalıştırmanızı sağlayacaktır.

Gereksinimler
Go 1.18 veya üstü
Gin Framework
JWT kütüphanesi
Bağımlılıklar için Go modül yönetimi
Kurulum
Repoyu klonlayın:
git clone https://github.com/erdemyuksel4/todoapp.git
Gerekli bağımlılıkları yükleyin:

bash Kopyala Düzenle cd todoapp go mod tidy Ortam değişkenlerini ayarlayın (isteğe bağlı):

Eğer özel bir port kullanmak isterseniz, .env dosyasını oluşturup PORT değerini ayarlayabilirsiniz.

Sunucuyu başlatın:

bash Kopyala Düzenle go run main.go Sunucu varsayılan olarak localhost:8080 adresinde çalışacaktır.

API Kullanımı Sağlık Durumu Kontrolü API'nin çalışıp çalışmadığını kontrol etmek için aşağıdaki endpoint'i kullanabilirsiniz:

http Kopyala Düzenle GET /health Yanıt:

json Kopyala Düzenle { "status": "up" } Kullanıcı Girişi (Login) Kullanıcı adı ve şifre ile giriş yapabilir ve JWT token alabilirsiniz.

http Kopyala Düzenle POST /login Giriş yapmak için aşağıdaki JSON verisini gönderin:

json Kopyala Düzenle { "username": "your-username", "password": "your-password" } Yanıt:

json Kopyala Düzenle { "token": "your-jwt-token" } JWT token, daha sonraki API isteklerinde Authorization başlığı ile kullanılmalıdır.

Kullanıcı Bilgileri JWT token'ı ile doğrulama yaptıktan sonra, kullanıcı bilgilerinizi alabilirsiniz:

http Kopyala Düzenle GET /api/me Yanıt:

json Kopyala Düzenle { "userId": 1, "type": 1 } Todo Listeleri Tüm Todo Listelerini almak için:

http Kopyala Düzenle GET /api/todolists Todo Listesine yeni bir liste eklemek için:

http Kopyala Düzenle POST /api/addlist Gönderilecek JSON:

json Kopyala Düzenle { "title": "New Todo List" } Todo İşlemleri Tüm Todoları listelemek için:

http Kopyala Düzenle GET /api/todos ID'ye Göre Todo'yu almak için:

http Kopyala Düzenle GET /api/todobyid/:id Todo'yu Tamamlamak için:

http Kopyala Düzenle PATCH /api/complete/:id Todo Mesajını Değiştirmek için:

http Kopyala Düzenle PATCH /api/changemessage Gönderilecek JSON:

json Kopyala Düzenle { "todoId": 1, "message": "New Todo Message" } Todo Silmek için:

http Kopyala Düzenle DELETE /api/deletetodo Gönderilecek JSON:

json Kopyala Düzenle { "todoId": 1 } Liste Silmek Todo Listesini Silmek için:

http Kopyala Düzenle DELETE /api/deletelist Gönderilecek JSON:

json Kopyala Düzenle { "id": 1 } Kimlik Doğrulama Bu uygulama, JWT (JSON Web Token) kullanarak kimlik doğrulaması yapmaktadır. API isteklerinde Authorization başlığı ile geçerli bir token sağlamalısınız.

Örnek başlık:

makefile Kopyala Düzenle Authorization: Bearer your-jwt-token Kullanıcılar Projede tanımlı üç örnek kullanıcı bulunmaktadır:

Kullanıcı Adı: Erdem Şifre: 123 Tür: TodoUser ID: 1

Kullanıcı Adı: user2 Şifre: 321 Tür: TodoUser ID: 2

Kullanıcı Adı: admin Şifre: 123 Tür: Admin ID: 3

Bu kullanıcılar, kimlik doğrulama için kullanılabilir. Kullanıcı adı ve şifre bilgilerini doğru girdiğinizde, JWT token alabilirsiniz.
