import { ReactNode } from "react";

const PageWrapper = ({ children }: { children: ReactNode }) => {
  return (
    <div className="flex flex-col pt-4 space-y-2 bg-graybackground flex-grow pb-4">
      {children}
    </div>
  );
};

export default PageWrapper;
