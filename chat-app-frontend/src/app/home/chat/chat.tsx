"use client";

import { Box, Button, Stack, TextField } from "@mui/material";
import styles from "./style.module.css";
import SendIcon from "@mui/icons-material/Send";
import { useForm } from "react-hook-form";
import { ChatItems } from "./chatItems";

export const Chat = () => {
  const ws = new WebSocket("ws://localhost:1323/v1/api/chat");

  ws.onopen = () => {
    console.log("Connected to WebSocket");
    ws.send("Hello, WebSocket!");
  };

  type ChatProps = {
    message: string;
  };

  ws.onerror = (event) => {
    console.error("WebSocket error:", event);
  };

  const onSend = (values: ChatProps) => {
    ws.send(values.message);
    reset();
  };

  const {
    register,
    handleSubmit,
    formState: { isValid },
    reset,
  } = useForm<ChatProps>({
    defaultValues: {
      message: "",
    },
  });

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        height: "100%",
      }}
    >
      <Box
        component="form"
        onSubmit={handleSubmit(onSend)}
        sx={{
          display: "flex",
          flexDirection: "column",
          flexGrow: 1,
          gap: 2,
        }}
      >
        <Box
          component="section"
          sx={{
            border: "1px dashed grey",
            flexGrow: 1,
            overflowY: "auto",
          }}
        >
          <Stack
            direction="column-reverse"
            sx={{
              height: "100%",
              alignContent: "flex-end",
              alignItems: "flex-end",
            }}
          >
            {/* <Typography variant="h6" gutterBottom>
              1
            </Typography>
            <Typography variant="h6" gutterBottom>
              2
            </Typography> */}
            <ChatItems messages={["123", "234"]} />
          </Stack>
        </Box>
        <Stack direction="row" spacing={2}>
          <TextField
            label="Chat"
            variant="filled"
            sx={{ flexGrow: 1 }}
            {...register("message", {
              required: true,
            })}
          />
          <Button
            variant="contained"
            size="small"
            type="submit"
            disabled={!isValid}
            endIcon={<SendIcon />}
          >
            Send
          </Button>
        </Stack>
      </Box>
    </Box>
  );
};
