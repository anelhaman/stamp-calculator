# สถานการณ์สมมุติ เรื่องราวการซื้อสแตมป์

วันก่อนได้มีโอกาสไปที่บัณฑิตวิทยาลัยเพื่อแจ้งว่าผมจะไม่เข้ารับพิธีพระราชทานปริญญาโทในปีนี้ เนื่องจากมีภารกิจบางอย่างที่ไม่สามารถหลีกเลี่ยงได้ จึงต้องการให้มหาวิทยาลัยส่งปริญญาบัตรทางไปรษณีย์มาที่บ้านแทน

เมื่อไปถึงที่บัณฑิตวิทยาลัย เจ้าหน้าที่แจ้งว่าในการส่งปริญญาบัตรทางไปรษณีย์นั้น จำเป็นต้องใช้สแตมป์รวม 70 บาท ซึ่งจะเป็นค่าใช้จ่ายในการจัดส่ง

หลังจากทราบข้อมูลดังนี้ ผมจึงต้องไปซื้อสแตมป์ที่เคาน์เตอร์ไปรษณีย์ แต่ปัญหาคือ สแตมป์ที่มีในร้านนั้นมีหลายราคาตามที่กำหนด ได้แก่ 9 บาท, 5 บาท และ 3 บาท ซึ่งทำให้ผมต้องคิดว่า จะต้องใช้สแตมป์กี่ดวงและราคาอะไรบ้างที่จะรวมกันได้พอดีเป็น 70 บาท

จริง ๆ แล้วการคำนวณสแตมป์ไม่ได้เป็นเรื่องที่ต้องจริงจังเท่าไหร่ เพราะมีเจ้าหน้าที่ไปรษณีย์คอยคำนวณให้ แต่ผมต้องการเช็คเพื่อความถูกต้อง ระหว่างที่เช็คก็ใช้เวลาคิดอยู่พักใหญ่ เพราะมันเป็นไปได้หลายแบบจริง ๆ

ผมเริ่มคิดว่าเราควรจะหาความเป็นไปได้ในการเลือกใช้สแตมป์จากราคาต่าง ๆ ที่มีอยู่เพื่อให้รวมกันได้พอดีตามจำนวนที่กำหนด

หลังจากการคิดและคำนวณ ผมจึงตัดสินใจเขียนโปรแกรมเพื่อช่วยคำนวณหาชุดสแตมป์ที่สามารถรวมกันได้พอดีกับจำนวนเงิน 70 บาท โดยไม่มากหรือน้อยกว่านั้น ซึ่งจะช่วยให้การเลือกซื้อสแตมป์เป็นไปได้อย่างสะดวกและรวดเร็ว

```log
Combination #1
---------------------------------------------
Stamp 9 THB:	6	=	54	THB
Stamp 5 THB:	2	=	10	THB
Stamp 3 THB:	2	=	6	THB
---------------------------------------------
Total:				70	THB

Combination #2
---------------------------------------------
Stamp 9 THB:	1	=	9	THB
Stamp 5 THB:	5	=	25	THB
Stamp 3 THB:	12	=	36	THB
---------------------------------------------
Total:				70	THB

Combination #3
---------------------------------------------
Stamp 9 THB:	1	=	9	THB
Stamp 5 THB:	2	=	10	THB
Stamp 3 THB:	17	=	51	THB
---------------------------------------------
Total:				70	THB

Combination #4
---------------------------------------------
Stamp 9 THB:	3	=	27	THB
Stamp 5 THB:	8	=	40	THB
Stamp 3 THB:	1	=	3	THB
---------------------------------------------
Total:				70	THB

Combination #5
---------------------------------------------
Stamp 5 THB:	5	=	25	THB
Stamp 3 THB:	15	=	45	THB
---------------------------------------------
Total:				70	THB

Combination #6
---------------------------------------------
Stamp 9 THB:	4	=	36	THB
Stamp 5 THB:	5	=	25	THB
Stamp 3 THB:	3	=	9	THB
---------------------------------------------
Total:				70	THB

Combination #7
---------------------------------------------
Stamp 9 THB:	3	=	27	THB
Stamp 5 THB:	5	=	25	THB
Stamp 3 THB:	6	=	18	THB
---------------------------------------------
Total:				70	THB

Combination #8
---------------------------------------------
Stamp 5 THB:	14	=	70	THB
---------------------------------------------
Total:				70	THB

Combination #9
---------------------------------------------
Stamp 9 THB:	1	=	9	THB
Stamp 5 THB:	11	=	55	THB
Stamp 3 THB:	2	=	6	THB
---------------------------------------------
Total:				70	THB

Combination #10
---------------------------------------------
Stamp 9 THB:	2	=	18	THB
Stamp 5 THB:	2	=	10	THB
Stamp 3 THB:	14	=	42	THB
---------------------------------------------
Total:				70	THB

Combination #11
---------------------------------------------
Stamp 5 THB:	2	=	10	THB
Stamp 3 THB:	20	=	60	THB
---------------------------------------------
Total:				70	THB

Combination #12
---------------------------------------------
Stamp 9 THB:	4	=	36	THB
Stamp 5 THB:	2	=	10	THB
Stamp 3 THB:	8	=	24	THB
---------------------------------------------
Total:				70	THB

Combination #13
---------------------------------------------
Stamp 9 THB:	3	=	27	THB
Stamp 5 THB:	2	=	10	THB
Stamp 3 THB:	11	=	33	THB
---------------------------------------------
Total:				70	THB

Combination #14
---------------------------------------------
Stamp 9 THB:	5	=	45	THB
Stamp 5 THB:	5	=	25	THB
---------------------------------------------
Total:				70	THB

Combination #15
---------------------------------------------
Stamp 9 THB:	5	=	45	THB
Stamp 5 THB:	2	=	10	THB
Stamp 3 THB:	5	=	15	THB
---------------------------------------------
Total:				70	THB

Combination #16
---------------------------------------------
Stamp 9 THB:	2	=	18	THB
Stamp 5 THB:	8	=	40	THB
Stamp 3 THB:	4	=	12	THB
---------------------------------------------
Total:				70	THB

Total combinations: 16

```

เพราะมันคำนวณยากแบบนี้เองหรือเปล่า เพื่อนจึงเลือกเข้าพิธีฯ แทนที่จะส่งไปรษณีย์! 😜