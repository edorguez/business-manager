'use client';

import { Icon } from "@iconify/react";

interface SimpleToastProps {
  title: string;
  description: string;
  status: string;
}

const SimpleToast: React.FC<SimpleToastProps> = ({
  title,
  description,
  status
}) => {
  const getToastStyles = (): {bg: string, icon: string} => {
    if (status === 'error') {
      return {bg: 'thirdcolor', icon: 'material-symbols:error'};
    }

    return {bg: 'maincolor', icon: 'ep:success-filled'};
  };
  const toastStyles: {bg: string, icon: string} = getToastStyles();

  return (
    <div className={`
      rounded 
      p-4 
      flex 
      items-center 
      text-white
      bg-${toastStyles.bg}`}>
      <Icon icon={toastStyles.icon} style={{ fontSize: '24px' }} />
      <div className="ml-3">
        <h1 className="text-md font-bold">{title}</h1>
        <p className="text-sm">{description}</p>
      </div>
    </div>
  )
}

export default SimpleToast;
