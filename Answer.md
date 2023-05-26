ข้อเขียน

1. เป็นมาตรฐานในการสร้าง Token ในระบบ Blockchain ของ Ethereum โดยเป็นการกำหนด function พื้นฐานที่จำเป็นต้องมีใน smart contract ของแต่ละ Token เพื่อให้ developer สามารถเขียนคำสั่งเรียกใช้งาน function พื้นฐานได้ function ที่กำหนดไว้ก็เช่น transfer(), getbalance(), etc.

2.

3.

4. เนื่องจากแต่ละ Blockchain ไม่สามารถส่งข้อมูลข้ามไปมาหากันได้อย่างง่ายและปลอดภัย ทำให้ Token จะถูกจำกัดการใช้งานอยู่แค่บน chain ที่ Token นั้นๆถูกสร้างขึ้นเท่านั้น 

	ทำให้เกิด Wrapped Token ขึ้นมา เพื่อแก้ไขปัญหาข้างต้น ซึ่ง Wrapped Token จะเป็น ERC20 ที่สะท้อนมูลค่าแบบ 1:1 กับ Token ที่ถูก Wrap จาก chain อื่น เพื่อให้สามารถใช้งาน Token นั้นๆบน Ethereum ได้ ตัวอย่างเช่น wBTC ที่เป็น Wraped Token ของ Bitcoin โดยเราสามารถใช้ wBTC แทน BTC ในการทำธุรกรรมใน Dex บน Ethereum ได้ ซึ่ง Wrapped Token ก็สามารถมีได้ในทุก chain ไม่จำกัดเฉพาะ Ethereum

	Wrapped Token ทำให้การใช้งาน Dex สามารถทำได้หลากหลายมากขึ้นและทำให้ผู้ใช้งานสามารถนำ asset ที่มีไปเลือกใช้งานใน chain ต่างๆได้อย่างอิสระ
 
