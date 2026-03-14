# Alman Hesabi - User Stories

> Tricount benzeri bir hesap bölüşme uygulaması.
> Bu dosya, AI-native geliştirme sürecinde "Source of Truth" görevi görür.
> MVP için hedef: **sıfır sürtünme**, **kayıt/giriş yok**, **link tabanlı gruplar**.

---

## MVP Epic: Frictionless Grup & Harcama

### Story 1: Anonim Grup Oluşturma (Sign up / Sign in YOK)

**User Story:** Bir kullanıcı olarak, uygulamaya girer girmez tek tıkla yeni bir harcama grubu oluşturmak istiyorum, böylece kayıt/giriş yapmadan harcamaları listelemeye başlayabileyim.

**Acceptance Criteria:**
- Landing'e giren kullanıcıdan email/şifre istenmez, direkt grup oluşturma ekranı gösterilir.
- "Yeni grup oluştur" butonuna basıldığında benzersiz bir grup ID ve gizli token üretilir.
- Kullanıcıdan sadece görünen bir **isim** istenebilir (opsiyonel; girmezse "Sen" gibi bir varsayılan isim atanabilir).
- Kullanıcının tarayıcısı değişmediği ve veriler silinmediği sürece, aynı cihazdan girdiğinde aynı katılımcı kimliği ile grup içinde kalmalıdır.

---

### Story 2: Link ile Davet / Katılım

**User Story:** Bir kullanıcı olarak, oluşturduğum grubu WhatsApp üzerinden link ile paylaşmak istiyorum, böylece arkadaşlarım kayıt olmadan bu gruba katılabilsin.

**Acceptance Criteria:**
- Grup oluşturulduktan sonra, ekranda paylaşılabilir bir grup linki gösterilmelidir.
- Kullanıcı için en azından "**Linki kopyala**" butonu olmalıdır (ek olarak "WhatsApp ile paylaş" kısa yolu olabilir).
- Link'e tıklayan yeni kullanıcıdan sadece görünen bir isim istenmelidir (email/şifre zorunlu değildir).
- Link ile gelen kullanıcı, isim girdikten sonra doğrudan ilgili grubun harcama listesini ve katılımcılarını görebilmelidir.

---

### Story 3: Hızlı Harcama Ekleme (Basit UI)

**User Story:** Bir grup üyesi olarak, tek bir sade ekrandan hızlıca harcama eklemek istiyorum, böylece kimse karmaşık form veya kayıt süreciyle uğraşmadan listeyi doldurabilelim.

**Acceptance Criteria:**
- Harcama ekleme alanı, aynı ekranda listeyle birlikte görünür olmalıdır (ayrı sayfaya gitmeye gerek kalmamalı).
- Minimum alanlar:
  - Tutar (`amount`, zorunlu, 0'dan büyük olmalı).
  - Açıklama (`description`, kısa metin).
- Ek alanlar:
  - Tarih (`date`) girilebilir, boş bırakılırsa otomatik olarak "bugün" atanmalıdır.
  - Ödeyen kişi (`payer`), gruptaki katılımcı listesinden seçilmelidir (varsayılan olarak son kullanan veya current user seçilebilir).
- Yeni harcama oluşturulduğunda, harcama listesi anında güncellenmelidir.
- Varsayılan para birimi TRY olmalıdır.

---

### Story 4: Eşit Bölüşme (MVP)

**User Story:** Bir grup üyesi olarak, her harcamayı seçili üyeler arasında otomatik eşit bölmek istiyorum, böylece manuel ayarlarla uğraşmadan herkesin payı ortaya çıksın.

**Acceptance Criteria:**
- MVP'de sadece **eşit bölüşme** desteklenir (`split_type = "equal"`).
- Her harcama için, katılımcı listesi varsayılan olarak tüm grup üyeleri olacak; kullanıcı isterse harcamaya dahil olmayanları kaldırabilmelidir.
- Harcama tutarı, seçili katılımcı sayısına eşit bölünmelidir.
- Kuruş farkları ilk katılımcıya eklenmelidir (örnek: 100 TL / 3 = 33.34 + 33.33 + 33.33).
- Bölüşme sonuçları dahili olarak saklanmalı, ileride borç/alışveriş hesabı yapılırken kullanılmalıdır.

---

### Story 5: Borç/Alacak Özeti (Kim Kime Ne Kadar Borçlu)

**User Story:** Bir grup üyesi olarak, gruptaki borç/alacak durumumu sade bir ekranda görmek istiyorum, böylece kimin kime ne kadar borçlu olduğunu hemen anlayabileyim.

**Acceptance Criteria:**
- Her üye için net bakiye (`balance`) hesaplanmalıdır:
  - Pozitif bakiye = o kişinin alacaklı olduğu tutar.
  - Negatif bakiye = o kişinin borçlu olduğu tutar.
- "Kim kime ne kadar borçlu" şeklinde sade bir liste/akış gösterilmelidir (örnek format: `Ali → Ayşe: 120 TL`).
- Minimum transfer sayısını azaltan basit bir optimizasyon (greedy yaklaşım) uygulanmalıdır; gereksiz küçük transferler gösterilmemelidir.
- Net bakiyesi 0 olan üyeler, "settled" olarak kabul edilmelidir (UI'da gizlenebilir veya "tamam" olarak işaretlenebilir).

---

## Gelecek Sürüm (MVP Sonrası) Epic: Kullanıcı Yönetimi (Auth)

> Not: Bu bölüm MVP kapsamı **dışındadır**. Email/şifre ile kayıt ve giriş, ileride ihtiyaç olursa aktive edilecektir.

### Future Story A: Kayıt (Sign Up)

**User Story:** Bir kullanıcı olarak, email ve şifremle kayıt olmak istiyorum, böylece gruplarım ve harcamalarım hesap bazlı olarak saklanabilsin.

**Acceptance Criteria (taslak):**
- Email formatı doğrulanmalı (RFC 5322).
- Şifre en az 8 karakter, 1 büyük harf ve 1 rakam içermeli.
- Aynı email ile birden fazla hesap oluşturulamamalı (409 Conflict).
- Başarılı kayıt sonrası JWT token dönülmeli.

---

### Future Story B: Giriş (Sign In)

**User Story:** Kayıtlı bir kullanıcı olarak, email ve şifremle giriş yapmak istiyorum, böylece geçmiş gruplarıma ve harcamalarıma erişebileyim.

**Acceptance Criteria (taslak):**
- Geçersiz email/şifre kombinasyonunda `401 Unauthorized` dönülmeli.
- Başarılı giriş sonrası JWT token dönülmeli.
- Token süresi 24 saat olmalı.

---

## Gelecek Sürüm (MVP Sonrası) Epic: İleri Harcama Özellikleri

### Future Story C: Özel Tutarla Bölüşme

**User Story:** Bir grup üyesi olarak, harcamayı farklı tutarlarla bölmek istiyorum, böylece herkes gerçek payını ödesin.

**Acceptance Criteria (taslak):**
- Her katılımcıya özel tutar girilebilmeli.
- Girilen tutarların toplamı harcama tutarına eşit olmalı (validation).
- Eşit olmadığında anlamlı bir hata mesajı gösterilmeli.

