export async function POST(request: Request) {
  const formData = await request.formData();

  formData.append("grant_type", "password");

  const response = await fetch("http://localhost:1323/v1/auth/login", {
    method: "POST",
    body: formData,
  });

  return response;
}
