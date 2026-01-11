"use client";

import { useForm } from "react-hook-form";
import styles from "./form.module.css";
import Image from "next/image";
import { useRouter } from "next/navigation";

type formValue = {
  username: string;
  password: string;
};

export const Form = () => {
  const router = useRouter();

  const {
    handleSubmit,
    register,
    setError,
    formState: { errors, isSubmitting },
  } = useForm<formValue>({
    mode: "onSubmit",
    delayError: 5,
  });

  const onSubmit = async (values: formValue) => {
    try {
      const formData = new FormData();
      formData.append("username", values.username);
      formData.append("password", values.password);

      const response = await fetch("/api/login", {
        method: "POST",
        body: formData,
      });

      if (response.ok) {
        router.push("/home");
      }

      if (response.status === 401) {
        setError("root", {
          type: "manual",
          message: "unable to login",
        });
      }
    } catch (e) {
      setError("root", {
        type: "manual",
        message: e instanceof Error ? e.message : "An error occurred",
      });
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className={styles.form}>
      <div className={styles.usernameContainer}>
        <Image
          src="/icons/username.png"
          className={styles.usernameLogo}
          alt="Username Logo"
          width={22}
          height={23}
        />
        <input
          {...register("username", {
            required: "Username can not be empty",
            pattern: {
              value: /\S+@\S+\.\S+/,
              message: "Username does not match email format",
            },
          })}
          type="text"
          placeholder="Username"
          className={styles.username}
        />
        <div className={styles.errorContainer}>
          {errors.username && (
            <p className={styles.errorMsg}>{errors.username.message}</p>
          )}
        </div>
      </div>
      <div className={styles.passwordContainer}>
        <Image
          src="/icons/password.png"
          className={styles.passwordLogo}
          alt="Password Logo"
          width={23}
          height={24}
        />
        <input
          {...register("password", {
            required: "Password can not be empty",
            minLength: {
              value: 5,
              message: "Password should be minimum 10 characters long",
            },
          })}
          type="password"
          placeholder="Password"
          className={styles.password}
        />
        <div className={styles.errorContainer}>
          {errors.password && (
            <p className={styles.errorMsg}>{errors.password.message}</p>
          )}
        </div>
      </div>
      {errors.root && <p>Error for login</p>}
      <div className={styles.forgotPassCont}>
        <button className={styles.forgotPassBtn} type="button">
          Forgot password?
        </button>
      </div>
      <button
        type="submit"
        className={styles.loginButton}
        disabled={isSubmitting}
      >
        {isSubmitting ? (
          <div className={styles.loggingBtn}>Logging in</div>
        ) : (
          "Login"
        )}
      </button>
    </form>
  );
};
