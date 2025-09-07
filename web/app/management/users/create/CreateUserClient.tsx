"use client";

import getCurrentUser from "@/app/actions/getCurrentUser";
import { GetRolesRequest } from "@/app/services/roles";
import { CreateUserRequest } from "@/app/services/users";
import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import SimpleCard from "@/app/components/cards/SimpleCard";
import { PASSWORD, USER_ROLE_ID } from "@/app/constants";
import useLoading from "@/app/hooks/useLoading";
import { BreadcrumItem } from "@/app/types";
import { CurrentUser } from "@/app/types/auth";
import { Role } from "@/app/types/role";
import { CreateUser } from "@/app/types/user";
import { Button, Input, Link, Select, useToast } from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import { useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";
import {
  validEmail,
  validLettersAndNumbers,
  validWithNoSpaces,
} from "@/app/utils/InputUtils";

const CreateUserClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Userios",
      href: "/management/users",
    },
    {
      label: "Crear Usuario",
      href: "/management/users/create",
    },
  ];

  const isLoading = useLoading();
  const toast = useToast();
  const { push } = useRouter();
  const [roles, setRoles] = useState<Role[]>([]);
  const [formData, setFormData] = useState<CreateUser>({
    companyId: 0,
    roleId: 0,
    email: "",
    password: "",
  });

  const getRoles = useCallback(async () => {
    isLoading.onStartLoading();
    let data: Role[] = await GetRolesRequest();
    setRoles(data.filter((x) => x.id !== USER_ROLE_ID.SUPER_ADMIN));
    isLoading.onEndLoading();
  }, []);

  useEffect(() => {
    const currentUser: CurrentUser | null = getCurrentUser();
    if (currentUser) {
      formData.companyId = currentUser.companyId;
    }

    getRoles();
  }, [getRoles]);

  const handleChange = (event: any) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const handleEmailChange = (event: any) => {
    const { name, value } = event.target;
    if (value && !validWithNoSpaces(value)) return;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const handlePasswordChange = (event: any) => {
    const { name, value } = event.target;
    if (value && !validLettersAndNumbers(value)) return;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const onSubmit = async () => {
    if (isFormValid()) {
      isLoading.onStartLoading();
      let createUser: any = await CreateUserRequest(formData);
      if (!createUser.error) {
        isLoading.onEndLoading();
        showSuccessCreationMessage("Usuario creado exitosamente");
        push("/management/users");
      } else {
        showErrorMessage(createUser.error);
        isLoading.onEndLoading();
      }
    } else {
      showErrorMessage("Algunos campos son requeridos o inválidos");
    }
  };

  const isFormValid = (): boolean => {
    if (!formData.companyId) return false;
    if (!formData.roleId) return false;
    if (!formData.email) return false;
    if (!validEmail(formData.email)) return false;
    if (!formData.password) return false;
    if (formData.password.length < PASSWORD.MIN_PASSWORD_LEGTH) return false;
    if (!validLettersAndNumbers(formData.password)) return false;

    return true;
  };

  const showSuccessCreationMessage = (msg: string) => {
    toast({
      title: "Usuario",
      description: msg,
      variant: "customsuccess",
      position: "top-right",
      duration: 3000,
      isClosable: true,
    });
  };

  const showErrorMessage = (msg: string) => {
    toast({
      title: "Error",
      description: msg,
      variant: "customerror",
      position: "top-right",
      duration: 3000,
      isClosable: true,
    });
  };

  return (
    <div>
      <SimpleCard>
        <BreadcrumbNavigation items={bcItems} />
        <hr className="my-3" />
        <div className="flex items-center">
          <div>
            <Link href="/management/users">
              <div className="rounded p-2 hover:bg-thirdcolor hover:text-white duration-150">
                <Icon icon="fa-solid:arrow-left" />
              </div>
            </Link>
          </div>
          <h1 className="ml-2 font-bold">Crear Usuario</h1>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <div className="mt-2">
            <label className="text-sm">
              Correo <span className="text-thirdcolor">*</span>
            </label>
            <Input
              size="sm"
              name="email"
              value={formData.email}
              onChange={handleEmailChange}
              maxLength={100}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">
              Contraseña <span className="text-thirdcolor">*</span>
            </label>
            <Input
              size="sm"
              name="password"
              value={formData.password}
              onChange={handlePasswordChange}
              maxLength={20}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">
              Rol <span className="text-thirdcolor">*</span>
            </label>
            <Select
              size="sm"
              name="roleId"
              value={formData.roleId}
              onChange={handleChange}
            >
              <option value="">-</option>
              {roles &&
                roles.map((val: Role, index: number) => {
                  return (
                    <option key={index} value={val.id}>
                      {val.name}
                    </option>
                  );
                })}
            </Select>
          </div>
        </SimpleCard>
      </div>

      <div className="mt-3">
        <Button variant="main" className="w-full" onClick={onSubmit}>
          Crear
        </Button>
      </div>
    </div>
  );
};

export default CreateUserClient;
