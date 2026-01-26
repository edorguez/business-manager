"use client";

import useWarningModal from '@/app/hooks/useWarningModal';
import {
  Modal,
  ModalOverlay,
  ModalContent,
  ModalBody,
  Button,
} from '@chakra-ui/react'
import { Icon } from '@iconify/react';

interface WarningModalProps {
  title: string;
  description: string;
  confirmText: string;
  onSubmit: () => void;
}

const WarningModal: React.FC<WarningModalProps> = ({
  title,
  description,
  confirmText,
  onSubmit
}) => {
  const warningModal = useWarningModal();

  return (
    <>
      <Modal isOpen={warningModal.isOpen} onClose={warningModal.onClose}>
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
              <Button colorScheme='gray' onClick={warningModal.onClose} className='w-full'>
                Cancelar
              </Button>
              <Button variant='third' className='ml-2 w-full' onClick={onSubmit}>
                { confirmText }
              </Button>
            </div>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  )
}

export default WarningModal;
