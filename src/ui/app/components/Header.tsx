"use client";

import React from "react";
import Link from "next/link";

const Header = () => {

  return (
    <div
      className={'sticky inset-x-0 top-0 z-30 w-full transition-all bg-whitebackground shadow-lg'}
    >
      <div className="flex h-[40px] items-center justify-between px-4">
        <div className="flex items-center space-x-4">
          <Link
            href="/"
            className="flex flex-row space-x-3 items-center justify-center md:hidden"
          >
            <img src="images/logo.png" className="h-7 w-7 rounded-full" />
            <span className="font-bold text-xl flex ">Business Manager</span>
          </Link>
        </div>

        <div className="hidden md:block">
          <img src="images/user_profile.png" className="h-7 w-7 rounded-full" />
        </div>
      </div>
    </div>
  );
};

export default Header;
