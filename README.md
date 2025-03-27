# TestWallet

# Database 

database docker run --name MyWallet -e POSTGRES_PASSWORD=Samurai7 -d -p 5432:5432  postgres

create table if not exists public.transactions
(
    transaction_id varchar,
    user_id        integer,
    amount         integer,
    payment_method varchar(50),
    status         varchar(50),
    expires_at     timestamp,
    balance        integer
);

create table if not exists public.users
(
    user_id        integer,
    amount         integer,
    payment_method varchar(50)
);

## เหตุผลที่ให้ Amont เป็น integerเพราะว่าถ้าเป็น fload64 จะเกิดปัญหา FloatingPointError เลยใช้เป็น int แทนครับ 
## ฟังชัััน has ที่ไม่ได้ทำไปเพราะคิดว่าใช้้เวลานานหว่านี้ครับ ส่วนตัวทีเคยทำมาของ payment-gateway ใช้ของ ais(mapy) ครับ
## วิธีที่ใช้ของmpay คือเค้าจะให้ secretkeyมา และให้ตัวอย่างเข้ารหัสมาครับ ใช้รูปแบบเอาdatajson+sceretkey ทำเป็น signatureส่งไปพร้อมกับdatajson ตอนแกะออกก็เอา datajason ที่เค้าส่งมาเอามาเข้ารหัสก่อนดูว่าได้ signature ตัวเดียวกันมั้ยถ้าใช้ก็สามารถเชื่อถือdataได้ครับ
## Prevent duplicate confirmations ผมไม่แน่ใจว่ามันคือเรื่อง Idempotent หรือป่าวครับ 

## วิธีการ Run
    1  สรา้งdatabaseใส่ข้อมูล
    2  install lip
    3  run go (go run main.go)
    4  http://localhost:8080/users
    5  http://localhost:8080/wallet/verify { "amount":10050,    "user_id":1, "payment_method": "credit_card"}
    6  http://localhost:8080/wallet/confirm {"transaction_id": "abc123"}
