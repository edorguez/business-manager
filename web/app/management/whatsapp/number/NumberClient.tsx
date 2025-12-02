"use client"

import getCurrentUser from "@/app/actions/getCurrentUser";
import SimpleCard from "@/app/components/cards/SimpleCard"
import useLoading from "@/app/hooks/useLoading";
import { CreateBusinessPhoneRequest, EditBusinessPhoneRequest, GetBusinessPhoneByCompanyIdRequest } from "@/app/services/whatsapp";
import { CurrentUser } from "@/app/types/auth";
import { BusinessPhone } from "@/app/types/whatsapp";
import { validNumbers, validPhone } from "@/app/utils/InputUtils";
import { Button, Input, useToast } from "@chakra-ui/react"
import { Icon } from "@iconify/react";
import { useCallback, useEffect, useState } from "react";

const NumberClient = () => {
  const toast = useToast();
  const isLoading = useLoading();
  const [isEditPhone, setIsEditPhone] = useState<boolean>(false);
  const [phone, setPhone] = useState<string>("");
  const [companyId, setComanyId] = useState<number>(0);
  const [isEditMode, setIsEditMode] = useState<boolean>(false);

  const handleNumberChange = (event: any) => {
    const { value } = event.target;
    if (value && !validNumbers(value)) return;
    setPhone(value);
  };

  const handleEditMode = () => {
    setIsEditMode((prev) => !prev);
  };

  const getBusinessPhone = useCallback(async () => {
    isLoading.onStartLoading();
    const currentUser: CurrentUser | null = getCurrentUser();

    if (currentUser) {
      setComanyId(currentUser.companyId);
      let data: BusinessPhone = await GetBusinessPhoneByCompanyIdRequest(
        currentUser.companyId
      );
      if (data?.phone) {
        setIsEditPhone(true);
        setPhone(data.phone);
      }
    }
    isLoading.onEndLoading();
  }, []);

  useEffect(() => {
    getBusinessPhone();
  }, []);

  const onSubmit = async () => {
    if (isFormValid()) {
      isLoading.onStartLoading();
      let res: any;

      if (isEditPhone) {
        let res: any = await EditBusinessPhoneRequest({
          companyId: companyId,
          phone: phone,
        });
      } else {
        let res: any = await CreateBusinessPhoneRequest({
          companyId: companyId,
          phone: phone,
        });
      }

      if (res?.error) {
        showErrorMessage(res.error);
        isLoading.onEndLoading();
      } else {
        showSuccessMessage("Número guardado exitosamente");
        isLoading.onEndLoading();
      }

      setIsEditMode(false);
    } else {
      showErrorMessage("El número es requerido y debe tener el siguiente formato: 04141234567");
    }
  };

  const isFormValid = (): boolean => {
    if (phone && !validPhone(phone)) return false;

    return true;
  };

  const showSuccessMessage = (msg: string) => {
    toast({
      title: "Whatsapp Business",
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
    <SimpleCard>
      <h1 className="font-bold">Whatsapp Business</h1>
      <div className="grid grid-cols-1 md:grid-cols-5 gap-2 ">
        <div className="md:col-span-4 mt-2">
          <label className="text-sm">Teléfono</label>
          <Input
            size="sm"
            name="phone"
            disabled={!isEditMode}
            value={phone}
            onChange={handleNumberChange}
            maxLength={11}
          />
        </div>
        <div className="flex items-end">
          {!isEditMode ? (
            <Button
              size="sm"
              variant="main"
              className="mx-1"
              onClick={handleEditMode}
            >
              <Icon icon="lucide:edit" />
            </Button>
          ) : (
            <>
              <Button
                size="sm"
                variant="third"
                className="mx-1"
                onClick={handleEditMode}
              >
                <Icon icon="material-symbols:cancel-outline" />
              </Button>
              <Button
                size="sm"
                variant="main"
                className="mx-1"
                onClick={onSubmit}
              >
                <Icon icon="lucide:check" />
              </Button>
            </>
          )}
        </div>
      </div>
      <p className="my-2 text-xs text-gray-600">
        Este número será usado para notificarte cuando tengas una nueva order y
        para que tus clientes se comuniquen contigo
      </p>
    </SimpleCard>
  );

}

export default NumberClient;
