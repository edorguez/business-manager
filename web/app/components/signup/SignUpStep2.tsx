'use client'

import { SignUp } from "@/app/types/signup";
import { validLettersAndNumbers, validWithNoSpaces } from "@/app/utils/InputUtils";
import { Button, Input } from "@chakra-ui/react";
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


  return (
    <>
      <div className="mt-2">
        <label className="text-sm">Correo</label>
        <Input
          size="sm"
          name="email"
          type="email"
          maxLength={100}
          value={userForm.email}
          onChange={handleEmailChange}
        />
      </div>
      <div className="mt-2">
        <label className="text-sm">Contraseña</label>
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
        <label className="text-sm">Repetir Contraseña</label>
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

        <Button variant="main" className="w-40" onClick={() => onClickNextStep()}>
          Finalizar
        </Button>
      </div>
    </>
  )
}

export default SignUpStep2;
