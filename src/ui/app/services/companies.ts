import { EditCompany, GetCompany } from "@/app/types/company";
import Cookies from "js-cookie";

const baseUrl: string = "http://localhost:3001/api/companies";

export async function GetCompanyRequest(request: GetCompany) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    const res = await fetch(`${baseUrl}/${request.id}`, {
      method: "GET",
      headers: headers,
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function GetCompanyByNameRequest(name: string) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    const res = await fetch(`${baseUrl}/name/${name}`, {
      method: "GET",
      headers: headers,
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function GetCompanyByNameUrlRequest(nameUrl: string) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    const res = await fetch(`${baseUrl}/nameUrl/${nameUrl}`, {
      method: "GET",
      headers: headers,
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function EditCompanyRequest(request: EditCompany, images: File[]) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    request.name = request.name.trim();

    const formData = new FormData();

    // Add JSON data as a string
    formData.append("json", JSON.stringify(request));

    // Add images to the FormData
    images.forEach((image, index) => {
      formData.append(`files`, image);
    });

    const res = await fetch(`${baseUrl}/${request.id}`, {
      method: "PUT",
      headers: headers,
      body: formData,
    });

    if(res.status === 204) {
      return;
    }

    return await res.json();
  } catch (error: any) {
    console.log(error.toString());
  }
}
