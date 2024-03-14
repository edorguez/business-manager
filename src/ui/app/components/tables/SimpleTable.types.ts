export interface SimpleTableProps {
  columns: SimpleTableColumn[];
  data: any[];
  size?: string;
  showDetails?: boolean;
  showEdit?: boolean;
  showDelete?: boolean;
  onDelete?: () => void;
}

export interface SimpleTableColumn {
  key: string;
  name: string;
  display?: boolean;
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
