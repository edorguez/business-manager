import { NextPage } from "next";
import Image from "next/image";

const NotFound: NextPage = () => {
  return (
    <div className="mt-32 text-center">
      <div className="flex justify-center">
        <Image
          src="/images/general/warning-person.png"
          width={512}
          height={512}
          alt="Not Found"
        />
      </div>
      <h1 className="font-bold">404 - Página NO encontrada</h1>
      <p className="my-4">
        Oops! La página que estabas buscando no fue encontrada.
      </p>
      <a href="/" className="underline decoration-1 text-maincolor">
        Regresar al inicio
      </a>
    </div>
  );
};

export default NotFound;
