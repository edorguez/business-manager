'use client';

import SimpleCard from '../../components/cards/SimpleCard';
import { Input } from '@chakra-ui/react'

const CreateCustomerClient = () => {
  return (
    <div>
      <SimpleCard>
        <div>
          <label className='text-sm'>Nombre</label>
          <Input placeholder='Basic usage' size="sm"/>
        </div>
      </SimpleCard>
    </div>
  )
}

export default CreateCustomerClient;
