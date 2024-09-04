export interface SimpleTableProps {
  columns: SimpleTableColumn[];
  data: any[];
  offset: number | undefined;
  size?: string;
  showToggleActive?: boolean;
  showDetails?: boolean;
  showEdit?: boolean;
  showDelete?: boolean;
  onDelete?: (val: any) => void;
  onDetail?: (val: any) => void;
  onEdit?: (val: any) => void;
  onChangePage?: (val: string) => void;
  onChangeStatus?: (id: any, status: boolean) => void;
}

export interface SimpleTableColumn {
  key: string;
  name: string;
  type: ColumnType;
  display?: boolean;
}

export enum ColumnType {
  String,
  Image,
  ArrayImage,
  ArrayImageFirst,
  Number,
  Money
}

// export interface SimpleTableData {
//   value: any;
//   type: SimpleTableDataType;
// }
//
// enum SimpleTableDataType {
//   String,
//   Number,
// }
