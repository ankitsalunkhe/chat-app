import { Button } from "@mui/material";
import styles from "./styles.module.css";
import { Home, Logout } from "@mui/icons-material";

export const NavBar = () => {
  return (
    <nav className={styles.navbar}>
      <Button
        size="large"
        endIcon={<Home />}
        href="/"
      ></Button>
      <Button
        variant="contained"
        size="small"
        endIcon={<Logout />}
        href="/logout"
      >
        Logout
      </Button>
    </nav>
  );
};
