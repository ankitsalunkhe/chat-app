import { headers } from "next/headers";
import { Button } from "@mui/material";
import styles from "./style.module.css";
import ChatIcon from "@mui/icons-material/Chat";
import { Chat } from "./chat/chat";

export const Home = async () => {
  const h = await headers();
  const userId = h.get("x-user-id");

  return (
    <div className={styles.container}>
      <Button variant="contained" startIcon={<ChatIcon />}>
        Chat
      </Button>
      <Chat />
    </div>
  );
};

export default Home;
