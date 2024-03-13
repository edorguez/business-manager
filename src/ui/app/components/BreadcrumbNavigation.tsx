'use client';

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
} from '@chakra-ui/react';
import { Icon } from '@iconify/react';
import { BreadcrumItem } from '../types';

interface BreadcrumbProps {
  items: BreadcrumItem[]
}

const BreadcrumbNavigation: React.FC<BreadcrumbProps> = ({
  items
}) => {

  return (
    <>
      <Breadcrumb fontSize="sm" spacing='8px' separator={<Icon icon="octicon:chevron-right-12" />}>
        {items.map((element: BreadcrumItem, index: number) => (
          <BreadcrumbItem key={index}>
            <BreadcrumbLink href={`${element.href}`}>
              {element.label}
            </BreadcrumbLink>
          </BreadcrumbItem>
        ))}
      </Breadcrumb>
    </>
  )
}

export default BreadcrumbNavigation;
