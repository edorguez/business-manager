import { Login } from "@/app/types/auth"

export async function login(
  request: Login,
) {
  const res = await fetch('http://localhost:3001/api/auth/login', {
    method: 'POST',
    body: JSON.stringify(request),
  })

  const data = await res.json()

  return Response.json(data)
}
