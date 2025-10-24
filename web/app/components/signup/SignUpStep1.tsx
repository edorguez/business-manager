'use client'

import { SignUp } from "@/app/types/signup";
import { validNumbers } from "@/app/utils/InputUtils";
import { Button, Input } from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import ImagesUpload from "../uploads/ImagesUpload";

interface SignUpStep1Props {
  companyForm: SignUp['company'],
  onCompanyChange: (company: SignUp['company']) => void;
  onClickNextStep: () => void
}

const SignUpStep1: React.FC<SignUpStep1Props> = ({
  companyForm,
  onCompanyChange,
  onClickNextStep
}) => {
  const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    onCompanyChange({
      ...companyForm,
      [name]: value
    });
  };

  const handlePhoneChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    if (value && !validNumbers(value)) return;
    onCompanyChange({
      ...companyForm,
      [name]: value
    });
  };

  const handleUploadFiles = (files: File[]) => {
    onCompanyChange({
      ...companyForm,
      image: files
    })
  };

  return (
    <>
      <div className="mt-2">
        <label className="text-sm">Nombre Empresa</label>
        <Input
          size="sm"
          name="name"
          maxLength={50}
          value={companyForm.name}
          onChange={handleNameChange}
        />
      </div>
      <div className="mt-2">
        <label className="text-sm">Número de WhatsApp</label>
        <Input
          size="sm"
          type="tel"
          name="phone"
          maxLength={11}
          value={companyForm.phone}
          onChange={handlePhoneChange}
        />
      </div>
      <div className="mt-2">
        <label className="text-sm">Imagen</label>
        <div className="border rounded py-5 px-3">
          <ImagesUpload
            maxImagesNumber={1}
            isViewOnlyImage={false}
            onUploadFiles={handleUploadFiles}
          />
        </div>
      </div>
      <div className="mt-3 flex justify-end">
        <Button variant="main" className="w-40" onClick={() => onClickNextStep()}>
          Siguiente
          <Icon className="ml-2" icon="fa-solid:arrow-right" />
        </Button>
      </div>
    </>
  )
}

export default SignUpStep1;
