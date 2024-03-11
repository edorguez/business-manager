export interface SimpleTableColumn {
  name: string;
  display?: boolean;
}

export interface SimpleTableData {
  value: any;
  type: SimpleTableDataType;
}

enum SimpleTableDataType {
  String,
  Number,
}
