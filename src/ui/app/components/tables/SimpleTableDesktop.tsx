'use client';

import { SimpleTableColumn, SimpleTableProps } from './SimpleTable.types';
import {
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  TableContainer,
  Button,
  TableCaption,
  Switch
} from '@chakra-ui/react'
import { Icon } from '@iconify/react';
import SimpleTableField from './SimpleTableField';

const SimpleTableDesktop: React.FC<SimpleTableProps> = ({
  columns,
  data,
  size = 'md',
  showToggleActive = false,
  showDetails = false,
  showEdit = false,
  showDelete = false,
  onDelete
}) => {

  return (
    <TableContainer>
      <Table size={size}>
        <TableCaption>
          <div className='flex justify-end'>
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
            {columns.map((col: SimpleTableColumn, index: number) => (
              <Th key={index}>
                {col.name}
              </Th>
            ))}

            {(showEdit || showDelete) &&
              <Th></Th>
            }
          </Tr>
        </Thead>
        <Tbody>
          {data.map((dataVal: any, dataIndex: number) => (

            <Tr key={dataIndex} className='hover:bg-thirdcolorhov transition duration-150 text-sm'>
              {columns.map((col: SimpleTableColumn, colIndex: number) => (

                <Td key={colIndex}>
                  {
                    <SimpleTableField data={dataVal} col={col} />
                  }
                </Td>

              ))}

              <Td>
                <div className='flex justify-end'>

                  {showToggleActive && (
                    <div className='mr-2 flex items-center justify-center'>
                      <span className="text-xs font-bold text-maincolor mr-2">Activo</span>
                      <Switch size='md' colorScheme='main' />
                    </div>
                  )}

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
                    <Button size="sm" variant="third" onClick={onDelete}>
                      <Icon icon="wpf:delete" />
                    </Button>
                  )}
                </div>
              </Td>

            </Tr>

          ))}

        </Tbody>
      </Table>
    </TableContainer>
  );
}

export default SimpleTableDesktop;
