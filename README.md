## Vatansoft Case
Bu proje Vatansoft adına stajyerlik pozisyonu başvurusu için yapılmıştır. 

# Öğrenci Ekleme

127.0.0.1:1323/createStudent adlı adrese post isteği olarak postman'da body kısmında raw alt seçeneğinden JSON formatında bu bilgileri gönderiyoruz : 
"email" : Öğrenci Email
"password" : Öğrenci Şifresi
"name": Öğrenci Adı
"surname" : Öğrenci Soyadı 

![](https://i.hizliresim.com/nz30mjd.png)

# Öğrenci Güncelleme

127.0.0.1:1323/updateStudent adlı adrese post isteği olarak postman'da body kısmında raw alt seçeneğinden JSON formatında bu bilgileri gönderiyoruz , sistem bu bilgilere göre otomatik olarak veritabanında öğrenciyi buluyor ve güncelliyor: 

"email" : Öğrenci Email
"password" : Öğrenci Şifresi
"name": Öğrenci Adı
"surname" : Öğrenci Soyadı 

![](https://i.hizliresim.com/7z1pj84.png)

# Öğrenci Silme 

127.0.0.1:1323/deleteStudent adlı adrese postman'da parametre kısmında id değişkenini giriyoruz , post isteği atıyoruz ve sistem bu bilgilere göre otomatik olarak veritabanında öğrenciyi buluyor ve siliyor : 

![](https://i.hizliresim.com/qpot4k4.png)

# ID Vererek Öğrenci Bulma

127.0.0.1:1323/getStudentById adlı adrese postman'da post isteği atıyoruz ve belirtilen ID değerindeki öğrenci geliyor : 

![](https://www.hizliresim.com/rre308t)

# Bütün Öğrencileri Bulma

127.0.0.1:1323/getStudents adlı adrese postman'da get isteği atıyoruz ve öğrenciler geliyor : 

![](https://i.hizliresim.com/5qdpokl.png)

# Plan Ekleme

127.0.0.1:1323/addPlan adlı adrese post isteği olarak postman'da body kısmında raw alt seçeneğinden JSON formatında bu bilgileri gönderiyoruz :

"lesson" : Ders veya yapılacak işin konusu
"lessonContent" : Ders veya yapılacak işin içeriği
"day" : Yapılacak gün bilgisidir, verilen bilgi bu şekilde olmalıdır "YYYY-MM-DD"
"startTime" : Başlangıç saati bilgisi , verilen bilgi bu şekilde olmalıdır "15.00" , "HH-MM"
"finishTime" : Bitiş saati bilgisi , verilen bilgi bu şekilde olmalıdır "15.00" , "HH-MM"
"state" : Otomatik olarak verilmektedir ve yapılacak olarak işaretlenmektedir

![](https://i.hizliresim.com/mk0u9aw.png)

# Plan Güncelleme 

127.0.0.1:1323/updatePlan adlı adrese post isteği olarak postman'da body kısmında raw alt seçeneğinden JSON formatında bu bilgileri gönderiyoruz :

"lesson" : Ders veya yapılacak işin konusu
"lessonContent" : Ders veya yapılacak işin içeriği
"day" : Yapılacak gün bilgisidir, verilen bilgi bu şekilde olmalıdır "YYYY-MM-DD"
"startTime" : Başlangıç saati bilgisi , verilen bilgi bu şekilde olmalıdır "15.00" , "HH-MM"
"finishTime" : Bitiş saati bilgisi , verilen bilgi bu şekilde olmalıdır "15.00" , "HH-MM"
"state" : Otomatik olarak verilmektedir ve yapılacak olarak işaretlenmektedir

![](https://i.hizliresim.com/5uz5dyj.png)

# Plan Silme

127.0.0.1:1323/deletePlan adlı adrese post isteği olarak postman'da parametre olarak id veriyoruz ve plan siliniyor :

![](https://i.hizliresim.com/4jiis14.png)

# Plan Durumunu Bitti Olarak Güncelleme

127.0.0.1:1323/finishPlan adlı adrese post isteği olarak postman'da parametre olarak id veriyoruz ve planın durumu bitti olarak güncelleniyor :

![](https://i.hizliresim.com/k6f7ds6.png)

# Plan Durumunu İptal Olarak Güncelleme

127.0.0.1:1323/cancelPlan adlı adrese post isteği olarak postman'da parametre olarak id veriyoruz ve planın durumu iptal olarak güncelleniyor :

![](https://i.hizliresim.com/6wax4m1.png)

# ID Vererek Plan Bulma

127.0.0.1:1323/getPlanById adlı adrese postman'da post isteği atıyoruz ve belirtilen ID değerindeki plan geliyor : 

![](https://www.hizliresim.com/q24rzhw)

# Bütün Planları Görüntülemek

127.0.0.1:1323/getPlans adlı adrese get isteği olarak postman'da istek atıyoruz ve bütün planları görüntülüyoruz :

![](https://i.hizliresim.com/ce2svdr.png)

# Haftalık Planları Görüntülemek

127.0.0.1:1323/getPlansWeekly adlı adrese get isteği olarak postman'da istek atıyoruz ve haftalık bütün planları görüntülüyoruz :

![](https://i.hizliresim.com/dd9v47a.png)

# Aylık Planları Görüntülemek

127.0.0.1:1323/getPlansMonthly adlı adrese get isteği olarak postman'da istek atıyoruz ve aylık bütün planları görüntülüyoruz :

![](https://i.hizliresim.com/c57s5fq.png)








