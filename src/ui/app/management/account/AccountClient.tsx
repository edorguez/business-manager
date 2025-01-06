"use client";

import React, { useCallback, useEffect, useState } from "react";
import {
  Box,
  Button,
  Container,
  Flex,
  FormControl,
  Heading,
  Input,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Tabs,
  Text,
  VStack,
  Avatar,
  useToast,
} from "@chakra-ui/react";
import SimpleCard from "@/app/components/cards/SimpleCard";
import ImagesUpload from "@/app/components/uploads/ImagesUpload";
import { Company, EditCompany } from "@/app/types/company";
import {
  EditCompanyRequest,
  GetCompanyRequest,
} from "@/app/services/companies";
import { CurrentUser } from "@/app/types/auth";
import getCurrentUser from "@/app/actions/getCurrentUser";
import useLoading from "@/app/hooks/useLoading";
import { EditEmail, EditPassword, User } from "@/app/types/user";
import {
  EditEmailRequest,
  EditPasswordRequest,
  GetUserRequest,
} from "@/app/services/users";
import useWarningModal from "@/app/hooks/useWarningModal";
import WarningModal from "@/app/components/modals/WarningModal";
import { useRouter } from "next/navigation";
import deleteUserSession from "@/app/actions/deleteUserSession";
import isUserAdmin from "@/app/actions/isUserAdmin";
import { PASSWORD } from "@/app/constants";
import {
  validEmail,
  validLettersAndNumbers,
  validWithNoSpaces,
} from "@/app/utils/InputUtils";
import { formatCompanyNameToUrlName } from "@/app/utils/Utils";
import useChangeCompanyImage from "@/app/hooks/useChangeCompanyImage";

