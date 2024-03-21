'use client';

import { ReactNode } from "react";

const SimpleCardItem = ({ children }: { children: ReactNode }) => {

  return (
    <div className="bg-white rounded-md border border-slate-200 py-2 px-4">
      {children}
    </div>
  );
}

export default SimpleCardItem;
