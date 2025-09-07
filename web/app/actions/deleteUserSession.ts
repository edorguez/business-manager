'use server'

import { cookies } from 'next/headers'

export default async function deleteUserSession() {
  cookies().delete('token');
}
