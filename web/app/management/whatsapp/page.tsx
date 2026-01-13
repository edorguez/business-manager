'use client';

import Container from "@/app/components/Container";
import WhatsAppClient from "./WhatsAppClient";
import { redirect } from "next/navigation";

const WhatsAppPage = () => {
  redirect('/management/home');

  return null;
  // return (
  //   <Container>
  //     <WhatsAppClient />
  //   </Container>
  // )
}

export default WhatsAppPage;
