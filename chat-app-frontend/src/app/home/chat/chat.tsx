"use client";

import { Card, CardContent } from "@mui/material";
import styles from "./style.module.css";

export const Chat = () => {
  const ws = new WebSocket("ws://localhost:1323/v1/api/chat");
  
  ws.onopen = () => {
    console.log("Connected to WebSocket");
    ws.send("Hello, WebSocket!");
  };

  ws.onerror = (event) => {
    console.error("WebSocket error:", event);
  };
  
  return (
    // <Card sx={{ minWidth: 275 }}>
    //   <CardContent className={styles.cardContent}>test</CardContent>
    // </Card>
    <></>
  );
};
