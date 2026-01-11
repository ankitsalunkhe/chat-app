import { headers } from "next/headers";

export const Home = async () => {
  const h = await headers();
  const userId = h.get("x-user-id");

  return <p>Id: {userId ?? "unknown"}</p>;
};

export default Home;