const AccountClient = () => {
  const { push } = useRouter();
  const isLoading = useLoading();
  const toast = useToast();
  const [isAdmin, setIsAdmin] = useState<boolean>(false);
  const { triggerSignal } = useChangeCompanyImage();
  const [imagesToSave, setImagesToSave] = useState<File[]>([]);
  const [imagesLoaded, setImagesLoaded] = useState<File[]>([]);
  const [companyImageUrl, setCompanyImageUrl] = useState<string>("/images/account/user.png");
  const confirmChangeModal = useWarningModal();

  const [companyFormData, setCompanyFormData] = useState<EditCompany>({
    id: 0,
    name: "",
    nameFormatUrl: "",
  });
  const [emailFormData, setEmailFormData] = useState<EditEmail>({
    id: 0,
    email: "",
  });
  const [passwordFormData, setPasswordFormData] = useState<EditPassword>({
    id: 0,
    password: "",
    passwordRepeat: "",
  });

  const handleCompanyFormChange = (event: any) => {
    const { name, value } = event.target;

    if (value && !validLettersAndNumbers(value, true)) return;

    let formatUrl = formatCompanyNameToUrlName(value);
    setCompanyFormData((prev) => ({
      ...prev,
      [name]: value,
      ["nameFormatUrl"]: formatUrl,
    }));
  };

  const handleEmailFormChange = (event: any) => {
    const { name, value } = event.target;
    if (value && !validWithNoSpaces(value)) return;
    setEmailFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handlePasswordFormChange = (event: any) => {
    const { name, value } = event.target;
    if (value && !validLettersAndNumbers(value)) return;
    setPasswordFormData((prev) => ({ ...prev, [name]: value }));
  };

  const getCompany = useCallback(async () => {
    const currentUser: CurrentUser | null = getCurrentUser();
    if (currentUser) {
      isLoading.onStartLoading();
      let company: Company = await GetCompanyRequest({
        id: currentUser?.companyId,
      });
      if (company) {
        setCompanyFormData({
          id: company.id ?? 0,
          name: company.name ?? "",
          nameFormatUrl: company.nameFormatUrl ?? "",
        });

        if (company.imageUrl) {
          // Convert image URL to File[] type
          setCompanyImageUrl(company.imageUrl);
          const response = await fetch(company.imageUrl);
          const blob = await response.blob();
          const fileName = company.imageUrl.split("/").pop() || "image.jpg";
          const fileExtension = fileName.split(".").pop();
          const fileBaseName = fileName.split(".").slice(0, -1).join(".");
          const fileImage = new File(
            [blob],
            `${fileBaseName}.${fileExtension}`,
            { type: blob.type }
          );
          setImagesLoaded([fileImage]);
        }
      }

      isLoading.onEndLoading();
    }
  }, []);

  const getUser = useCallback(async () => {
    const currentUser: CurrentUser | null = getCurrentUser();
    if (currentUser) {
      isLoading.onStartLoading();
      let user: User = await GetUserRequest({
        id: currentUser?.id,
      });
      if (user) {
        setEmailFormData({
          id: user.id ?? 0,
          email: user.email ?? "",
        });
        setPasswordFormData({
          id: user.id ?? 0,
          password: "",
          passwordRepeat: "",
        });
      }

      isLoading.onEndLoading();
    }
  }, []);

  useEffect(() => {
    const user: CurrentUser | null = getCurrentUser();
    if (user) setIsAdmin(isUserAdmin(user.roleId));

    getCompany();
    getUser();
  }, [getCompany, getUser]);

  const handleSpaceKeyDown = (event: any) => {
    if (event.key === " ") {
      event.preventDefault();
    }
  };

  const isCompanyFormValid = (): boolean => {
    if (!companyFormData.name) return false;
    if (!companyFormData.nameFormatUrl) return false;
    if (!validLettersAndNumbers(companyFormData.name, true)) return false;

    return true;
  };

  const isEmailFormValid = (): boolean => {
    if (!emailFormData.email) return false;
    if (!validEmail(emailFormData.email)) return false;

    return true;
  };

  const isPasswordFormValid = (): boolean => {
    if (!passwordFormData.password || !passwordFormData.passwordRepeat)
      return false;
    if (passwordFormData.password !== passwordFormData.passwordRepeat)
      return false;
    if (passwordFormData.password.length < PASSWORD.MIN_PASSWORD_LEGTH)
      return false;
    if (!validLettersAndNumbers(passwordFormData.password)) return false;

    return true;
  };

  const onCompanySubmit = async () => {
    if (isCompanyFormValid()) {
      isLoading.onStartLoading();
      let editCompany: any = await EditCompanyRequest(
        companyFormData,
        imagesToSave
      );
      if (editCompany?.error) {
        isLoading.onEndLoading();
        showErrorMessage(editCompany.error);
      } else {
        triggerSignal();
        isLoading.onEndLoading();
        showSuccessCreationMessage("Empresa modificada exitosamente");
      }
    } else {
      showErrorMessage("Algunos campos son requeridos o inválidos");
    }
  };

  const onEmailSubmit = async () => {
    confirmChangeModal.onClose();
    if (isEmailFormValid()) {
      isLoading.onStartLoading();
      let editEmail: any = await EditEmailRequest(emailFormData);
      if (editEmail?.error) {
        isLoading.onEndLoading();
        showErrorMessage("Correo ya está en uso");
      } else {
        isLoading.onEndLoading();
        showSuccessCreationMessage("Correo actualizado exitosamente");
        deleteUserSession();
        push("/login");
      }
    } else {
      showErrorMessage("Algunos campos son requeridos o inválidos");
    }
  };

  const onPasswordSubmit = async () => {
    if (isPasswordFormValid()) {
      isLoading.onStartLoading();
      let editPassword: any = await EditPasswordRequest(passwordFormData);
      if (editPassword?.error) {
        isLoading.onEndLoading();
        showErrorMessage(editPassword.error);
      } else {
        isLoading.onEndLoading();
        showSuccessCreationMessage("Contraseña cambiada exitosamente");
        setPasswordFormData((prev) => ({
          ...prev,
          password: "",
          passwordRepeat: "",
        }));
      }
    } else {
      showErrorMessage(
        "La contraseña es requerida, mínimo 6 caracteres o las contraseñas no son iguales"
      );
    }
  };

  const showSuccessCreationMessage = (msg: string) => {
    toast({
      title: "Empresa",
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

  const handleUploadFiles = (files: File[]) => {
    if (files && files.length > 0) {
      const reader = new FileReader();
      reader.onload = (event) => {
        const imageUrl = event.target?.result as string;
        setCompanyImageUrl(imageUrl);
      };

      reader.readAsDataURL(files[0]);
    } else {
      setCompanyImageUrl("/images/account/user.png");
    }

    setImagesToSave(files);
  };

  return (
    <Container maxW="4xl" py={8}>
      <SimpleCard>
        <WarningModal
          onSubmit={onEmailSubmit}
          title="Actualizar Correo"
          description="¿Estás seguro que quieres actualizar tu correo? Tendrás que volver a iniciar sesión una vez hagas el cambio"
          confirmText="Aceptar"
        />
        <div className="p-4">
          <Flex align="center" gap={4}>
            <Avatar
              size="xl"
              src={
                companyImageUrl
              }
            />
            <Box>
              <Heading size="lg" className="break-all">
                {companyFormData.name}
              </Heading>
              <span className="text-gray-500">
                <a
                  href={`https://www.${companyFormData.nameFormatUrl}.com`}
                  target="_blank"
                >
                  www.{companyFormData.nameFormatUrl}.com
                </a>
              </span>
            </Box>
          </Flex>
          <Tabs isFitted variant="enclosed" className="mt-8">
            <TabList mb="1em">
              <Tab>Empresa</Tab>
              <Tab>Cuenta</Tab>
            </TabList>
            <TabPanels>
              <TabPanel>
                <VStack spacing={4} align="stretch">
                  <Heading size="md" mb={2}>
                    Información de la Empresa
                  </Heading>
                  <Text size="sm" color="gray.500">
                    Actualiza los datos de tu empresa aquí
                  </Text>
                  <FormControl>
                    <label className="text-sm">Nombre</label>
                    <Input
                      size="sm"
                      name="name"
                      maxLength={50}
                      value={companyFormData.name}
                      onChange={handleCompanyFormChange}
                      disabled={!isAdmin}
                    />
                  </FormControl>
                  <FormControl>
                    <label className="text-sm">Imagen</label>
                    <div className="border rounded py-5 px-3">
                      <ImagesUpload
                        maxImagesNumber={1}
                        isViewOnlyImage={!isAdmin}
                        onUploadFiles={handleUploadFiles}
                        defaultImages={imagesLoaded}
                      />
                    </div>
                  </FormControl>
                  {isAdmin && (
                    <Button
                      variant="main"
                      alignSelf="flex-start"
                      className="mt-4"
                      onClick={onCompanySubmit}
                    >
                      Guardar Cambios
                    </Button>
                  )}
                </VStack>
              </TabPanel>
              <TabPanel>
                <VStack spacing={4} align="stretch">
                  <Heading size="md" mb={2}>
                    Ajuste de Cuenta
                  </Heading>
                  <Text size="sm" color="gray.500">
                    Actualiza el correo de tu usuario
                  </Text>
                  <FormControl>
                    <label className="text-sm">Correo</label>
                    <Input
                      size="sm"
                      name="email"
                      value={emailFormData.email}
                      onChange={handleEmailFormChange}
                      max={100}
                    />
                  </FormControl>
                  <Button
                    variant="main"
                    alignSelf="flex-start"
                    className="mt-4"
                    onClick={() => {
                      confirmChangeModal.onOpen();
                    }}
                  >
                    Actualizar Correo
                  </Button>
                  <hr className="my-4" />
                  <Text size="sm" color="gray.500">
                    Cambiar tu contraseña
                  </Text>
                  <FormControl>
                    <label className="text-sm">Contraseña Nueva</label>
                    <Input
                      size="sm"
                      type="password"
                      name="password"
                      maxLength={20}
                      value={passwordFormData.password}
                      onChange={handlePasswordFormChange}
                      onKeyDown={handleSpaceKeyDown}
                    />
                  </FormControl>
                  <FormControl>
                    <label className="text-sm">Repetir Contraseña Nueva</label>
                    <Input
                      size="sm"
                      type="password"
                      name="passwordRepeat"
                      maxLength={20}
                      value={passwordFormData.passwordRepeat}
                      onChange={handlePasswordFormChange}
                      onKeyDown={handleSpaceKeyDown}
                    />
                  </FormControl>
                  <Button
                    variant="main"
                    alignSelf="flex-start"
                    className="mt-4"
                    onClick={onPasswordSubmit}
                  >
                    Cambiar Contraseña
                  </Button>
                </VStack>
              </TabPanel>
            </TabPanels>
          </Tabs>
        </div>
      </SimpleCard>
    </Container>
  );
};

export default AccountClient;
