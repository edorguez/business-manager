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

interface SimpleTableProps {
  columns: SimpleTableColumn[];
  data: any[];
  size?: string;
  showEdit?: boolean;
  showDelete?: boolean;
}

const SimpleTable: React.FC<SimpleTableProps> = ({
  columns,
  data,
  size = 'md',
  showEdit = false,
  showDelete = false
}) => {

  return (
    <TableContainer>
      <Table size={size}>
        <TableCaption>
          <div className='flex justify-end'>
            <div className='flex items-center select-none text-thirdcolor'>
              <Icon icon="fa:chevron-left" className='cursor-pointer' />
              <span className='mx-2 font-bold'>
                Página 1
              </span>
              <Icon icon="fa:chevron-right" className='cursor-pointer' />
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

            <Tr key={dataIndex} className='hover:bg-thirdcolorhov'>
              {columns.map((col: SimpleTableColumn, colIndex: number) => (

                <Td key={colIndex}>
                  {dataVal[col.key]}
                </Td>

              ))}

              <Td>
                {showEdit && showDelete && (
                  <div className='flex'>
                    <Button size="sm" variant="main" className="mr-1">
                      <Icon icon="lucide:edit" />
                    </Button>
                    <Button size="sm" variant="third">
                      <Icon icon="wpf:delete" />
                    </Button>
                  </div>
                )}

                {showEdit && !showDelete && (
                  <Button size="sm" variant="main">
                    <Icon icon="lucide:edit" />
                  </Button>
                )}

                {!showEdit && showDelete && (
                  <Button size="sm" variant="third">
                    <Icon icon="wpf:delete" />
                  </Button>
                )}
              </Td>

            </Tr>

          ))}

        </Tbody>
      </Table>
    </TableContainer>
  );
}

export default SimpleTable;
