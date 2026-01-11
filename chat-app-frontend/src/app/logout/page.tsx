import {  logOut } from "./service";

export const Logout = async () => {
  const response = await logOut()

  if (response) {
    return <p>Logged out successfully</p>;
  }

  return <p>Something went wrong logging out</p>;
};

export default Logout;
