'use client';

import Image from "next/image";

const UserBar = () => {
  return (
    <div className="bg-maincolor flex items-center p-1">
      <Image className="rounded-full" src={'https://marketplace.canva.com/print-mockup/bundle/E9Me4jcyzMX/fit:female,pages:double-sided,surface:marketplace/product:171618,surface:marketplace/EAFam5QLuIc/1/0/933w/canva-black-white-typography-motivation-tshirt-WKRZLU21i2c.png?sig=bc03703936ce8090247068bcf3a44f0e&width=800'} alt="" width={38} height={38} />
      <span className="text-white text-sm font-bold ml-2">Cliente 1</span>
    </div>
  )
}

export default UserBar;
