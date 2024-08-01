'use client';

import { SimpleTableProps } from './SimpleTable.types';
import SimpleTableDesktop from './SimpleTableDesktop';
import SimpleTableMobile from './SimpleTableMobile';

const SimpleTable: React.FC<SimpleTableProps> = ({
  columns,
  data,
  offset,
  size = 'md',
  showToggleActive = false,
  showDetails = false,
  showEdit = false,
  showDelete = false,
  onDelete,
  onDetail,
  onEdit,
  onChangePage
}) => {

  const tableProps: SimpleTableProps = {
    columns,
    data,
    offset,
    size,
    showToggleActive,
    showDetails,
    showEdit,
    showDelete,
    onDelete,
    onDetail,
    onEdit,
    onChangePage
  };

  return (
    <div>
      <div className='xs:hidden sm:hidden md:block'>
        <SimpleTableDesktop {...tableProps} />
      </div>
      <div className='md:hidden lg:hidden xl:hidden'>
        <SimpleTableMobile {...tableProps} />
      </div>
    </div>
  );
}

export default SimpleTable;
