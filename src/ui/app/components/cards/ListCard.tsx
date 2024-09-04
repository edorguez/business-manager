"use client";

import SimpleTable from "../tables/SimpleTable";
import { SimpleTableColumn } from "../tables/SimpleTable.types";

interface ListCardProps {
  title: string,
  columns: SimpleTableColumn[],
  data: any[],
}

const ListCard: React.FC<ListCardProps> = ({
  title,
  columns,
  data
}) => {
  return (
    <div className="col-span-2 bg-white shadow-lg rounded-md">
      <header className="px-5 py-4">
        <h2 className="font-semibold text-maincolor">
          { title }
        </h2>
      </header>
      <div className="p-1">
          <SimpleTable columns={columns} data={data} offset={undefined} />
      </div>
    </div>
  );
};

export default ListCard;
