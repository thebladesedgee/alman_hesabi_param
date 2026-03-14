# Alman Hesabi - User Stories

> Tricount benzeri bir hesap bölüşme uygulaması.
> Bu dosya, AI-native geliştirme sürecinde "Source of Truth" görevi görür.

---

## Epic: Kullanıcı Yönetimi

### Story 1: Kayıt (Sign Up)

**User Story:** Bir kullanıcı olarak, email ve şifremle kayıt olmak istiyorum, böylece harcama gruplarına katılabileyim.

**Acceptance Criteria:**
- Email formatı doğrulanmalı (RFC 5322).
- Şifre en az 8 karakter, 1 büyük harf ve 1 rakam içermeli.
- Aynı email ile birden fazla hesap oluşturulamamalı (409 Conflict).
- Başarılı kayıt sonrası JWT token dönülmeli.

**Technical Hints:**
- `POST /api/auth/register` endpoint.
- Şifreler bcrypt ile hash'lenmeli.
- `users` tablosu: `id`, `name`, `email`, `password_hash`, `created_at`.

---

### Story 2: Giriş (Sign In)

**User Story:** Kayıtlı bir kullanıcı olarak, email ve şifremle giriş yapmak istiyorum, böylece mevcut gruplarıma erişebileyim.

**Acceptance Criteria:**
- Geçersiz email/şifre kombinasyonunda `401 Unauthorized` dönülmeli.
- Başarılı giriş sonrası JWT token dönülmeli.
- Token süresi 24 saat olmalı.

**Technical Hints:**
- `POST /api/auth/login` endpoint.
- JWT secret `.env` dosyasından okunmalı.

---

## Epic: Grup Yönetimi

### Story 3: Grup Oluşturma

**User Story:** Bir kullanıcı olarak, yeni bir harcama grubu oluşturmak istiyorum, böylece arkadaşlarımla harcamaları paylaşabileyim.

**Acceptance Criteria:**
- Grup adı zorunlu (min 1, max 100 karakter).
- Grup oluşturan kişi otomatik olarak üye olmalı.
- Benzersiz bir davet kodu (invite code) üretilmeli.
- Grup para birimi seçilebilmeli (varsayılan: TRY).

**Technical Hints:**
- `POST /api/groups` endpoint.
- `groups` tablosu: `id`, `name`, `currency`, `invite_code`, `created_by`, `created_at`.
- `group_members` tablosu: `group_id`, `user_id`, `joined_at`.

---

### Story 4: Gruba Katılma

**User Story:** Bir kullanıcı olarak, davet kodu ile bir gruba katılmak istiyorum, böylece grubun harcamalarını görebilip ekleyebileyim.

**Acceptance Criteria:**
- Geçerli bir davet kodu ile gruba katılım sağlanmalı.
- Geçersiz kod girildiğinde `404 Not Found` dönülmeli.
- Zaten üye olan kullanıcıya uygun hata mesajı gösterilmeli.
- Katılım sonrası grup detayları ve mevcut üye listesi dönülmeli.

**Technical Hints:**
- `POST /api/groups/join` endpoint, body: `{ "invite_code": "ABC123" }`.

---

## Epic: Harcama Takibi

### Story 5: Harcama Ekleme

**User Story:** Bir grup üyesi olarak, yaptığım harcamayı gruba eklemek istiyorum, böylece masraf diğer üyelerle paylaşılsın.

**Acceptance Criteria:**
- Tutar (amount), açıklama (description) ve tarih (date) girilmeli.
- Harcamayı ödeyen kişi (payer) seçilmeli.
- Harcamaya katılan üyeler seçilebilmeli (varsayılan: tüm grup üyeleri).
- Tutar 0'dan büyük olmalı.
- Bölüşme tipi seçilebilmeli: eşit (equal) veya özel (custom).

**Technical Hints:**
- `POST /api/groups/:groupId/expenses` endpoint.
- `expenses` tablosu: `id`, `group_id`, `description`, `amount`, `currency`, `paid_by`, `split_type`, `date`, `created_at`.
- `expense_splits` tablosu: `expense_id`, `user_id`, `amount`.

---

### Story 6: Eşit Bölüşme

**User Story:** Bir grup üyesi olarak, harcamayı seçilen üyeler arasında eşit bölmek istiyorum, böylece herkes eşit pay ödesin.

**Acceptance Criteria:**
- Harcama tutarı, seçilen katılımcı sayısına eşit bölünmeli.
- Kuruş farkları ilk katılımcıya eklenmeli (ör: 100 TL / 3 = 33.34 + 33.33 + 33.33).
- Bölüşme detayları `expense_splits` tablosuna kaydedilmeli.

**Technical Hints:**
- `split_type: "equal"` olarak işaretlenmeli.
- Backend'de `services/split.go` içinde hesaplama yapılmalı.

---

### Story 7: Özel Tutarla Bölüşme

**User Story:** Bir grup üyesi olarak, harcamayı farklı tutarlarla bölmek istiyorum, böylece herkes gerçek payını ödesin.

**Acceptance Criteria:**
- Her katılımcıya özel tutar girilebilmeli.
- Girilen tutarların toplamı harcama tutarına eşit olmalı (validation).
- Eşit olmadığında hata mesajı gösterilmeli.

**Technical Hints:**
- `split_type: "custom"` olarak işaretlenmeli.

---

### Story 8: Borç/Alacak Özeti

**User Story:** Bir grup üyesi olarak, gruptaki borç/alacak durumumu görmek istiyorum, böylece kimin kime borçlu olduğunu bileyim.

**Acceptance Criteria:**
- Her üye için net bakiye (balance) gösterilmeli.
- "Kim kime ne kadar borçlu" listesi oluşturulmalı.
- Minimum transfer sayısı ile borçlar optimize edilmeli (debt simplification).
- Bakiye sıfır olan üyeler "settled" olarak işaretlenmeli.

**Technical Hints:**
- `GET /api/groups/:groupId/balances` endpoint.
- Backend'de `services/balance.go` içinde hesaplama yapılmalı.
- Greedy algoritma ile minimum transfer optimizasyonu.
