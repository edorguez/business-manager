'use client'

import { PASSWORD } from "@/app/constants";
import { SignUp } from "@/app/types/auth";
import { validEmail, validLettersAndNumbers, validWithNoSpaces } from "@/app/utils/InputUtils";
import { Button, Input, useToast } from "@chakra-ui/react";
import { Icon } from "@iconify/react";

interface SignUpStep2Props {
  userForm: SignUp['user'],
  onUserChange: (user: SignUp['user']) => void;
  onClickBackStep: () => void
  onClickNextStep: () => void
}

const SignUpStep2: React.FC<SignUpStep2Props> = ({
  userForm,
  onUserChange,
  onClickBackStep,
  onClickNextStep
}) => {
  const toast = useToast();

  const handleEmailChange = (event: any) => {
    const { value } = event.target;
    if (value && !validWithNoSpaces(value)) return;
    onUserChange({
      ...userForm,
      email: value
    });
  };

  const handlePasswordChange = (event: any) => {
    const { name, value } = event.target;
    if (value && !validLettersAndNumbers(value)) return;
    onUserChange({
      ...userForm,
      [name]: value
    });
  };

  const onFinish = () => {
    if(isFormValid()) {
      onClickNextStep();
    } else if(userForm.password && userForm.password.length < PASSWORD.MIN_PASSWORD_LEGTH) {
      showErrorMessage(`La contraseña debe tener al menos ${PASSWORD.MIN_PASSWORD_LEGTH} caracteres`);
    } else if (userForm.password !== userForm.passwordRepeat) {
      showErrorMessage("Las contraseñas deben ser iguales");
    } else {
      showErrorMessage("Algunos campos son requeridos o inválidos");
    }
  };

  const isFormValid = (): boolean => {
    if (!userForm.email) return false;
    if (!validEmail(userForm.email)) return false;
    if (!userForm.password) return false;
    if (userForm.password.length < PASSWORD.MIN_PASSWORD_LEGTH) return false;
    if (!validLettersAndNumbers(userForm.password)) return false;
    if (!userForm.passwordRepeat) return false;

    return true;
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
    <>
      <div className="mt-2">
        <label className="text-sm">
          Correo<span className="text-thirdcolor">*</span>
        </label>
        <Input
          size="sm"
          name="email"
          maxLength={100}
          value={userForm.email}
          onChange={handleEmailChange}
        />
      </div>
      <div className="mt-2">
        <label className="text-sm">
          Contraseña<span className="text-thirdcolor">*</span>
        </label>
        <Input
          size="sm"
          type="password"
          name="password"
          maxLength={20}
          value={userForm.password}
          onChange={handlePasswordChange}
        />
      </div>
      <div className="mt-2">
        <label className="text-sm">
          Repetir Contraseña<span className="text-thirdcolor">*</span>
        </label>
        <Input
          size="sm"
          type="password"
          name="passwordRepeat"
          maxLength={20}
          value={userForm.passwordRepeat}
          onChange={handlePasswordChange}
        />
      </div>
      <div className="mt-3 flex justify-between">
        <Button variant="main" className="w-40" onClick={() => onClickBackStep()}>
          <Icon className="mr-2" icon="fa-solid:arrow-left" />
          Atrás
        </Button>

        <Button variant="main" className="w-40" onClick={onFinish}>
          Finalizar
        </Button>
      </div>
    </>
  )
}

export default SignUpStep2;
