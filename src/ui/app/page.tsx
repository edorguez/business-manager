"use client";

import { Button } from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import Image from "next/image";
import { useEffect } from "react";
import AOS from "aos";
import "aos/dist/aos.css";
import FloatingItem from "./components/framer-motion/FloatingItem";

export default function Home() {
  useEffect(() => {
    AOS.init({
      duration: 1000,
    });
  }, []);

  return (
    <div className="flex flex-col min-h-screen bg-white">
      <header className="px-4 lg:px-6 h-16 flex items-center bg-maincolor">
        <div className="flex items-center justify-center">
          <Image src="/images/logo.png" alt="Logo" width={28} height={28} />
          <span className="ml-2 text-2xl font-bold text-white">Edezco</span>
        </div>
        <nav className="ml-auto sm:flex gap-4 sm:gap-6 hidden">
          <a
            className="text-sm font-medium hover:underline underline-offset-4 text-white"
            href="#features"
          >
            Características
          </a>
          <a
            className="text-sm font-medium hover:underline underline-offset-4 text-white"
            href="#plans"
          >
            Planes
          </a>
          <a
            className="text-sm font-medium hover:underline underline-offset-4 text-white"
            href="#contact"
          >
            Contáctanos
          </a>
        </nav>
      </header>
      <main className="flex-1">
        <section className="relative w-full py-48 xl:py-64 bg-gradient-to-b from-maincolor to-white overflow-hidden mb-20">
          {/* Image Container with Absolute Positioning */}
          <div className="absolute inset-0 pointer-events-none">
            <FloatingItem className="absolute top-10 left-1/4 animate-float">
              <Image
                src="/images/home/bag.png"
                alt="Bag"
                width={64}
                height={64}
              />
            </FloatingItem>
            <FloatingItem className="absolute top-5 right-80 animate-float">
              <Image
                src="/images/home/box.png"
                alt="Box"
                width={64}
                height={64}
              />
            </FloatingItem>
            <FloatingItem className="absolute bottom-10 left-1/4 animate-float">
              <Image
                src="/images/home/graph.png"
                alt="Graph"
                width={64}
                height={64}
              />
            </FloatingItem>
            <FloatingItem className="absolute top-1/2 left-80 animate-float">
              <Image
                src="/images/home/order-processing.png"
                alt="Order"
                width={64}
                height={64}
              />
            </FloatingItem>
            <FloatingItem className="absolute bottom-20 right-1/4 animate-float">
              <Image
                src="/images/home/smartphone.png"
                alt="Smartphone"
                width={64}
                height={64}
              />
            </FloatingItem>
            <FloatingItem className="absolute top-40 right-60 animate-float">
              <Image
                src="/images/home/whatsapp.png"
                alt="WhatsApp"
                width={64}
                height={64}
              />
            </FloatingItem>
          </div>

          {/* Main Content Section */}
          <div className="container relative z-10 mx-auto px-4 md:px-6">
            <div className="flex flex-col items-center space-y-4 text-center">
              <div className="space-y-2">
                <div className="rounded-full flex justify-center">
                  <Image
                    className="rounded-full border-rose-500 "
                    src="/images/logo.png"
                    width={100}
                    height={100}
                    alt="logo"
                  />
                </div>
                <h1 className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl lg:text-6xl/none relative">
                  Automatiza tus ventas
                </h1>
                <p className="mx-auto max-w-[700px] text-gray-700 md:text-xl">
                  Crea tu tienda online y lleva control de todos tus productos y
                  ventas en cualquier lugar
                </p>
              </div>
              <div className="space-x-4 transition-transform transform hover:scale-125 duration-300 ease-in-out">
                <Button variant="third">
                  <a href="https://wa.link/bpguv0" target="_blank">
                    Comenzar
                  </a>
                </Button>
              </div>
            </div>
          </div>
        </section>
        <section className="container mx-auto px-4 py-8 sm:py-12 md:py-16">
          <div className="flex flex-col lg:flex-row items-center justify-between gap-8">
            <div data-aos="fade-up" className="w-full lg:w-1/2">
              <div className="relative aspect-video">
                <Image
                  src="/images/home/page-pc.png"
                  alt="Product on Laptop"
                  layout="fill"
                  objectFit="contain"
                  className="rounded-lg object-cover"
                />
              </div>
            </div>
            <div
              className="w-full lg:w-1/2 space-y-4 text-center"
              data-aos="fade-up"
            >
              <h2 className="text-2xl sm:text-3xl md:text-4xl font-bold text-gray-800">
                Tu tienda en segundos
              </h2>
              <p className="text-base sm:text-lg text-gray-600">
                Vende tus productos a través de tu propia tienda personalizada a
                tu medida
              </p>
            </div>
          </div>
        </section>
        <section className="w-full bg-maincolor">
          <div className="container mx-auto px-4 py-8 sm:py-12 md:py-16">
            <div className="flex flex-col lg:flex-row items-center justify-between gap-8">
              <div
                className="w-full lg:w-1/2 space-y-4 text-center"
                data-aos="fade-up"
              >
                <h2 className="text-2xl sm:text-3xl md:text-4xl font-bold text-white">
                  Notificaciones
                </h2>
                <p className="text-base sm:text-lg text-white">
                  Obtén notificaciones a través de WhatsApp con información
                  detallada sobre tus pedidos
                </p>
              </div>
              <div data-aos="fade-up" className="w-full lg:w-1/2">
                <div className="relative aspect-video">
                  <Image
                    src="/images/home/page-phone.png"
                    alt="Product on Phone"
                    layout="fill"
                    objectFit="contain"
                    className="rounded-lg object-cover"
                  />
                </div>
              </div>
            </div>
          </div>
        </section>
        <section className="w-full bg-white">
          <div className="container mx-auto px-4 py-8 sm:py-12 md:py-16">
            <div className="flex flex-col lg:flex-row items-center justify-between gap-8">
              <div data-aos="fade-up" className="w-full lg:w-1/2">
                <div className="relative aspect-video">
                  <Image
                    src="/images/home/page-laptop.png"
                    alt="Product on Laptop"
                    layout="fill"
                    objectFit="contain"
                    className="rounded-lg object-cover"
                  />
                </div>
              </div>
              <div
                className="w-full lg:w-1/2 space-y-4 text-center"
                data-aos="fade-up"
              >
                <h2 className="text-2xl sm:text-3xl md:text-4xl font-bold text-gray-800">
                  Analíticas
                </h2>
                <p className="text-base sm:text-lg text-gray-600">
                  Obtén información sobre tus productos, clientes y métodos de
                  pago
                </p>
              </div>
            </div>
          </div>
        </section>
        <section className="w-full bg-thirdcolor">
          <div className="container mx-auto px-4 py-8 sm:py-12 md:py-16">
            <div className="flex flex-col lg:flex-row items-center justify-between gap-8">
              <div
                className="w-full lg:w-1/2 space-y-4 text-center"
                data-aos="fade-up"
              >
                <h2 className="text-2xl sm:text-3xl md:text-4xl font-bold text-white">
                  Administra
                </h2>
                <p className="text-base sm:text-lg text-white">
                  Lleva control detallado de tus productos y clientes
                </p>
              </div>
              <div data-aos="fade-up" className="w-full lg:w-1/2">
                <div className="relative aspect-video">
                  <Image
                    src="/images/home/page-tablet.png"
                    alt="Product on Laptop"
                    layout="fill"
                    objectFit="contain"
                    className="rounded-lg object-cover"
                  />
                </div>
              </div>
            </div>
          </div>
        </section>
        <section
          className="w-full py-12 md:py-24 lg:py-32 bg-gray-100"
          id="features"
        >
          <div className="container mx-auto px-4 md:px-6">
            <h2
              className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl text-center mb-14"
              data-aos="fade-up"
            >
              Características
            </h2>
            <div
              className="grid gap-10 sm:grid-cols-2 md:grid-cols-3"
              data-aos="fade-up"
            >
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
        <section className="w-full py-12 md:py-24 lg:py-32 bg-white" id="plans">
          <div className="container mx-auto px-4 md:px-6">
            <h2
              className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl text-center mb-14"
              data-aos="fade-up"
            >
              Planes
            </h2>
            <div className="grid gap-6 sm:grid-cols-3 lg:px-52">
              <div
                className="flex flex-col p-6 bg-white rounded-lg shadow-lg border border-maincolor"
                data-aos="fade-up"
                data-aos-delay="400"
              >
                <h3 className="text-2xl font-bold text-maincolor mb-4">
                  Gratis
                </h3>
                <p className="text-4xl font-bold mb-4">
                  $0<span className="text-xl text-gray-500">/mensual</span>
                </p>
                <ul className="mb-6 space-y-2">
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Todo el plan
                    básico
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> 7 días de
                    período de prueba
                  </li>
                </ul>
              </div>
              <div
                className="flex flex-col p-6 bg-white rounded-lg shadow-lg border border-thirdcolor"
                data-aos="fade-up"
                data-aos-delay="400"
              >
                <h3 className="text-2xl font-bold text-thirdcolor mb-4">
                  Básico
                </h3>
                <p className="text-4xl font-bold mb-4 text-black">
                  $15<span className="text-xl text-black">/mensual</span>
                </p>
                <ul className="mb-6 space-y-2 text-black">
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Tienda online
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Pedidos por
                    WhatsApp
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Analíticas
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Registro de
                    clientes
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Registro de
                    productos (limitado)
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Registro de
                    métodos de pago (limitado)
                  </li>
                </ul>
              </div>
              <div
                className="flex flex-col p-6 bg-maincolor rounded-lg shadow-lg"
                data-aos="fade-up"
                data-aos-delay="400"
              >
                <h3 className="text-2xl font-bold text-white mb-4">Pro</h3>
                <p className="text-4xl font-bold text-white mb-4">
                  $25<span className="text-xl text-gray-200">/mensual</span>
                </p>
                <ul className="mb-6 space-y-2 text-white">
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Tienda online
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Pedidos por
                    WhatsApp
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Analíticas
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Registro de
                    clientes
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Registro de
                    productos
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Registro de
                    métodos de pago
                  </li>
                  <li className="flex items-center">
                    <Icon icon="material-symbols-light:check" /> Creación de
                    multiples usuarios
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </section>
        <section className="w-full py-12 md:py-24 lg:py-32 bg-thirdcolor">
          <div className="container mx-auto px-4 md:px-6">
            <div className="flex flex-col items-center justify-center space-y-4 text-center">
              <div className="space-y-2">
                <h2
                  className="text-3xl font-bold tracking-tighter sm:text-4xl md:text-5xl text-white"
                  data-aos="fade-up"
                >
                  ¿Preparado para transformar tu negocio?
                </h2>
                <p
                  className="mx-auto max-w-[600px] text-gray-100 md:text-xl"
                  data-aos="fade-up"
                  data-aos-delay="400"
                >
                  Únete a nuestros clientes satisfechos y toma el primer paso
                  para el éxito
                </p>
              </div>
              <div
                className="w-full max-w-sm space-y-2"
                data-aos="fade-up"
                data-aos-delay="400"
              >
                <div className="space-x-4 transition-transform transform hover:scale-125 duration-300 ease-in-out">
                  <Button variant="main">
                    <a href="https://wa.link/bpguv0" target="_blank">
                      Comenzar
                    </a>
                  </Button>
                </div>
              </div>
            </div>
          </div>
        </section>
        <section className="w-full bg-white" id="contact">
          <div className="container mx-auto px-4 py-8 sm:py-12 md:py-16">
            <div className="flex flex-col lg:flex-row items-center justify-between gap-8">
              <div
                className="w-full lg:w-1/2 space-y-4 text-center"
                data-aos="fade-up"
              >
                <h2 className="text-2xl sm:text-3xl md:text-4xl font-bold text-black">
                  Contáctanos
                </h2>
                <div
                  className="grid sm:grid-cols-2"
                  data-aos="fade-up"
                >
                  <div className="flex flex-col items-center text-center">
                    <Icon
                      icon="material-symbols:mail-outline"
                      className="text-maincolor text-6xl"
                    />
                    <h3 className="text-xl font-bold mt-4 mb-2">
                      Correo
                    </h3>
                    <p className="text-gray-600">
                      info@edezco.com
                    </p>
                  </div>
                  <div className="flex flex-col items-center text-center">
                    <Icon
                      icon="lineicons:phone"
                      className="text-maincolor text-6xl"
                    />
                    <h3 className="text-xl font-bold mt-4 mb-2">
                      Teléfono
                    </h3>
                    <p className="text-gray-600">
                      +58 0412-0238498
                    </p>
                  </div>
                </div>
              </div>
              <div data-aos="fade-up" className="w-full lg:w-1/2">
                <div className="relative aspect-video">
                  <Image
                    src="/images/home/contact.png"
                    alt="Contact"
                    layout="fill"
                    objectFit="contain"
                    className="rounded-lg object-cover"
                  />
                </div>
              </div>
            </div>
          </div>
        </section>
      </main>
      <footer className="flex flex-col gap-2 sm:flex-row py-6 w-full shrink-0 items-center px-4 md:px-6 border-t">
        <p className="text-xs text-gray-500">
          © {new Date().getFullYear()} Edezco
        </p>
      </footer>
    </div>
  );
}
