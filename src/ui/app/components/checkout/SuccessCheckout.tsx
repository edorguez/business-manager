"use client";

import useWindowSize from "@/app/hooks/useWindowSize";
import { Icon } from "@iconify/react";
import { useEffect } from "react";
import Confetti from "react-confetti";

const SuccessCheckout = () => {
  const { width, height, setWindowSize } = useWindowSize();

  useEffect(() => {
    const handleResize = () => {
      setWindowSize(window.innerWidth, window.innerHeight);
    };

    window.addEventListener("resize", handleResize);
    handleResize();

    return () => window.removeEventListener("resize", handleResize);
  }, [setWindowSize]);

  return (
    <>
        <Confetti width={width} height={height} numberOfPieces={150} style={{ position: 'fixed', top: 0, left: 0 }} />
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
