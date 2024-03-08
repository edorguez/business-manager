'use client';

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
} from '@chakra-ui/react'

interface BreadcrumbProps {
  items: string[]
}

const BreadcrumbNavigation: React.FC<BreadcrumbProps> = ({
  items
}) => {

  return (
    <>
      <Breadcrumb fontSize="sm">
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
