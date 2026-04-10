# FAN IN/FAN OUT

![alt text](<Screenshot 2026-03-27 070621.png>)
![alt text](<Screenshot 2026-03-27 070655.png>) 
![alt text](<Screenshot 2026-03-27 070710.png>)
![alt text](<Screenshot 2026-03-27 070729.png>)

- For monitoring tools use ticker with go routines not sleep coz we need continous data

![alt text](<Screenshot 2026-03-29 025438.png>)


🔑 Context ke important types

1️⃣ Background

context.Background()


Base context.

2️⃣ WithCancel

Manual stop.

ctx, cancel := context.WithCancel(context.Background())

defer cancel()

3️⃣ WithTimeout

Time limit.

ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

4️⃣ WithDeadline

Exact time.

ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))