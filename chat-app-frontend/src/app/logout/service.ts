import { cookies } from "next/headers";

export const logOut = async (): Promise<boolean> => {
  const cookieStore = await cookies();
  const token = cookieStore.get("chatapp-session");

  if (!token) {
    return true;
  }

  const response = await fetch("http://localhost:1323/v1/auth/logout", {
    method: "POST",
    headers: {
      cookie: "chatapp-session=" + token.value,
    },
  });

  if (response.ok) {
    return true;
  }

  return false;
};
