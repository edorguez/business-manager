"use client";

import React from "react";
import Link from "next/link";
import Image from "next/image";

const Header = () => {

  return (
    <div className={'sticky inset-x-0 top-0 z-30 w-full transition-all bg-whitebackground shadow-lg'}>
      <div className="flex h-[40px] items-center justify-between px-4">
        <div className="flex items-center space-x-4">
          <Link href="/management/home" className="flex flex-row space-x-3 items-center justify-center md:hidden">
            <Image src='/images/logo.png' alt="Logo" width={28} height={28} />
            <span className="font-bold text-xl flex ">Business Manager</span>
          </Link>
        </div>

        <div className="hidden md:block">
          <Image src='/images/user_profile.png' alt="User Profile" width={28} height={28} />
        </div>
      </div>
    </div>
  );
};

export default Header;
