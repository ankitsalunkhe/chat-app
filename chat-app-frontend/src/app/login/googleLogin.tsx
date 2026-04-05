"use client";

import Image from "next/image";
import styles from "./login.module.css";
import Link from "next/link";

const GoogleLogin = () => {
  return (
    <Link href={"http://localhost:1323/googleLogin"}>
      <Image
        src="/icons/google-login.png"
        className={styles.googleLogin}
        alt="Google Login"
        width={30}
        height={30}
      />
    </Link>
  );
};

export { GoogleLogin };
