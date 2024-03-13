'use client';

import { SimpleTableColumn } from './SimpleTable.types';
import {
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  TableContainer,
  Button,
  TableCaption
} from '@chakra-ui/react'
import { Icon } from '@iconify/react';
import { Fragment } from 'react';

interface SimpleTableMobileProps {
  columns: SimpleTableColumn[];
  data: any[];
  size?: string;
  showDetails?: boolean;
  showEdit?: boolean;
  showDelete?: boolean;
}

const SimpleTableMobile: React.FC<SimpleTableMobileProps> = ({
  columns,
  data,
  size = 'md',
  showDetails = false,
  showEdit = false,
  showDelete = false

}) => {
  return (
    <TableContainer>
      <Table variant="none" size={size}>
        <TableCaption>
          <div className='flex justify-center'>
            <div className='flex items-center select-none text-thirdcolor'>
              <div className='cursor-pointer p-2 rounded hover:bg-maincolor duration-150 hover:text-white'>
                <Icon icon="fa:chevron-left" />
              </div>
              <span className='mx-2 font-bold'>
                Página 1
              </span>
              <div className='cursor-pointer p-2 rounded hover:bg-maincolor duration-150 hover:text-white'>
                <Icon icon="fa:chevron-right" />
              </div>
            </div>
          </div>
        </TableCaption>
        <Thead>
          <Tr>
            <Th></Th>
            <Th></Th>
          </Tr>
        </Thead>
        <Tbody>
          {data.map((dataVal: any, dataIndex: number) => 
            [
              columns.map((col: SimpleTableColumn, colIndex: number) => (
                <Tr key={colIndex}>
                  <Th className='text-sm'>
                    {col.name}
                  </Th>
                  <Td className='text-sm'>
                    {dataVal[col.key]}
                  </Td>
                </Tr>
              )),

              <Tr key={dataIndex}>
                <Td>
                </Td>
                <Td>
                  <div className='flex'>
                  {showDetails && (
                    <Button size="sm" variant="fifth">
                      <Icon icon="lucide:info" />
                    </Button>
                  )}

                  {showEdit && (
                    <Button size="sm" variant="main" className="mx-1">
                      <Icon icon="lucide:edit" />
                    </Button>
                  )}

                  {showDelete && (
                    <Button size="sm" variant="third">
                      <Icon icon="wpf:delete" />
                    </Button>
                  )}
                </div>
                </Td>
              </Tr>,

              <Tr key={dataIndex + 1} className='border-t-2 border-gray-200'>
                <Td>
                </Td>
                <Td>
                </Td>
              </Tr>
            ]
              
          )}

        </Tbody>
      </Table>
    </TableContainer>
  )
}

export default SimpleTableMobile;
