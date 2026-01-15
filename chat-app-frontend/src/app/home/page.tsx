import { headers } from "next/headers";
import { Chat } from "./chat/chat";

export const Home = async () => {
  const h = await headers();
  const userId = h.get("x-user-id");

  return <Chat />;
};

export default Home;
