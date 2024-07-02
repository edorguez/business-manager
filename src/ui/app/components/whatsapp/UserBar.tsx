'use client';

import Image from "next/image";

interface UserBarProps {
  imageUrl: string,
  name: string
}

const UserBar: React.FC<UserBarProps> = ({
  imageUrl,
  name
}) => {
  const src = 'https://canto-wp-media.s3.amazonaws.com/app/uploads/2019/08/19194138/image-url-3.jpg';

  return (
    <div className="bg-maincolor flex items-center p-1 border-b-2 border-slate-200">
      <div className="w-[38px] h-[38px]">
        <Image className="rounded-full" alt={name} loader={() => src} src={src}  width={38} height={38}/>
      </div>
      <span className="text-white text-sm font-bold ml-2">{name}</span>
    </div>
  )
}

export default UserBar;
