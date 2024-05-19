'use client';

import { Input } from "@chakra-ui/react";

const MessageBar = () => {
  return (
    <div className="bg-slate-200 border-t-2 border-slate-300">
      <Input bg="white" size='sm' placeholder='Basic usage' />
    </div>
  )
}

export default MessageBar;
