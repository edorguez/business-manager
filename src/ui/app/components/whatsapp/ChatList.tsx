'use client';

import { Box, Skeleton, SkeletonCircle, SkeletonText, Stack } from "@chakra-ui/react";

const ChatList = () => {
  const chatsLoading: any[] = [...Array(13)].map(item => {
    return {...item, id: crypto.randomUUID()}
  });

  return (
    <div>
      <Stack>
        {
          chatsLoading.map((item) => (
            <Skeleton key={item.id} height='60px' />
          ))
        }
      </Stack>
    </div>
  )
}

export default ChatList;
