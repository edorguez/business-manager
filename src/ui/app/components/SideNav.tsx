"use client";

import React, { useEffect, useState } from "react";

import Link from "next/link";
import { usePathname } from "next/navigation";

import { SIDENAV_ITEMS } from "../constants";
import { SideNavItem } from "../types";
import { Icon } from "@iconify/react";
import Image from "next/image";
import { CurrentUser } from "../types/auth";
import getCurrentUser from "../actions/getCurrentUser";
import { validateUserInRoles } from "../utils/Utils";

const SideNav = () => {
  const [currentUser, setCurrentUser] = useState<CurrentUser>();

  useEffect(() => {
    const user: CurrentUser | null = getCurrentUser();
    if (user) {
      setCurrentUser(user);
    }
  }, []);

  return (
    <div className="md:w-60 bg-whitebackground shadow-lg h-screen flex-1 fixed hidden md:flex">
      <div className="flex flex-col space-y-6 w-full">
        <Link
          href="/management/home"
          className="flex flex-row space-x-3 items-center justify-center md:justify-start md:px-6 h-12 w-full"
        >
          <Image src="/images/logo.png" alt="Logo" width={28} height={28} />
          <span className="font-bold text-sm hidden md:flex text-thirdcolor">
            Business Manager
          </span>
        </Link>

        <div className="flex flex-col space-y-2 md:px-3 ">
          {SIDENAV_ITEMS.map((item, idx) => {
            if (
              currentUser &&
              validateUserInRoles(item.roleIds, currentUser.roleId)
            )
              return <MenuItem key={idx} item={item} />;
            else return null;
          })}
        </div>
      </div>
    </div>
  );
};

export default SideNav;

const MenuItem = ({ item }: { item: SideNavItem }) => {
  const pathname = usePathname();
  const [subMenuOpen, setSubMenuOpen] = useState(false);
  const toggleSubMenu = () => {
    setSubMenuOpen(!subMenuOpen);
  };

  return (
    <div className="">
      {item.submenu ? (
        <>
          <button
            onClick={toggleSubMenu}
            className={`flex flex-row items-center p-2 rounded-md w-full justify-between hover:bg-thirdcolorhov ${
              pathname.includes(item.path) ? "bg-darksecondary" : ""
            }`}
          >
            <div className="flex flex-row space-x-4 items-center text-iconcolor">
              {item.icon}
              <span className="text-xs text-black flex">{item.title}</span>
            </div>

            <div
              className={`${
                subMenuOpen ? "rotate-180" : ""
              } flex text-iconcolor`}
            >
              <Icon icon="lucide:chevron-down" width="20" height="20" />
            </div>
          </button>

          {subMenuOpen && (
            <div className="my-2 ml-12 flex flex-col space-y-4">
              {item.subMenuItems?.map((subItem, idx) => {
                return (
                  <Link key={idx} href={subItem.path}>
                    <span className="text-black text-xs">{subItem.title}</span>
                  </Link>
                );
              })}
            </div>
          )}
        </>
      ) : (
        <Link
          href={item.path}
          className={`flex flex-row space-x-4 items-center p-2 rounded-md hover:bg-thirdcolorhov ${
            item.path === pathname
              ? "bg-thirdcolorhov text-thirdcolor"
              : "text-iconcolor"
          }`}
        >
          {item.icon}
          <span
            className={`text-xs flex ${
              item.path === pathname
                ? "text-thirdcolor font-bold"
                : "text-black"
            }`}
          >
            {item.title}
          </span>
        </Link>
      )}
    </div>
  );
};
