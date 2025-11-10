'use client'

import { validLettersAndNumbers, validNumbers, validPhone, validSubdomain } from "@/app/utils/InputUtils";
import { Button, Input, InputGroup, InputRightAddon, useToast } from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import ImagesUpload from "../uploads/ImagesUpload";
import { SignUp } from "@/app/types/auth";

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
  const toast = useToast();

  const handleNameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    if (value && !validLettersAndNumbers(value, true)) return;

    let urlValue: string = value.trimStart().replaceAll(/ /g, '-');

    onCompanyChange({
      ...companyForm,
      [name]: value.trimStart(),
      nameFormatUrl: urlValue
    });
  };

  const handleNameFormatUrlChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = event.target;
    if (value && !validSubdomain(value)) return;
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
      images: files
    })
  };

  const onNextStep = () => {
    if(isFormValid()) {
      onClickNextStep();
    } else {
      showErrorMessage("Algunos campos son requeridos o inválidos");
    }
  };

  const isFormValid = (): boolean => {
    if (!companyForm.name) return false;
    if (companyForm.name && !validLettersAndNumbers(companyForm.name, true)) return false;
    if (!companyForm.nameFormatUrl) return false;
    if (!companyForm.phone) return false;
    if (companyForm.phone && !validPhone(companyForm.phone)) return false;

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
          Nombre Empresa<span className="text-thirdcolor">*</span>
        </label>
        <Input
          size="sm"
          name="name"
          maxLength={50}
          value={companyForm.name}
          onChange={handleNameChange}
        />
      </div>
      <div className="mt-2">
        <label className="text-sm">
          URL Página Web<span className="text-thirdcolor">*</span>
        </label>
        
        <InputGroup size='sm'>
          <Input
          size="sm"
          name="nameFormatUrl"
          maxLength={50}
          value={companyForm.nameFormatUrl}
          onChange={handleNameFormatUrlChange}
          />
          <InputRightAddon>.edezco.com</InputRightAddon>
        </InputGroup>
      </div>
      <div className="mt-2">
        <label className="text-sm">
          Número de WhatsApp<span className="text-thirdcolor">*</span>
        </label>
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
        <Button variant="main" className="w-40" onClick={onNextStep}>
          Siguiente
          <Icon className="ml-2" icon="fa-solid:arrow-right" />
        </Button>
      </div>
    </>
  )
}

export default SignUpStep1;
