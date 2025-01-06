"use client";

import React, { useCallback, useEffect, useState } from "react";
import Link from "next/link";
import Image from "next/image";
import {
  Avatar,
  Menu,
  MenuButton,
  MenuDivider,
  MenuItem,
  MenuList,
} from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import { useRouter } from "next/navigation";
import deleteUserSession from "../actions/deleteUserSession";
import useChangeCompanyImage from "../hooks/useChangeCompanyImage";
import { CurrentUser } from "../types/auth";
import getCurrentUser from "../actions/getCurrentUser";
import { Company } from "../types/company";
import { GetCompanyRequest } from "../services/companies";

const defaultUserImage: string = '/images/account/user.png';

const Header = () => {
  const { push } = useRouter();
  const { emitSignal, resetSignal } = useChangeCompanyImage();
  const [userImage, setUserImage] = useState<string>(
    defaultUserImage
  );

  const onCloseSession = () => {
    deleteUserSession();
    push("/login");
  };

  const onAccount = () => {
    push("/management/account");
  };

  useEffect(() => {
    getCompanyImage();
  }, []);

  useEffect(() => {
    if (emitSignal) {
      getCompanyImage();
      resetSignal();
    }
  }, [emitSignal, resetSignal]);

  const getCompanyImage = useCallback(async () => {
    const currentUser: CurrentUser | null = getCurrentUser();
    if (currentUser) {
      let company: Company = await GetCompanyRequest({
        id: currentUser?.companyId,
      });
      if (company?.imageUrl) {
        setUserImage(company.imageUrl);
      } else {
        setUserImage(defaultUserImage);
      }
    }
  }, []);

  return (
    <div
      className={
        "sticky inset-x-0 top-0 z-30 w-full transition-all bg-whitebackground shadow-lg"
      }
    >
      <div className="flex h-[40px] items-center justify-between px-4">
        <div className="flex items-center space-x-4">
          <Link
            href="/management/home"
            className="flex flex-row space-x-3 items-center justify-center md:hidden"
          >
            <Image src={userImage} alt="Logo" width={28} height={28} />
            <span className="font-bold text-xl flex ">Business Manager</span>
          </Link>
        </div>

        <div className="hidden md:block">
          <Menu>
            <MenuButton>
              <Avatar name="User Profile" src={userImage} size="sm" />
            </MenuButton>
            <MenuList>
              <MenuItem onClick={onAccount}>
                <Icon icon="mdi:user" />
                Cuenta
              </MenuItem>
              <MenuDivider />
              <MenuItem onClick={onCloseSession}>
                <Icon icon="ci:exit" />
                Cerrar Sesión
              </MenuItem>
            </MenuList>
          </Menu>
        </div>
      </div>
    </div>
  );
};

export default Header;
