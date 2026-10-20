import Cookies from "js-cookie";

const baseUrl: string =
  `${process.env.NEXT_PUBLIC_API_URL ?? "/api"}/roles`;

export async function GetRolesRequest() {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    const res = await fetch(
      `${baseUrl}?` +
        new URLSearchParams({
          limit: "99",
          offset: "0",
        }),
      {
        method: "GET",
        headers: headers,
      }
    );

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}
