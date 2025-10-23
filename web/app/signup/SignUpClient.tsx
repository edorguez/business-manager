"use client";

import Image from "next/image";
import SimpleCard from "../components/cards/SimpleCard";
import {
  Box,
  Button,
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
import { Icon } from "@iconify/react";
import { SignUp } from "../types/signup";
import SignUpStep1 from "../components/signup/SignUpStep1";

const SignUpClient = () => {
  const isLoading = useLoading();
  const toast = useToast();
  const { push } = useRouter();
  const [formData, setFormData] = useState<SignUp>({ company: { name: "", phone: ""}, user: { email: "", password: "" } });

  const updateCompany = (updatedCompany: SignUp['company']) => {
    setFormData(prev => ({
      ...prev,
      company: updatedCompany
    }));
  };

  const { activeStep, setActiveStep } = useSteps({
    index: 0,
    count: 2,
  });

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
                      <SignUpStep1 companyForm={formData.company} onCompanyChange={updateCompany} onClickNextStep={() => console.log('hola')} />
                    </div>
                  )}
                  {activeStep === 1 && (
                    <div className="mt-7">
                      Chao
                      <div className="mt-3">
                        <Button variant="main" className="w-40">
                          Siguiente
                          <Icon className="ml-2" icon="fa-solid:arrow-right" />
                        </Button>
                      </div>
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
