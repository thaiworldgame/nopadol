## Database Model : ERP Project 
### Customer 

* add  

        สร้างรายการลูกค้าใหม่ 
* update 

          update ข้อมูลลูกค้าที่มีอยู่ 
* inactive

       ปรับสถานะเป็น ไม่ใช้งาน 
* changecode

        เปลียนรหัสเป็นรหัสใหม่ 
* getByCode

        ดึงข้อมูลจากฐานข้อมูลมาเก็บไว้ที่ Object โดยเรียกจากรหัสลูกค้า
* getByID
        
        ดึงข้อมูลจากฐานข้อมูลเก็บใน Object โดยอ้างอิงจาก ID
        
### item 
* add 
* update 
* inactive
* changecode