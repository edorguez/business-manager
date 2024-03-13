'use client';

import { SimpleTableProps } from './SimpleTable.types';
import SimpleTableDesktop from './SimpleTableDesktop';
import SimpleTableMobile from './SimpleTableMobile';

const SimpleTable: React.FC<SimpleTableProps> = ({
  columns,
  data,
  size = 'md',
  showDetails = false,
  showEdit = false,
  showDelete = false
}) => {

  const tableProps = {
    columns,
    data,
    size,
    showDetails,
    showEdit,
    showDelete
  }

  return (
    <SimpleTableDesktop {...tableProps} />
  );
}

export default SimpleTable;
