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
import CreateCustomerComponent from "../orders/CreateOrderCustomer";
import SuccessCheckout from "../checkout/SuccessCheckout";
import { useCallback, useState } from "react";
import {
  CreateOrder,
  CreateOrderCustomer,
  CreateOrderProduct,
} from "@/app/types/order";
import useProductsCart, { CartItem } from "@/app/hooks/useProductsCart";
import useCompanyInfo from "@/app/hooks/useCompanyInfo";
import { CreateOrderRequest } from "@/app/services/orders";

const CheckoutModal = () => {
  const [allowCloseModal, setAllowCloseModal] = useState<boolean>(true);
  const checkoutModal = useCheckoutModal();
  const cart = useProductsCart();
  const companyInfo = useCompanyInfo();

  const { activeStep } = useSteps({
    index: 0,
    count: 2,
  });

  const createOrder = useCallback(async (request: CreateOrder) => {
    const res = await CreateOrderRequest(request);
    console.log("res");
    console.log(res);
  }, []);

  const handleStartCreateOrder = async (customer: CreateOrderCustomer) => {
    setAllowCloseModal(false);
    console.log("hola");
    console.log(customer);
    console.log(cart.items);
    console.log(companyInfo.company);

    const request: CreateOrder = {
      companyId: companyInfo.company?.id ?? 0,
      customer: customer,
      products: cart.items.reduce(
        (res: CreateOrderProduct[], item: CartItem) => {
          res.push({
            productId: item.id,
            quantity: item.quantity,
            price: item.price,
          });

          return res;
        },
        []
      ),
    };

    await createOrder(request);
  };

  return (
    <>
      <Modal
        closeOnOverlayClick={allowCloseModal}
        isOpen={checkoutModal.isOpen}
        onClose={checkoutModal.onClose}
      >
        <ModalOverlay />
        <ModalContent>
          <ModalBody className="my-3">
            <Stepper size="sm" colorScheme="main" index={activeStep}>
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
            {activeStep == 0 && (
              <div className="mt-7">
                <CreateCustomerComponent
                  onStartCreateOrderCustomer={handleStartCreateOrder}
                />
              </div>
            )}
            {activeStep == 1 && (
              <div className="mt-7">
                <SuccessCheckout />
              </div>
            )}
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
};

export default CheckoutModal;
