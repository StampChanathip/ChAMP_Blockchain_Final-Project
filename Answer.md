ข้อเขียน

1. เป็นมาตรฐานในการสร้าง Token ในระบบ Blockchain ของ Ethereum โดยเป็นการกำหนด function พื้นฐานที่จำเป็นต้องมีใน smart contract ของแต่ละ Token เพื่อให้ developer สามารถเขียนคำสั่งเรียกใช้งาน function พื้นฐานได้ function ที่กำหนดไว้ก็เช่น transfer(), getbalance(), etc.

2. EVM เป็นเครื่องมือที่คอยทำการคำนวณและจัดการ state ต่างๆบน Blockchain ช่วยให้สามารถเขียนและเรียกใช้ smart contract บน chain ที่ใช้งาน EVM ได้อย่างเช่น Ethereum, Binanace Smart Chain,etc. ส่วน Non-EVM คือ Virtual Machine รูปแบบอื่นๆที่แต่ละ Blockchain เลือกใช้ เพื่อเสนอ value ที่แตกต่างจากตัว Ethereum อย่างเช่น Solana ที่ชูจุดเด่นในเรื่องของ Transaction Speed 
    
    EVM 
    - ข้อดี: ใช้ภาษา solidity ในการเขียน smart contract ซึ่งมีความใกล้เคียงกับ javascript ทำให้ developer ส่วนใหญ่เรียนรู้ได้ไม่ยาก และในปัจจุบันมีหลาย chain ที่เลือกใช้ EVM ทำให้การเขียน Dapp ที่ support EVM นำไปใช้เข้าได้กับหลากหลาย chain ที่ใช้ EVM เหมือนกัน
    - ข้อเสีย: ด้วยความที่ตัว EVM ผูกติดกับ Blockchain ที่หลักการทำงานโดยตั้งต้นมาจาก Ethereum การพัฒนา Dapp โดยใช้ EVM จึงพบปัญหาแบบเดียวกับปัญหาของตัว Ethereum เองเช่น ปัญหา Transaction speed หรือ Transaction Cost ที่สูงเป็นต้น
    
    Non-EVM 
    - ข้อดี: โดยหลักคิดของ Non-EVM ที่เลือกใช้เพื่อแก้ไขข้อด้อยของ EVM ทำให้ Dapp ที่พัฒนามี Transaction speed ที่สูงกว่าและ Transaction Cost ที่ต่ำกว่า EVM
    - ข้อเสีย: มีกำแพงในด้านการเรียนรู้ที่สูงกว่าเนื่องจากภาษาที่มช้เขียน smart contract บน Non-EVM ส่วนใหญ่มีความยากและซับซ้อนกว่า Solidity ที่ใช้บน EVM


3.

4. เนื่องจากแต่ละ Blockchain ไม่สามารถส่งข้อมูลข้ามไปมาหากันได้อย่างง่ายและปลอดภัย ทำให้ Token จะถูกจำกัดการใช้งานอยู่แค่บน chain ที่ Token นั้นๆถูกสร้างขึ้นเท่านั้น 

	ทำให้เกิด Wrapped Token ขึ้นมา เพื่อแก้ไขปัญหาข้างต้น ซึ่ง Wrapped Token จะเป็น ERC20 ที่สะท้อนมูลค่าแบบ 1:1 กับ Token ที่ถูก Wrap จาก chain อื่น เพื่อให้สามารถใช้งาน Token นั้นๆบน Ethereum ได้ ตัวอย่างเช่น wBTC ที่เป็น Wraped Token ของ Bitcoin โดยเราสามารถใช้ wBTC แทน BTC ในการทำธุรกรรมใน Dex บน Ethereum ได้ ซึ่ง Wrapped Token ก็สามารถมีได้ในทุก chain ไม่จำกัดเฉพาะ Ethereum

	Wrapped Token ทำให้การใช้งาน Dex สามารถทำได้หลากหลายมากขึ้นและทำให้ผู้ใช้งานสามารถนำ asset ที่มีไปเลือกใช้งานใน chain ต่างๆได้อย่างอิสระ
 
