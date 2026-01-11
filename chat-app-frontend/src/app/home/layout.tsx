import { NavBar } from "@/libs/navbar/navbar";
import styles from "./styles.module.css";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className={styles.home}>
      <div className={styles.header}>
        <NavBar />
      </div>
      <div className={styles.sidebar}>sidebar</div>
      <div className={styles.main}>{children}</div>
    </div>
  );
}
