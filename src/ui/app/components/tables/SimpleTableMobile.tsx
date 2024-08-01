"use client";

import { SimpleTableColumn, SimpleTableProps } from "./SimpleTable.types";
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
  Switch,
} from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import SimpleTableField from "./SimpleTableField";
import { useState } from "react";

const SimpleTableMobile: React.FC<SimpleTableProps> = ({
  columns,
  data,
  offset,
  size = "md",
  showToggleActive = false,
  showDetails = false,
  showEdit = false,
  showDelete = false,
  onDelete,
  onDetail,
  onEdit,
  onChangePage,
}) => {
  
  const [pageNumber, setPageNumber] = useState<number>(1);

  const handleChangePage = (val: string) => {
    setPageNumber((prevValue) => val === 'NEXT' ? ++prevValue : --prevValue);

    if (onChangePage) {
      return onChangePage(val);
    }
  }

  const handleEditItem = (val: any) => {
    if(onEdit)
      onEdit(val);
  }

  const handleDeleteItem = (val: any) => {
    if(onDelete)
      onDelete(val);
  }

  const handleDetailItem = (val: any) => {
    if(onDetail)
      onDetail(val);
  }
  return (
    <TableContainer>
      <Table variant="none" size={size}>
        <TableCaption>
          <div className="flex justify-center">
            <div className="flex items-center select-none text-thirdcolor">
              {
                offset >= 10 &&
                <div className='cursor-pointer p-2 rounded hover:bg-maincolor duration-150 hover:text-white' onClick={() => handleChangePage('PREV')}>
                  <Icon icon="fa:chevron-left" />
                </div>
              }
              <span className="mx-2 font-bold">Página { pageNumber }</span>
              {
                data.length !== 0 &&
                <div className='cursor-pointer p-2 rounded hover:bg-maincolor duration-150 hover:text-white' onClick={() => handleChangePage('NEXT')}>
                  <Icon icon="fa:chevron-right" />
                </div>
              }
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
          {data.map((dataVal: any, dataIndex: number) => [
            columns.map((col: SimpleTableColumn, colIndex: number) => (
              <Tr key={colIndex}>
                <Th className="text-sm">{col.name}</Th>
                <Td className="text-sm">
                  <SimpleTableField data={dataVal} col={col} />
                </Td>
              </Tr>
            )),

            <Tr key={dataIndex}>
              <Td></Td>
              <Td>
                {showToggleActive && (
                  <div className="flex items-center justify-center mb-4">
                    <span className="text-xs font-bold text-maincolor mr-2">
                      Activo
                    </span>
                    <Switch size="lg" colorScheme="main" />
                  </div>
                )}

                <div className="flex justify-center">
                  {showDetails && (
                    <Button size="sm" variant="fifth" onClick={() => handleDetailItem(dataVal)}>
                      <Icon icon="lucide:info" />
                    </Button>
                  )}

                  {showEdit && (
                    <Button size="sm" variant="main" className="mx-1" onClick={() => handleEditItem(dataVal)}>
                      <Icon icon="lucide:edit" />
                    </Button>
                  )}

                  {showDelete && (
                    <Button size="sm" variant="third" onClick={() => handleDeleteItem(dataVal) }>
                      <Icon icon="wpf:delete" />
                    </Button>
                  )}
                </div>
              </Td>
            </Tr>,

            <Tr key={dataIndex + 1} className="border-t-2 border-gray-200">
              <Td></Td>
              <Td></Td>
            </Tr>,
          ])}
          
          {
            data.length === 0 && (
              <Tr>
                <Td colSpan={6} >
                  <div className='flex justify-center'>

                  <span>No hay registros para mostrar</span>
                  </div>
                </Td>
              </Tr>
            )
          }
        </Tbody>
      </Table>
    </TableContainer>
  );
};

export default SimpleTableMobile;
