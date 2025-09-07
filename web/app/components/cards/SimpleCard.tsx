'use client';

import { ReactNode } from "react";

const SimpleCard = ({ children }: { children: ReactNode }) => {

    return (
        <div className="bg-white shadow-lg rounded-md py-2 px-5">
            {children}
        </div>
    );
}

export default SimpleCard;