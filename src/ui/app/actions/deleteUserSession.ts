export default function deleteUserSession() {
  localStorage.removeItem("token");
}
