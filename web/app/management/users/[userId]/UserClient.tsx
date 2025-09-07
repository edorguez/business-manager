"use client";

import { Button, Input, Select, useToast } from "@chakra-ui/react";
import { BreadcrumItem } from "@/app/types";
import { Icon } from "@iconify/react";
import Link from "next/link";
import SimpleCard from "@/app/components/cards/SimpleCard";
import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import { useCallback, useEffect, useState } from "react";
import { useRouter, useParams, useSearchParams } from "next/navigation";
import useLoading from "@/app/hooks/useLoading";
import { EditUserRequest, GetUserRequest } from "@/app/services/users";
import { Role } from "@/app/types/role";
import { EditUser } from "@/app/types/user";
import { GetRolesRequest } from "@/app/services/roles";
import { PASSWORD, USER_ROLE_ID } from "@/app/constants";
import { CurrentUser } from "@/app/types/auth";
import getCurrentUser from "@/app/actions/getCurrentUser";
import {
  validEmail,
  validLettersAndNumbers,
  validWithNoSpaces,
} from "@/app/utils/InputUtils";

const UserClient = () => {
  const isLoading = useLoading();
  const toast = useToast();
  const params = useParams();
  const searchParams = useSearchParams();
  const { push } = useRouter();

  const bcItems: BreadcrumItem[] = [
    {
      label: "Usuarios",
      href: "/management/users",
    },
    {
      label: "Usuario",
      href: `/management/users/${params.userId}`,
    },
  ];

  const [isEdit, setIsEdit] = useState(false);
  const [roles, setRoles] = useState<Role[]>([]);
  const [formData, setFormData] = useState<EditUser>({
    id: 0,
    roleId: 0,
    email: "",
    password: null,
  });

  const getRoles = useCallback(async () => {
    isLoading.onStartLoading();
    let data: Role[] = await GetRolesRequest();
    setRoles(data.filter((x) => x.id !== USER_ROLE_ID.SUPER_ADMIN));
    isLoading.onEndLoading();
  }, []);

  const getUser = useCallback(async () => {
    isLoading.onStartLoading();
    let user: any = await GetUserRequest({ id: +params.userId });
    if (user) {
      const currentUser: CurrentUser | null = getCurrentUser();
      if (currentUser && user.id === currentUser.id) {
        push("/management/users");
      }

      setFormData({
        id: user.id,
        roleId: user.roleId,
        email: user.email,
        password: null,
      });
    }
    isLoading.onEndLoading();
  }, [params.customerId]);

  useEffect(() => {
    let paramIsEdit = searchParams.get("isEdit");
    if (paramIsEdit) {
      setIsEdit(true);
    }
    getUser();
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
      let editUser: any = await EditUserRequest(formData);
      if (!editUser?.error) {
        isLoading.onEndLoading();
        showSuccessEditMessage("Usuario editado exitosamente");
        push("/management/users");
      } else {
        showErrorMessage(editUser.error);
        isLoading.onEndLoading();
      }
    } else {
      showErrorMessage("Algunos campos son requeridos o inválidos");
    }
  };

  const isFormValid = (): boolean => {
    if (!formData.roleId) return false;
    if (!formData.email) return false;
    if (!validEmail(formData.email)) return false;
    if (
      formData.password &&
      formData.password.length < PASSWORD.MIN_PASSWORD_LEGTH
    )
      return false;
    if (formData.password && !validLettersAndNumbers(formData.password))
      return false;

    return true;
  };

  const showSuccessEditMessage = (msg: string) => {
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
          <h1 className="ml-2 font-bold">{`${
            isEdit ? "Editar" : ""
          } Usuario`}</h1>
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
              disabled={!isEdit}
            />
          </div>
          {isEdit && (
            <div className="mt-2">
              <label className="text-sm">Contraseña</label>
              <Input
                size="sm"
                name="password"
                value={formData.password ?? ""}
                onChange={handlePasswordChange}
                maxLength={20}
              />
            </div>
          )}

          <div className="mt-2">
            <label className="text-sm">
              Rol <span className="text-thirdcolor">*</span>
            </label>
            <Select
              size="sm"
              name="roleId"
              value={formData.roleId}
              onChange={handleChange}
              disabled={!isEdit}
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

      {isEdit && (
        <div className="mt-3">
          <Button variant="main" className="w-full" onClick={onSubmit}>
            Editar
          </Button>
        </div>
      )}
    </div>
  );
};

export default UserClient;
