'use client';

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
} from '@chakra-ui/react';
import { Icon } from '@iconify/react';

interface BreadcrumbProps {
  items: string[]
}

const BreadcrumbNavigation: React.FC<BreadcrumbProps> = ({
  items
}) => {

  return (
    <>
      <Breadcrumb fontSize="sm" spacing='8px' separator={<Icon icon="octicon:chevron-right-12" />}>
        {items.map((element: string, index: number) => (
          <BreadcrumbItem key={index}>
            <BreadcrumbLink href='#'>
              {element}
            </BreadcrumbLink>
          </BreadcrumbItem>
        ))}
      </Breadcrumb>
    </>
  )
}

export default BreadcrumbNavigation;
