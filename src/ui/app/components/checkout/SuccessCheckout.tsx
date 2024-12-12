"use client";

import { Icon } from "@iconify/react";

const SuccessCheckout = () => {
  return (
    <>
      <div className="flex justify-center text-9xl text-maincolor">
        <Icon icon="icon-park-twotone:check-one" />
      </div>
      <div className="mt-3">
        <h1 className="text-center font-bold text-lg">¡Pedido realizado!</h1>
        <p className="text-center text-sm">
          Nos comunicaremos contigo para concretar la orden
        </p>
      </div>
    </>
  );
};

export default SuccessCheckout;
