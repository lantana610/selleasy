const API_BASE = "http://localhost:8080/api/v1";

localStorage.removeItem("currentUser");

function getSession() {
  try {
    return JSON.parse(localStorage.getItem("session"));
  } catch (e) {
    return null;
  }
}

function saveSession(session) {
  localStorage.setItem("session", JSON.stringify(session));
}

function clearSession() {
  localStorage.removeItem("session");
}

async function api(path, options = {}) {
  const headers = {};
  const session = getSession();
  if (session) {
    headers["Authorization"] = "Bearer " + session.token;
  }
  if (options.body !== undefined) {
    headers["Content-Type"] = "application/json";
  }

  const response = await fetch(API_BASE + path, {
    method: options.method || "GET",
    headers: headers,
    body: options.body === undefined ? undefined : JSON.stringify(options.body),
  });

  const text = await response.text();
  if (!response.ok) {
    throw new Error(text.trim() || "Something went wrong");
  }
  return text ? JSON.parse(text) : null;
}

function esc(value) {
  return String(value ?? "").replace(/[&<>"']/g, function (ch) {
    return { "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;", "'": "&#39;" }[ch];
  });
}

function setMessage(element, text, kind) {
  if (!element) return;
  element.textContent = text || "";
  element.className = "message" + (kind ? " " + kind : "");
}

function announceSessionChange() {
  window.dispatchEvent(new Event("session-changed"));
}

function renderNav(active) {
  const nav = document.getElementById("navbar");
  const session = getSession();

  nav.innerHTML =
    '<a class="brand" href="index.html">SellEasy</a>' +
    '<div class="links">' +
    '<a href="index.html">Browse</a>' +
    '<a href="dashboard.html">Dashboard</a>' +
    "</div>" +
    '<div class="auth">' +
    (session
      ? "<span>" + esc(session.email) + '</span><button type="button" id="logoutButton">Log out</button>'
      : '<form id="loginForm">' +
        '<input type="email" id="loginEmail" placeholder="Email" required>' +
        '<input type="password" id="loginPassword" placeholder="Password" required>' +
        '<button type="submit">Log in</button>' +
        "</form>") +
    '<span id="authMessage" class="message"></span>' +
    "</div>";

  if (session) {
    document.getElementById("logoutButton").addEventListener("click", logOut);
  } else {
    document.getElementById("loginForm").addEventListener("submit", logIn);
  }
}

async function logIn(event) {
  event.preventDefault();
  const message = document.getElementById("authMessage");
  try {
    const user = await api("/login", {
      method: "POST",
      body: {
        email: document.getElementById("loginEmail").value,
        password: document.getElementById("loginPassword").value,
      },
    });
    saveSession({ token: user.token, id: user.id, email: user.email });
    announceSessionChange();
  } catch (err) {
    setMessage(message, err.message, "error");
  }
}

async function logOut() {
  try {
    await api("/logout", { method: "POST" });
  } catch (err) {
    // we log out locally either way
  }
  clearSession();
  announceSessionChange();
}

window.addEventListener("session-changed", function () {
  renderNav();
});
renderNav();
