"use client";

import SignUpClient from "./SignUpClient";
import { useEffect, useState } from "react";
import Cookies from "js-cookie";
import { useRouter } from "next/navigation";

const SignUpPage = () => {
  const router = useRouter();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const token = Cookies.get('token'); 

    if (token) {
      router.push("/management/home");
    } else {
      setLoading(false);
    }
  }, [router]);

  if (loading) {
    return <div>Loading...</div>;
  }

  return <SignUpClient />;
};

export default SignUpPage;
