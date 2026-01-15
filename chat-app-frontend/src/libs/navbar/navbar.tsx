import { Button, Link } from "@mui/material";
import styles from "./styles.module.css";
import { Logout } from "@mui/icons-material";

export const NavBar = () => {
  return (
    <nav className={styles.navbar}>
      {/* <a href="/">Home</a> */}
      <Link href="/" underline="none">
        Home
      </Link>
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
