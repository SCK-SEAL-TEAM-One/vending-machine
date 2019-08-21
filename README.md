# vending-machine

## Directory Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
product
```

## File Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
product.go
```

## Package Name
- ใช้ตัวอักษรพิมพ์เล็กทั้งหมด เช่น
```
product
```

## Test Function Name
- ใช้รูปแบบการตั้งชื่อฟังก์ชันเป็นแบบ **Snake_Case** เช่น
```
Test_GetProduct_Input_Product_ID_B00UEIC4ZC_Should_Be_Marvel_Avengers_8_Character_Ultimate_Protectors Pack
```

## Variable Name
- ชื่อตัวแปรเป็น **camelCase** เช่น
```
name, price, quantity
```

- ชื่อตัวแปรเก็บค่าที่เป็นพหูพจน์ ให้เติม s ต่อท้ายตัวแปรเสมอ เช่น
```
images
```

- ชื่อตัวแปร struct ให้ตั้งชื่อขึ้นต้นคำแรกด้วยตัวอักษรพิมพ์ใหญ่ ในรูปแบบ **camelCase** เช่น
```
Prodruct, Customer
```

- ชื่อตัวแปร Constant ให้ตังชื่อเป็นตัวพิมพ์เล็กก่อน เว้นแต่เมื่อมีการใช้ข้าม package ถึงจะใช้ Capital Case เช่น
```
Months, Minute, threshold
```

## รูปแบบข้อมูล json 

ใช้เป็น **snakeCase** เช่น
```
shipping_weight, product_id
```

# Error Message Pattern
- ใช้รูปแบบ verb + noun + "error" เช่น
```
Get product detail error
```

## ข้อตกลง Commit Message ร่วมกัน
`[Created]: สร้างไฟล์ใหม่`

`[Edited]: แก้ไข code ในไฟล์เดิมที่มีอยู่แล้ว รวมถึงกรณี refactor code`

`[Added]: กรณีเพิ่ม function, function test ใหม่เข้ามา`

`[Deleted]: ลบไฟล์ออก`

* ให้เขียนรายละเอียดด้วยว่าแก้ไขอะไรและทำที่ตรงไหน
