import { Form } from "./form";
import styles from "./login.module.css";
import Image from "next/image";

const LoginPage = () => {
  return (
    <div className={styles.container}>
      <div className={styles.loginContainer}>
        <div className={styles.loginTitle}>Sign in</div>
        <Form />
        <div className={styles.externalLogin}>
          <p className={styles.externalSignTitle}>Or sign in with</p>
          <Image
            src="/icons/google-login.png"
            className={styles.googleLogin}
            alt="Google Login"
            width={30}
            height={30}
          />
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
