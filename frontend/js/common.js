const API_BASE = "http://localhost:8080/api/v1";

let currentUser = null;
const savedUser = localStorage.getItem("currentUser");
if (savedUser) {
  currentUser = JSON.parse(savedUser);
}

function updateLoginStatus() {
  const status = document.getElementById("loginStatus");
  if (status && currentUser) {
    status.textContent = `Logged in as ${currentUser.email}`;
  }
}
updateLoginStatus();

function wireLoginForm() {
  const loginButton = document.getElementById("loginButton");
  if (!loginButton) return;

  loginButton.addEventListener("click", () => {
    const email = document.getElementById("loginEmail").value;
    const password = document.getElementById("loginPassword").value;

    fetch(`${API_BASE}/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    })
      .then(response => {
        if (!response.ok) {
          throw new Error("Login failed");
        }
        return response.text();
      })
      .then(result => {
        const userID = result.split(": ")[1];
        currentUser = { email: email, id: userID };
        localStorage.setItem("currentUser", JSON.stringify(currentUser));
        updateLoginStatus();
      })
      .catch(error => {
        document.getElementById("loginStatus").textContent = "Login failed";
        console.error(error);
      });
  });
}
wireLoginForm();
