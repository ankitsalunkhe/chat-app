import styles from "./styles.module.css";

export const NavBar = () => {
  return (
    <nav className={styles.navbar}>
      <a href="/">Home</a>
      <a href="/logout" className={styles.logout}>Logout</a>
    </nav>
  );
};
