"use client";

import Image from "next/image";
import SimpleCard from "../components/cards/SimpleCard";
import { Button, Container, Input, useToast } from "@chakra-ui/react";
import { useState } from "react";
import { Login } from "../types/auth";
import { login } from "../services/auth";
import { useRouter } from "next/navigation";
import useLoading from "../hooks/useLoading";
import { validLettersAndNumbers, validWithNoSpaces } from "../utils/InputUtils";

const LoginClient = () => {
  const isLoading = useLoading();
  const toast = useToast();
  const { push } = useRouter();
  const [formData, setFormData] = useState<Login>({ email: "", password: "" });

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

  const onLogin = async () => {
    isLoading.onStartLoading();
    let result: any = await login(formData);
    if (!result?.error) {
      push("/management/home");
    } else {
      toast({
        title: "Error",
        description: result.error,
        variant: "customerror",
        position: "top-right",
        duration: 4000,
        isClosable: true,
      });
    }
  };

  const handleKeyDown = (event: any) => {
    if (event.key === "Enter") {
      onLogin();
    }
  };

  return (
    <>
      <div className="grid grid-cols-8 h-screen">
        <div className="hidden md:flex col-span-3 bg-gradient-to-b from-thirdcolor to-fourthcolor flex-col items-center justify-center">
          <h1 className="mb-10 text-white font-medium text-xl text-center">
            ¡Ingresa a tu cuenta y empieza a administrar tu negocio!
          </h1>
          <Image
            className="rounded-lg"
            src="/images/login/main_image.png"
            alt="Logo"
            width={300}
            height={300}
          />
        </div>
        <div className="col-span-8 md:col-span-5 bg-graybackground flex items-center">
          <Container>
            <div className="relative">
              <div className="absolute rounded-full border-[7px] border-white top-[-60px] left-[40%]">
                <Image
                  className="rounded-full border-rose-500 "
                  src="/images/logo.png"
                  width={100}
                  height={100}
                  alt="logo"
                />
              </div>
              <SimpleCard>
                <div className="px-1 py-8">
                  <div className="mt-2">
                    <label className="text-sm">Correo</label>
                    <Input
                      size="sm"
                      type="email"
                      name="email"
                      maxLength={100}
                      value={formData.email}
                      onChange={handleEmailChange}
                      onKeyDown={handleKeyDown}
                    />
                  </div>
                  <div className="mt-2">
                    <label className="text-sm">Contraseña</label>
                    <Input
                      size="sm"
                      type="password"
                      name="password"
                      maxLength={20}
                      value={formData.password}
                      onChange={handlePasswordChange}
                      onKeyDown={handleKeyDown}
                    />
                  </div>
                  <div className="mt-3">
                    <Button variant="main" className="w-full" onClick={onLogin}>
                      Iniciar Sesión
                    </Button>
                  </div>
                </div>
              </SimpleCard>
            </div>
          </Container>
        </div>
      </div>
    </>
  );
};

export default LoginClient;
