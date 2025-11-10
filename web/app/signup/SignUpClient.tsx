"use client";

import Image from "next/image";
import SimpleCard from "../components/cards/SimpleCard";
import {
  Box,
  Container,
  Step,
  StepDescription,
  StepIcon,
  StepIndicator,
  StepNumber,
  Stepper,
  StepSeparator,
  StepStatus,
  StepTitle,
  useSteps,
  useToast
} from "@chakra-ui/react";
import { useState } from "react";
import { useRouter } from "next/navigation";
import useLoading from "../hooks/useLoading";
import SignUpStep1 from "../components/signup/SignUpStep1";
import SignUpStep2 from "../components/signup/SignUpStep2";
import { SignUp } from "../types/auth";
import { signUp } from "../services/auth";

const SignUpClient = () => {
  const isLoading = useLoading();
  const toast = useToast();
  const { push } = useRouter();
  const [formData, setFormData] = useState<SignUp>({ 
    company: { 
      name: "", 
      nameFormatUrl: "",
      phone: "", 
      images: undefined 
    }, 
    user: { 
      email: "", 
      password: "", 
      passwordRepeat: "" 
    } 
  });

  const updateCompany = (updatedCompany: SignUp['company']) => {
    setFormData((prev: any) => ({
      ...prev,
      company: updatedCompany
    }));
  };

  const updateUser = (updatedUser: SignUp['user']) => {
    setFormData((prev: any) => ({
      ...prev,
      user: updatedUser
    }));
  };

  const { activeStep, setActiveStep } = useSteps({
    index: 0,
    count: 2,
  });

  const handleSignup = async () => {
    isLoading.onStartLoading();
    let result: any = await signUp(formData);
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
      isLoading.onEndLoading();
    }
  }

  return (
    <>
      <div className="grid grid-cols-8 h-screen">
        <div className="hidden md:flex col-span-3 bg-gradient-to-b from-thirdcolor to-fourthcolor flex-col items-center justify-center">
          <h1 className="mb-10 text-white font-medium text-xl text-center">
            ¡Únete y empieza a administrar tu negocio!
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
                  <Stepper size="sm" colorScheme="main" index={activeStep}>
                    <Step key={0}>
                      <StepIndicator>
                        <StepStatus
                          complete={<StepIcon />}
                          incomplete={<StepNumber />}
                          active={<StepNumber />}
                        />
                      </StepIndicator>

                      <Box flexShrink="0" className="select-none">
                        <StepTitle>Empresa</StepTitle>
                        <StepDescription>Datos empresa</StepDescription>
                      </Box>

                      <StepSeparator />
                    </Step>
                    <Step key={1}>
                      <StepIndicator>
                        <StepStatus
                          complete={<StepIcon />}
                          incomplete={<StepNumber />}
                          active={<StepNumber />}
                        />
                      </StepIndicator>

                      <Box flexShrink="0" className="select-none">
                        <StepTitle>Usuario</StepTitle>
                        <StepDescription>Datos usuario</StepDescription>
                      </Box>

                      <StepSeparator />
                    </Step>
                  </Stepper>
                  {activeStep === 0 && (
                    <div className="mt-7">
                      <SignUpStep1 
                        companyForm={formData.company} 
                        onCompanyChange={updateCompany} 
                        onClickNextStep={() => setActiveStep(1)} 
                      />
                    </div>
                  )}
                  {activeStep === 1 && (
                    <div className="mt-7">
                      <SignUpStep2
                        userForm={formData.user} 
                        onUserChange={updateUser} 
                        onClickBackStep={() => setActiveStep(0)} 
                        onClickNextStep={handleSignup} 
                      />
                    </div>                  
                  )}
                </div>
              </SimpleCard>
            </div>
          </Container>
        </div>
      </div>
    </>
  );
};

export default SignUpClient;
