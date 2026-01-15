import { Stack, Typography } from "@mui/material";

type ChatItemsProps = {
  messages: string[]
}
export const ChatItems = ({messages}: ChatItemsProps) => {
  return (
    <Stack
      direction="column-reverse"
      sx={{
        height: "100%",
        alignContent: "flex-end",
        alignItems: "flex-end",
      }}
    >
      {messages.map((message, index) => (
        <Typography key={index} variant="h6" gutterBottom>
          {message}
        </Typography>
      ))}
    </Stack>
  );
};
