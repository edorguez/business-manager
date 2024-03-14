"use client";

import useDeleteModal from '@/app/hooks/useDeleteModal';
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalBody,
  Button,
} from '@chakra-ui/react'
import { Icon } from '@iconify/react';

interface DeleteModalProps {
  title: string;
  description: string;
  onSubmit: () => void;
}

const DeleteModal: React.FC<DeleteModalProps> = ({
  title,
  description,
  onSubmit
}) => {
  const deleteModal = useDeleteModal();

  return (
    <>
      <Modal isOpen={deleteModal.isOpen} onClose={deleteModal.onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalBody>
            <div className='flex justify-center mt-5'>
              <div className='rounded-full bg-thirdcolorhov text-thirdcolor text-5xl px-2 pt-1 pb-2'>
                <Icon icon="clarity:warning-solid" />
              </div>
            </div>
            <div className='text-center my-3'>
              <h1 className='font-bold text-lg'>{title}</h1>
              <span className='text-sm mt-2'>{description}</span>
            </div>
            <div className='mt-7 mb-4 flex'>
              <Button colorScheme='gray' onClick={deleteModal.onClose} className='w-full'>
                Cancelar
              </Button>
              <Button variant='third' className='ml-2 w-full'>
                Eliminar
              </Button>
            </div>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  )
}

export default DeleteModal;
