'use client';

import Image from "next/image";
import { ColumnType, SimpleTableColumn } from "./SimpleTable.types";


interface SimpleTableFieldProps {
  data: any,
  col: SimpleTableColumn
}

const SimpleTableField: React.FC<SimpleTableFieldProps> = ({
  data,
  col
}) => {
  
  const getField: any = () => {
    if(col.type == ColumnType.String) {
      return data[col.key]
    }

    if(col.type == ColumnType.Image) {
      return <Image src={data[col.key]} alt="" width={38} height={38} />
    }

    if(col.type == ColumnType.Number) {
      return data[col.key]
    }

    if(col.type == ColumnType.Money) {
      return data[col.key]
    }

    return null
  };

  return (
    <>{ getField() }</>
  )
}

export default SimpleTableField;
