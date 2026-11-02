const gate = document.getElementById("loginGate");
const content = document.getElementById("dashboardContent");

function showDashboard() {
  const session = getSession();
  gate.hidden = Boolean(session);
  content.hidden = !session;
  if (session) {
    document.getElementById("welcome").textContent = "Welcome, " + session.email;
  }
}

window.addEventListener("session-changed", showDashboard);
showDashboard();