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
  showDelete = false,
  onDelete
}) => {

  const tableProps: SimpleTableProps = {
    columns,
    data,
    size,
    showDetails,
    showEdit,
    showDelete,
    onDelete
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
