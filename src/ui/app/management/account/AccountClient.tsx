"use client";

import React, { useState } from "react";
import {
  Box,
  Button,
  Container,
  Flex,
  FormControl,
  FormLabel,
  Heading,
  Input,
  Stack,
  Switch,
  Tab,
  TabList,
  TabPanel,
  TabPanels,
  Tabs,
  Text,
  VStack,
  Avatar,
  useColorModeValue,
} from "@chakra-ui/react";
import SimpleCard from "@/app/components/cards/SimpleCard";
import ImagesUpload from "@/app/components/uploads/ImagesUpload";

const AccountClient = () => {
  const [name, setName] = useState("Sofia Davis");
  const [email, setEmail] = useState("sofia.davis@example.com");

  return (
    <Container maxW="4xl" py={8}>
      <SimpleCard>
        <div className="p-4">
          <Flex align="center" gap={4}>
            <Avatar
              size="xl"
              name={name}
              src="/placeholder.svg?height=80&width=80"
            />
            <Box>
              <Heading size="lg">
                {name}
              </Heading>
            </Box>
          </Flex>
          <Tabs isFitted variant="enclosed" className="mt-8">
            <TabList mb="1em">
              <Tab>Empresa</Tab>
              <Tab>Cuenta</Tab>
            </TabList>
            <TabPanels>
              <TabPanel>
                <VStack spacing={4} align="stretch">
                  <Heading size="md" mb={2}>
                    Información de la Empresa
                  </Heading>
                  <Text size="sm" color="gray.500">
                    Actualiza los datos de tu empresa aquí
                  </Text>
                  <FormControl>
                    <label className="text-sm">Nombre</label>
                    <Input
                    size="sm"
                      value={name}
                      onChange={(e) => setName(e.target.value)}
                    />
                  </FormControl>
                  <FormControl>
                    <label className="text-sm">Imagen</label>
                    <div className="border rounded py-5 px-3">
                        <ImagesUpload maxImagesNumber={1} showAddImage={true} />
                    </div>
                  </FormControl>
                  <Button variant="main" alignSelf="flex-start" className="mt-4">
                    Guardar Cambios
                  </Button>
                </VStack>
              </TabPanel>
              <TabPanel>
                <VStack spacing={4} align="stretch">
                  <Heading size="md" mb={2}>
                    Ajuste de Cuenta
                  </Heading>
                  <Text size="sm" color="gray.500">Actualiza los datos de tu cuenta</Text>
                  <FormControl>
                    <label className="text-sm">Correo</label>
                    <Input size="sm" value="sofiadavis" />
                  </FormControl>
                  <FormControl>
                    <label className="text-sm">Contraseña</label>
                    <Input size="sm" value="English" />
                  </FormControl>
                  <Button variant="main" alignSelf="flex-start" className="mt-4">
                    Actualizar Cuenta
                  </Button>
                </VStack>
              </TabPanel>
            </TabPanels>
          </Tabs>
        </div>
      </SimpleCard>
    </Container>
  );
};

export default AccountClient;
