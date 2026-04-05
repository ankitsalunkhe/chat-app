import { Form } from "./form";
import { GoogleLogin } from "./googleLogin";
import styles from "./login.module.css";

const LoginPage = () => {
  return (
    <div className={styles.container}>
      <div className={styles.loginContainer}>
        <div className={styles.loginTitle}>Sign in</div>
        <Form />
        <div className={styles.externalLogin}>
          <p className={styles.externalSignTitle}>Or sign in with</p>
          <GoogleLogin />
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
