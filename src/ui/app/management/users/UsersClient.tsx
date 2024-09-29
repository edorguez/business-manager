"use client";

import getCurrentUser from "@/app/actions/getCurrentUser";
import { DeleteUserRequest, GetUsersRequest } from "@/app/api/users/route";
import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import SimpleCard from "@/app/components/cards/SimpleCard";
import WarningModal from "@/app/components/modals/WarningModal";
import SimpleTable from "@/app/components/tables/SimpleTable";
import {
  ColumnType,
  SimpleTableColumn,
} from "@/app/components/tables/SimpleTable.types";
import useLoading from "@/app/hooks/useLoading";
import useWarningModal from "@/app/hooks/useWarningModal";
import { BreadcrumItem } from "@/app/types";
import { CurrentUser } from "@/app/types/auth";
import { User } from "@/app/types/user";
import { Button, useToast } from "@chakra-ui/react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";

const UsersClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Usuarios",
      href: "/management/users",
    },
  ];

  const usersCols: SimpleTableColumn[] = [
    {
      key: "roleName",
      name: "Rol",
      type: ColumnType.String,
    },
    {
      key: "email",
      name: "Correo",
      type: ColumnType.String,
    },
  ];

  const isLoading = useLoading();
  const toast = useToast();
  const { push } = useRouter();
  const [userData, setUserData] = useState<User[]>([]);
  const [offset, setOffset] = useState<number>(0);
  const deleteUserModal = useWarningModal();
  const [userIdDelete, setUserIdDelete] = useState<number>(0);

  const getUsers = useCallback(async () => {
    isLoading.onStartLoading();
    const currentUser: CurrentUser | null = getCurrentUser();

    if (currentUser) {
        let data: User[] = await GetUsersRequest({ companyId: currentUser.companyId, limit: 10, offset: offset });
        const formatData: User[] = data.map(x => {
        return {
          ...x,
          roleName: x.role.name
        }
      })
        setUserData(formatData);
    }
    isLoading.onEndLoading();
  }, [offset]);

  useEffect(() => {
    getUsers();
  }, [getUsers]);

  const handleChangePage = (val: string) => {
    setOffset((prevValue) =>
      val === "NEXT" ? (prevValue += 10) : (prevValue -= 10)
    );
  };

  const handleOpenDelete = (val: any) => {
    setUserIdDelete(val.id);
    deleteUserModal.onOpen();
  };

  const handleSubmitDelete = () => {
    onDelete(userIdDelete);
  };

  const onDelete = useCallback(async (id: number) => {
    await DeleteUserRequest({ id });
    getUsers();
    deleteUserModal.onClose();
    toast({
      title: "Usuario",
      description: "Usuario eliminado exitosamente",
      variant: "customsuccess",
      position: "top-right",
      duration: 3000,
      isClosable: true,
    });
  }, []);

  const handleOpenEdit = (val: any) => {
    push(`users/${val.id}?isEdit=true`);
  };

  const handleOpenDetail = (val: any) => {
    push(`users/${val.id}`);
  };

  return (
    <div>
      <SimpleCard>
        <WarningModal
          onSubmit={handleSubmitDelete}
          title="Eliminar Usuario"
          description="¿Estás seguro que quieres eliminar este usuario?"
          confirmText="Eliminar"
        />
        <BreadcrumbNavigation items={bcItems} />

        <hr className="mt-3" />

        <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 xl:grid-cols-6 gap-4">
          <div className="xl:col-start-6">
            <div className="flex flex-col">
              <span className="opacity-0">.</span>
              <Link href="/management/users/create">
                <Button size="sm" variant="main" className="w-full">
                  Crear Cliente
                </Button>
              </Link>
            </div>
          </div>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <SimpleTable
            columns={usersCols}
            data={userData}
            showDetails
            showEdit
            showDelete
            onEdit={handleOpenEdit}
            onDelete={handleOpenDelete}
            onDetail={handleOpenDetail}
            onChangePage={handleChangePage}
            offset={offset}
          />
        </SimpleCard>
      </div>
    </div>
  );
};

export default UsersClient;
