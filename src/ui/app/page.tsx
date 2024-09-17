"use client";

import { Button, Input } from "@chakra-ui/react";
import { Icon } from "@iconify/react";

export default function Home() {
  return (
    <div className="flex flex-col min-h-screen bg-white">
      <header className="px-4 lg:px-6 h-16 flex items-center bg-maincolor">
        <a className="flex items-center justify-center" href="#">
          <span className="sr-only">Your SaaS Company</span>
          <svg
            className="h-6 w-6 text-white"
            fill="none"
            height="24"
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth="2"
            viewBox="0 0 24 24"
            width="24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path d="M6 2 3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4Z" />
            <path d="M3 6h18" />
            <path d="m16 10-4 4-4-4" />
          </svg>
          <span className="ml-2 text-2xl font-bold text-white">SaaSy</span>
        </a>
        <nav className="ml-auto flex gap-4 sm:gap-6">
          <a
            className="text-sm font-medium hover:underline underline-offset-4 text-white"
            href="#"
          >
            Características
          </a>
          <a
            className="text-sm font-medium hover:underline underline-offset-4 text-white"
            href="#"
          >
            Planes
          </a>
          <a
            className="text-sm font-medium hover:underline underline-offset-4 text-white"
            href="#"
          >
            Contáctanos
          </a>
        </nav>
      </header>
      <main className="flex-1">
        <section className="w-full py-12 md:py-24 lg:py-32 xl:py-48 bg-gradient-to-b from-maincolor to-white">
          <div className="container mx-auto px-4 md:px-6">
            <div className="flex flex-col items-center space-y-4 text-center">
              <div className="space-y-2">
                <h1 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl lg:text-6xl/none">
                  Automatiza tus ventas
                </h1>
                <p className="mx-auto max-w-[700px] text-gray-700 md:text-xl">
                  Crea tu tienda online y lleva control de todos tus productos y
                  ventas
                </p>
              </div>
              <div className="space-x-4">
                <Button variant="third">Comenzar</Button>
              </div>
            </div>
          </div>
        </section>
        <section className="w-full py-12 md:py-24 lg:py-32 bg-gray-100">
          <div className="container mx-auto px-4 md:px-6">
            <h2 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl text-center mb-14">
              Características
            </h2>
            <div className="grid gap-10 sm:grid-cols-2 md:grid-cols-3">
              <div className="flex flex-col items-center text-center">
                <Icon
                  icon="ion:storefront"
                  className="text-thirdcolor text-6xl"
                />
                <h3 className="text-xl font-bold mt-4 mb-2">Tienda online</h3>
                <p className="text-gray-600">
                  Crea tu tienda online donde tus clientes puedan seleccionar y
                  comprar tus productos/servicios
                </p>
              </div>
              <div className="flex flex-col items-center text-center">
                <Icon
                  icon="ic:baseline-whatsapp"
                  className="text-maincolor text-6xl"
                />
                <h3 className="text-xl font-bold mt-4 mb-2">Automatización</h3>
                <p className="text-gray-600">
                  Recibe mensajes automatizados con los pedidos de tus clientes
                </p>
              </div>
              <div className="flex flex-col items-center text-center">
                <Icon
                  icon="dashicons:analytics"
                  className="text-thirdcolor text-6xl"
                />
                <h3 className="text-xl font-bold mt-4 mb-2">Administración</h3>
                <p className="text-gray-600">
                  Gestiona y administra todos tus productos en nuestro sistema
                  con analíticas incluidas
                </p>
              </div>
            </div>
          </div>
        </section>
        <section className="w-full py-12 md:py-24 lg:py-32 bg-white">
          <div className="container mx-auto px-4 md:px-6">
            <h2 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl text-center mb-14">
              Planes
            </h2>
            <div className="grid gap-6 sm:grid-cols-2 lg:px-52">
              <div className="flex flex-col p-6 bg-white rounded-lg shadow-lg border border-maincolor">
                <h3 className="text-2xl font-bold text-maincolor mb-4">
                  Básico
                </h3>
                <p className="text-4xl font-bold mb-4">
                  $19<span className="text-xl text-gray-500">/mo</span>
                </p>
                <ul className="mb-6 space-y-2">
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> 5 Projects
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> 10GB Storage
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Basic Analytics
                  </li>
                </ul>
                <Button className="mt-auto bg-maincolor text-white hover:bg-maincolorhov">
                  Choose Plan
                </Button>
              </div>
              <div className="flex flex-col p-6 bg-maincolor rounded-lg shadow-lg">
                <h3 className="text-2xl font-bold text-white mb-4">Pro</h3>
                <p className="text-4xl font-bold text-white mb-4">
                  $49<span className="text-xl text-gray-200">/mo</span>
                </p>
                <ul className="mb-6 space-y-2 text-white">
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Unlimited
                    Projects
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> 100GB Storage
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Advanced
                    Analytics
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Priority
                    Support
                  </li>
                </ul>
                <Button className="mt-auto bg-white text-maincolor hover:bg-gray-100">
                  Choose Plan
                </Button>
              </div>
            </div>
          </div>
        </section>
        <section className="w-full py-12 md:py-24 lg:py-32 bg-thirdcolor">
          <div className="container mx-auto px-4 md:px-6">
            <div className="flex flex-col items-center justify-center space-y-4 text-center">
              <div className="space-y-2">
                <h2 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl text-white">
                  ¿Preparado para transformar tu negocio?
                </h2>
                <p className="mx-auto max-w-[600px] text-gray-100 md:text-xl">
                  Únete a nuestros clientes satisfechos y toma el primer paso
                  para el éxito
                </p>
              </div>
              <div className="w-full max-w-sm space-y-2">
                <Button variant="main">
                  Comenzar
                </Button>
              </div>
            </div>
          </div>
        </section>
      </main>
      <footer className="flex flex-col gap-2 sm:flex-row py-6 w-full shrink-0 items-center px-4 md:px-6 border-t">
        <p className="text-xs text-gray-500">
          © {new Date().getFullYear()} Business Manager.
        </p>
        {/* <nav className="sm:ml-auto flex gap-4 sm:gap-6">
          <a
            className="text-xs hover:underline underline-offset-4 text-gray-500"
            href="#"
          >
            Terms of Service
          </a>
          <a
            className="text-xs hover:underline underline-offset-4 text-gray-500"
            href="#"
          >
            Privacy
          </a>
        </nav> */}
      </footer>
    </div>
  );
}
