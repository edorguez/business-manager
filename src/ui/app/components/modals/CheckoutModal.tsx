"use client";

import useCheckoutModal from "@/app/hooks/useCheckoutModal";
import {
  Box,
  Modal,
  ModalBody,
  ModalContent,
  ModalOverlay,
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
} from "@chakra-ui/react";
import CreateCustomerComponent from "../customers/CreateCustomer";

const CheckoutModal = () => {
  const checkoutModal = useCheckoutModal();

  const { activeStep } = useSteps({
    index: 1,
    count: 2,
  });

  return (
    <>
      <Modal isOpen={checkoutModal.isOpen} onClose={checkoutModal.onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalBody className="my-3">
            <Stepper size="sm" index={activeStep}>
              <Step key={0}>
                <StepIndicator>
                  <StepStatus
                    complete={<StepIcon />}
                    incomplete={<StepNumber />}
                    active={<StepNumber />}
                  />
                </StepIndicator>

                <Box flexShrink="0">
                  <StepTitle>Información</StepTitle>
                  <StepDescription>Datos básicos</StepDescription>
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

                <Box flexShrink="0">
                  <StepTitle>Completado</StepTitle>
                  <StepDescription>Pedido realizado</StepDescription>
                </Box>

                <StepSeparator />
              </Step>
            </Stepper>
            {
                activeStep == 0 &&
                <div className="mt-7">
                    <CreateCustomerComponent />
                </div>
            }
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
};

export default CheckoutModal;
